package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type MineBody struct {
	Proof int
	Miner string
}

type AddTransactionBody struct {
	Sender string
	Recipient string
	Amount int
}

var bc = NewBlockChain()

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mine", HandleMine).Methods("POST")
	router.HandleFunc("/transactions", HandleAddTransaction).Methods("POST")
	router.HandleFunc("/chain", HandleGetChain).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func HandleMine(w http.ResponseWriter, r *http.Request) {
	lastBlock, _ := bc.getLastBlock()
	lastProof := lastBlock.Proof

	var mineBody MineBody

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&mineBody)

	if bc.isValidProof(lastProof, mineBody.Proof) {
		bc.addTransaction("0", mineBody.Miner, 1)

		bc.addBlock(mineBody.Proof)

		fmt.Println(bc)

		json.NewEncoder(w).Encode("Valid PoW, block has been mined, you earned 1 AWECash !")

		return
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("%s%d", "Invalid PoW, the valid one was ", bc.getProofOfWork(lastProof)))
}

func HandleAddTransaction(w http.ResponseWriter, r *http.Request) {
	var addTransactionBody AddTransactionBody

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&addTransactionBody)

	bc.addTransaction(addTransactionBody.Sender, addTransactionBody.Recipient, addTransactionBody.Amount)

	fmt.Println(bc)

	json.NewEncoder(w).Encode("Transaction added")
}

func HandleGetChain(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bc)
}