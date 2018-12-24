package state

import (
	"container/list"
	"fmt"
	"github.com/spacemeshos/go-spacemesh/common"
	"github.com/spacemeshos/go-spacemesh/crypto/sha3"
	"github.com/spacemeshos/go-spacemesh/log"
	"github.com/spacemeshos/go-spacemesh/rlp"
	"github.com/spacemeshos/go-spacemesh/trie"
	"math/big"
	"sort"
)


type LayerID uint64

//todo: this object should be splitted into two parts: one is the actual value serialized into trie, and an containig obj with caches
type Transaction struct {
	AccountNonce 	uint64
	Price			*big.Int
	GasLimit		uint64
	Recipient 		*common.Address
	Origin			common.Address //todo: remove this, should be calculated from sig.
	Amount       	*big.Int
	Payload      	[]byte

	//todo: add signatures

	hash *common.Hash
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func (tx *Transaction) Hash() common.Hash{
	if tx.hash == nil {
		hash := rlpHash(tx)
		tx.hash = &hash
	}
	return *tx.hash
}

type Transactions []*Transaction

type PseudoRandomizer interface {
	Uint32() uint32
	Uint64() uint64
}

type StatePreImages struct {
	rootHash common.Hash
	preImages Transactions
}

type TransactionProcessor struct {
	rand PseudoRandomizer
	globalState *StateDB
	prevStates map[LayerID]common.Hash
	currentLayer LayerID
	rootHash common.Hash
	stateQueue list.List
	db *trie.Database
}

const maxPastStates = 20

func NewTransactionProcessor(rnd PseudoRandomizer, db *StateDB) *TransactionProcessor{
	return &TransactionProcessor{
		rand: rnd,
		globalState:db,
		prevStates : make(map[LayerID]common.Hash),
		currentLayer: 0,
		rootHash: common.Hash{},
		stateQueue: list.List{},
		db : db.TrieDB(),
	}
}

//should receive sort predicate
func (tp *TransactionProcessor) ApplyTransactions(layer LayerID, transactions Transactions) error{
	//todo: need to seed the mersenne twister with random beacon seed

	txs := tp.mergeDoubles(transactions)
	tp.Process(tp.randomSort(txs), tp.coalesceTransactionsBySender(txs))

	newHash , err := tp.globalState.Commit(false)
	if err != nil {
		log.Error("db write error %v", err)
		return err
	}

	tp.stateQueue.PushBack(newHash)
	if tp.stateQueue.Len() > maxPastStates {
		hash := tp.stateQueue.Remove(tp.stateQueue.Back())
		tp.db.Commit(hash.(common.Hash),false)
	}
	//tp.prevStates[layer] =
	tp.db.Reference(newHash, common.Hash{})
	//	//call merge
	//	//eliminate doubles
	//	//check for double spends? (how do i do this? check the nonce against the prev one?
	return nil
}

func (tp *TransactionProcessor) Reset(layer LayerID){
	if state, ok := tp.prevStates[layer]; ok {
		newState, err := New(state, tp.globalState.db)
		if err != nil {
			panic("wtf")
		}
		tp.globalState = newState
		//lock mutex
		tp.pruneAfterRevert(layer)
	}
}



func (tp *TransactionProcessor) mergeDoubles(transactions Transactions) Transactions{
	transactionSet := make(map[common.Hash]struct{})
	merged := make(Transactions, 0, len(transactions))
	for _,trns := range transactions {
		if _,ok := transactionSet[trns.Hash()]; !ok {
			transactionSet[trns.Hash()] = struct{}{}
			merged = append(merged, trns)
		}
	}
	return merged
}

func (tp *TransactionProcessor) randomSort(transactions Transactions) Transactions{
	vecLen := len(transactions)
	for i := range transactions {
		swp := int(tp.rand.Uint32()) % vecLen
		tmp := transactions[i]
		transactions[i] = transactions[swp]
		transactions[swp] = tmp
	}
	return transactions
}

func (tp *TransactionProcessor) coalesceTransactionsBySender(transactions Transactions) map[common.Address][]*Transaction {
	trnsBySender := make(map[common.Address][]*Transaction)
	for _, trns := range transactions {
		trnsBySender[trns.Origin] = append(trnsBySender[trns.Origin], trns)
	}

	for key := range trnsBySender{
		sort.Slice(trnsBySender[key], func(i, j int) bool {
			return trnsBySender[key][i].AccountNonce < trnsBySender[key][j].AccountNonce
		})
	}

	return trnsBySender
}

func (tp *TransactionProcessor) Process(transactions Transactions, trnsBySender map[common.Address][]*Transaction) {
	bySender := make(map[common.Hash]bool)
	for _, trans := range transactions {
		for _, trns := range trnsBySender[trans.Origin] {
			if _, ok := bySender[trns.Hash()]; !ok {
				bySender[trans.Hash()] = true
				//todo: should we abort all transaction processing if we failed this one?
				err := tp.ApplyTransaction(trans)
				if err != nil {
					log.Error("transaction aborted: %v", err)
				}
			}
		}
	}
}

func (tp *TransactionProcessor) pruneAfterRevert(targetLayerID LayerID){
	for i:= tp.currentLayer; i >= targetLayerID; i-- {
		if hash, ok := tp.prevStates[i]; ok {
			if tp.stateQueue.Front().Value != hash {
				panic("old state wasn't found")
			}
			tp.stateQueue.Remove(tp.stateQueue.Front())
			tp.db.Dereference(hash)
		}
	}
}

func (tp *TransactionProcessor) checkNonce(trns *Transaction) bool{
	return tp.globalState.GetNonce(trns.Origin) == trns.AccountNonce
}

var(
	ErrOrigin = "origin account doesnt exist"
	ErrFunds = "insufficient funds"
	ErrNonce = "incorrect nonce"
)
//todo: mining fees...
func (tp *TransactionProcessor) ApplyTransaction(trans *Transaction) error{

	if !tp.globalState.Exist(trans.Origin) {
		return  fmt.Errorf(ErrOrigin)
	}

	origin := tp.globalState.GetOrNewStateObj(trans.Origin)

	//todo: should we allow to spend all accounts data
	if origin.Balance().Cmp(trans.Amount) <= 0 {
		return  fmt.Errorf(ErrFunds)
	}

	if !tp.checkNonce(trans) {
		return  fmt.Errorf(ErrNonce)
	}

	tp.globalState.SetNonce(trans.Origin, tp.globalState.GetNonce(trans.Origin) + 1)
	transfer(tp.globalState,trans.Origin, *trans.Recipient, trans.Amount)

	return nil
	//check if dst account exists
	//check if src exist
	//check if src account has enough funds
	//set journal backup
	//add 1 to account nonce
	//verify current nonce
	//upate accounts accordingly
	//error if no funds
	//commit to tree
}

func transfer(db GlobalStateDB, sender, recipient common.Address, amount *big.Int) {
	db.SubBalance(sender, amount)
	db.AddBalance(recipient, amount)
}