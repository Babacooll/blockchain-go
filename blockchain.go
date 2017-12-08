package main

import (
	"time"
	"errors"
	"fmt"
)

type BlockChain struct {
	Blocks       []Block
	Transactions []Transaction
}

func NewBlockChain() BlockChain {
	var bc = BlockChain{}

	var firstBlock = Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Proof:        100,
		PreviousHash: "1",
	}

	bc.Blocks = append(bc.Blocks, firstBlock)

	return bc
}

func (bc *BlockChain) addTransaction(sender string, recipient string, amount int) {
	var newTransaction = Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}

	bc.Transactions = append(bc.Transactions, newTransaction)
}

func (bc *BlockChain) addBlock(proof int) {
	var lastBlock, err = bc.getLastBlock()

	if err == nil {
		var newBlock = Block{
			Index:        bc.getCurrentIndex() + 1,
			Timestamp:    time.Now().Unix(),
			Transactions: bc.Transactions,
			Proof:        proof,
			PreviousHash: lastBlock.hash(),
		}

		bc.Blocks = append(bc.Blocks, newBlock)

		bc.Transactions = []Transaction{}
	} else {
		fmt.Println("ERROR")
	}
}

func (bc BlockChain) getLastBlock() (Block, error) {
	var currentIndex = bc.getCurrentIndex()

	if currentIndex >= 0 {
		return bc.Blocks[currentIndex], nil
	}

	return Block{}, errors.New("there is no last block")
}

func (bc BlockChain) getCurrentIndex() int {
	return len(bc.Blocks) - 1
}

func (bc *BlockChain) getProofOfWork(lastProof int) int {
	var proof = 0

	for !bc.isValidProof(lastProof, proof) {
		proof += 1
	}

	return proof
}

func (bc *BlockChain) isValidProof(lastProof int, proof int) bool {
	guess := fmt.Sprintf("%d%d", lastProof, proof)
	guessHash := ComputeHashSha256([]byte(guess))

	return guessHash[:4] == "0000"
}
