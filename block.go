package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

//定义结构
type Block struct {
	//补充结构 版本 梅克尔根 时间戳  难度值  随机数
	Version uint64
	//前驱块Hash
	PreHash []byte
	//Merkel根
	MerkelRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Difficulty uint64
	//随机数
	Nonce uint64
	//当前块Hash
	Hash []byte
	//数据
	Data []*Transaction
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	/*
		使用gob进行序列化
		定义编码器
		使用编码器及进行编码
	*/
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic("序列化出错")
	}
	return buffer.Bytes()
}

func (b *Block) MakeMerkelRoot() []byte {
	var info []byte
	/*
		梅克尔根的生成
		不进行二叉树处理
		简单的对交易进行简单的拼接
	*/
	for _, tx := range b.Data {
		info = append(info, tx.TXID...)
	}
	hash := sha256.Sum256(info)
	return hash[:]
}

/*创建区块*/
func NewBlock(txs []*Transaction, preBlock []byte) *Block {
	block := Block{
		Version:    00,       //版本默认为00
		PreHash:    preBlock, //前一区块的hash值，创世块的该值为0
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,        //难度，随着区块链的不断延长难度值会逐渐上升
		Nonce:      0,        //随机数
		Hash:       []byte{}, //当前区块的Hash值
		//这两个值用来表示矿工的工作量证明
		Data: txs,
	}
	block.MerkelRoot = block.MakeMerkelRoot()
	pow := NewProoOffWork(&block)
	hash, nonce := pow.Run()
	//根据挖矿结果对区块数据进行更新
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

//创世块

/*
参数即为公钥地址
每一个区块的产生都要有coinbase即矿工挖矿所产生的的交易
*/
func GenesisBlock(address string) *Block {
	coinbase := NewCoinBase(address, "创世块")
	//创世块中只有一笔coinbase交易
	return NewBlock([]*Transaction{coinbase}, []byte{})
}
