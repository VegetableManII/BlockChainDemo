package main

import (
	"bytes"
	"encoding/binary"
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

//uint64转[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer

	//使用gob进行序列化
	//定义编码器
	//使用编器及进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic("序列化出错")
	}
	return buffer.Bytes()
}

func DeSerialize(data []byte) Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	//使用解码器解码
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("反序列化出错")
	}
	return block
}

//创建区块
func NewBlock(txs []*Transaction, preBlock []byte) *Block {
	block := Block{
		Version:    00,
		PreHash:    preBlock,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0, //无效值
		Nonce:      0,
		Hash:       []byte{},
		Data:       txs,
	}
	//block.SetHash()
	block.MerkelRoot = block.MakeMerkelRoot()
	pow := NewProoOffWork(&block)
	hash, nonce := pow.Run()
	//根据挖矿结果对区块数据进行更新(补充)
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

//创世块
func GenesisBlock(address string) *Block {
	coinbase := NewCoinBase(address, "创世块")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) MakeMerkelRoot() []byte {
	//todo
	return []byte{}
}

/*
//生成哈希
func (block *Block)SetHash()  {
	//var blockInfo []byte
	//拼装数据
	/*
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PreHash...)
	blockInfo = append(blockInfo, block.PreHash...)
	blockInfo = append(blockInfo, block.PreHash...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)

	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PreHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	blockInfo := bytes.Join(tmp,[]byte{})
	//SHA256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
*/
