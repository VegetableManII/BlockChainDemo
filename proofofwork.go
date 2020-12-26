package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
)

/*定义POW结构*/
type ProofOfWork struct {
	//挖取得区块
	block *Block
	//目标值
	target *big.Int
}

//uint64转[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

/*POW*/
func NewProoOffWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//难度值，需要计算出一个HASH值其要小于等于targetStr
	targetStr := "00001000000000000000000000000000" +
		"00000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

/*提供挖矿功能的函数*/
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	/*
		拼装数据 (不断变化的随机数nonce)
		做哈希运算
		与pow中的target比较
	*/
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
		}
		//对区块数据进行拼接
		blockInfo := bytes.Join(tmp, []byte{})
		//加密
		hash = sha256.Sum256(blockInfo)
		//
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		/*
			比较Hash值
			Cmp compares x and y and returns:

			  -1 if x <  y
			   0 if x == y
			  +1 if x >  y
		*/
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功~ hash:%X nonce:%d\n", hash, nonce)
			break
		} else {
			nonce++
		}

	}
	return hash[:], nonce
}
