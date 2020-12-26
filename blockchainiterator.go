package main

import (
	"bytes"
	"encoding/gob"
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	//Hash指针
	currentHashPointer []byte
}

func DeSerialize(data []byte) Block {
	var block Block
	/*
		使用gob进行反序列化
		定义解码器
		使用解码器及进行解码
	*/
	decoder := gob.NewDecoder(bytes.NewReader(data))
	//使用解码器解码
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("反序列化出错")
	}
	return block
}
func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		//最初指向最后一个区块  随着Next的调用向前移动
		bc.tail,
	}
}

/*
迭代器的Next方式
返回当前的区块
指针前移
*/
func (it *BlockChainIterator) Next() *Block {
	var block Block
	_ = it.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blockBucket))
		if bkt == nil {
			log.Panic("非法:Bucket内容为空")
		}
		blockTmp := bkt.Get(it.currentHashPointer)
		//取出来的数据是空的？
		if blockTmp != nil {
			//解码
			block = DeSerialize(blockTmp)
		}
		it.currentHashPointer = block.PreHash
		return nil
	})
	return &block
}
