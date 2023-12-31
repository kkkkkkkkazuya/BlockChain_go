package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

/**
* Blcok struct
**/
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

/**
* Blcok constructor
**/
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

/**
* Blcok　Print log
**/
func (b *Block) Print() {
	fmt.Printf("timestamp:        %d\n", b.timestamp)
	fmt.Printf("nonce:            %d\n", b.nonce)
	fmt.Printf("previousHash:     %s\n", b.previousHash)
	fmt.Printf("transactions:     %s\n", b.transactions)

}

/**
* Blcokchain
**/
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

/**
* Blcokchain constructor
**/
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

/**
* Create Block
**/
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockchain()
	blockChain.CreateBlock(1, "hash 1")
	blockChain.Print()
}
