package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义工作量证明结构
type ProofOfWork struct {
	block *Block
	//目标值
	target *big.Int
}

//创建POW的函数
func NewProoOffWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	targetStr := "00001000000000000000000000000000" +
		"00000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow

}

//提供不断计算hash的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//拼装数据 (不断变化的随机数)
	//做哈希运算
	//与pow中的target比较

	fmt.Println("挖矿中...")
	var nonce uint64
	var hash [32]byte
	block := pow.block
	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PreHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}

		blockInfo := bytes.Join(tmp, []byte{})

		hash = sha256.Sum256(blockInfo)

		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])

		//比较Hash值
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功~ hash:%x nonce:%d\n", hash, nonce)
			break
		} else {
			nonce++
		}

	}
	return hash[:], nonce
}
