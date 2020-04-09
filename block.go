package gochain

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain is a standard chain with a slice of Blocks
type BlockChain struct {
	Blocks []*Block
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
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, new)
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
