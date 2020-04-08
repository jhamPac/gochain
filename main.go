package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain is a standard chain with a slice of Blocks
type BlockChain struct {
	blocks []*Block
}

// Block represents a node in the BlockChain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash creates a hash for the calling Block/Node
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a Block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock adds a Block to the chain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, new)
}

// Genesis creates the first block to start a chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain inititates a chain with a Genesis block
func InitBlockChain() *BlockChain {
	slice := []*Block{Genesis()}
	return &BlockChain{slice}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n:", block.Hash)
	}
}
