package main

import (
	"BlockChainDemo/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	//Hash指针
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		//最初指向最后一个区块  随着Next的调用向前移动
		bc.tail,
	}
}

//迭代器的Next方式
//返回当前的区块
//指针前移
func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blockBucket))
		if bkt == nil {
			log.Panic("非法:Bucket内容为空")
		}
		blockTmp := bkt.Get(it.currentHashPointer)
		//解码
		block = DeSerialize(blockTmp)
		it.currentHashPointer = block.PreHash

		return nil
	})

	return &block
}
