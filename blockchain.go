package main

import (
	"fmt"
	"github.com/boltdb/bolt"
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
const blockBucket = "blockBucket" //bolt数据库中的bucket名称
/*
区块是保存在bolt数据库中
K是该区块的HASH值
V是该区块序列化之后的数据
*/
//定义一个区块链
func NewBlockChain(address string) *BlockChain {
	var lastHash []byte //最后一个块的hash
	//打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	/*
		几个块都保存在一个Bucket当中
		获取数据库中的最后一个区块
		当最后一个区块不存在则创建创世块
	*/
	_ = db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blockBucket))
		if bkt == nil {
			//没有bucket
			bkt, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
				return err
			}
			/*创建创世块*/
			genesisBlock := GenesisBlock(address)
			//向数据库中写数据
			_ = bkt.Put([]byte(genesisBlock.Hash), genesisBlock.Serialize())
			_ = bkt.Put([]byte("LastHashKey"), genesisBlock.Hash)
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
	/*返回区块链的最后一个区块的Hash值*/
	return &BlockChain{db, lastHash}

}

/*添加区块*/
func (bc *BlockChain) AddBlock(txs []*Transaction) {

	db := bc.db
	lastHash := bc.tail

	_ = db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blockBucket))
		if bkt == nil {
			//没有bucket 错误
			log.Panic("BlockBucket不应为空")
		}
		block := NewBlock(txs, lastHash)

		_ = bkt.Put(block.Hash, block.Serialize())
		_ = bkt.Put([]byte("LastHashKey"), block.Hash)
		//更新内存中的LastHash值
		bc.tail = block.Hash
		return nil
	})
}

/*根据address即交易输入的签名和交易输出的公钥地址，来搜集所有跟这个address有关的UTXO*/
func (bc *BlockChain) FindUTXOTransactions(address string) []*Transaction {
	var txs []*Transaction                   //收集到的所有跟address有关的交易
	spentOutputs := make(map[string][]int64) //用来在搜集过程中标记花费过的output

	it := bc.NewIterator()

	for {
		//1.遍历区块
		block := it.Next()
		//2.遍历交易
		for i, tx := range block.Data {
			fmt.Printf("第%d个交易---\n", i)
			fmt.Printf("-----%v\n", spentOutputs)
			//3.遍历output找到和address相关的utxo（在添加output之前检查是否被消耗过）
		OUTPUT:
			for it, output := range tx.TXOutputs {
				fmt.Printf("\t\t--这个交易的第%d个输出:数额%f值%s\n", it, output.Value, output.PubKeyHash)
				//过滤Output,将所有消耗过的output和当前即将添加的ouput对比一下
				//相同则跳过，否则添加到spentOutputs
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if j == int64(it) {
							continue OUTPUT
						}
					}
				}
				/*收集所有交易输出的公钥地址和我们目标地址一致的交易信息*/
				if output.PubKeyHash == address {
					txs = append(txs, tx)
				}
			}
			//跳过挖矿交易
			if !tx.IsCoinbase() {
				fmt.Printf("\t--此交易非挖矿交易\n")
				//4.遍历input找到address花费过的utxo的集合(把自己消耗过的标记出来)
				for _, input := range tx.TXInputs {
					if input.Sig == address {
						fmt.Printf("\t--此交易的交易txID：%X, txIndex：%d\n", input.TXid, input.Index)
						spentOutputs[string(input.TXid)] = append(spentOutputs[string(input.TXid)], input.Index)
					}
				}
			}
		}
		if len(block.PreHash) == 0 {
			fmt.Printf("区块遍历完成")
			break
		}
	}
	return txs
}

func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]int64, float64) {
	//找到的合理的utxos集合
	utxos := make(map[string][]int64)
	var calc float64

	txs := bc.FindUTXOTransactions(from)

	for _, tx := range txs {
		for i, output := range tx.TXOutputs {
			if from == output.PubKeyHash {
				if calc < amount {
					utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], int64(i))
					calc += output.Value
					if calc >= amount {
						return utxos, calc
					}
				}
			}
		}
	}
	return utxos, calc
}

/*遍历各个区块中的所有交易查找余额*/
func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput

	txs := bc.FindUTXOTransactions(address)
	for _, tx := range txs {
		for _, output := range tx.TXOutputs {
			if output.PubKeyHash == address {
				UTXO = append(UTXO, output)
			}
		}
	}
	return UTXO
}
