package gochain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Difficulty simulates the chain POW difficulty like on Bitcoin
const Difficulty = 18

// ProofOfWork object to determine network difficulty
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof creates a POW for the chain
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}
	return pow
}

// InitData initializes a POW object using a nonce
func (p *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			p.Block.PrevHash,
			p.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{})
	return data
}

// Run exectues a proof of work
func (p *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := p.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()
	return nonce, hash[:]
}

// Validate ensures blocks are valid with the POW target
func (p *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := p.InitData(p.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(p.Target) == -1
}

// ToHex converts an int to a hex number
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
