package main

import "BlockChainDemo/boltdb/bolt"

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
