package main

import (
	"BlockChainDemo/boltdb/bolt"
	"log"
)

//引入区块链
type BlockChain struct {
	//定义区块链数组
	//Blocks []*Block
	db *bolt.DB

	tail []byte
}

const blockChainDb = "./bcdb/blockChain.db"
const blockBucket = "blockBucket"

//定义一个区块链
func NewBlockChain() *BlockChain {
	//return &BlockChain{
	//	Blocks: []*Block{
	//		genesisBlock,
	//	},
	//}
	var lastHash []byte //最后一个块的hash
	//打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blockBucket))
		if bkt == nil {
			//没有bucket
			bkt, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
				return err
			}
			genesisBlock := GenesisBlock()
			//向数据库中写数据
			bkt.Put([]byte(genesisBlock.Hash), genesisBlock.Serialize())
			bkt.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash

			//测试读取数据反序列化
			//blockBytes := bkt.Get(genesisBlock.Hash)
			//blockInfo := DeSerialize(blockBytes)
			//fmt.Printf("读取到反序列化数据为:%s\n",blockInfo.Data)
		} else {
			lastHash = bkt.Get([]byte("LastHashKey"))
		}

		return nil
	})
	return &BlockChain{db, lastHash}

}

//添加区块
func (bc *BlockChain) AddBlock(data string) {
	//lastBlock := bc.Blocks[len(bc.Blocks)-1]
	//preHash := bc.tail
	//bc.Blocks = append(bc.Blocks, block)
	db := bc.db
	lastHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blockBucket))
		if bkt == nil {
			//没有bucket 错误
			log.Panic("BlockBucket不应为空")
		}
		block := NewBlock(data, lastHash)

		bkt.Put(block.Hash, block.Serialize())
		bkt.Put([]byte("LastHashKey"), block.Hash)
		//更新内存中的LastHash值
		bc.tail = block.Hash

		return nil
	})

}
