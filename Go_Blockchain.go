package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp         %d\n", b.timestamp)
	fmt.Printf("nonce             %d\n", b.nonce)
	fmt.Printf("previousHash      %s\n", b.previousHash)
	fmt.Printf("transactions      %s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _:= json.Marshal(b)
	fmt.Println(m)
	return sha256.Sum256([]byte(m))
}

type Blockchain struct {
	transactionsPool []string
	chain            []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreatBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) CreatBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, 
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())
/* 	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreatBlock(5, "hash 1")
    blockChain.Print()
    blockChain.CreatBlock(2, "hash 2")
    blockChain.Print()
 */}
