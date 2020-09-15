package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reward = 12.5

//定义交易结构
type Transaction struct {
	TXID      []byte     //交易ID
	TXInputs  []TXInput  //交易输入数组
	TXOutputs []TXOutput //交易输出数组
}
type TXInput struct {
	//引用的交易ID
	TXid []byte
	//引用的output的索引值
	Index int64
	//解锁脚本，用地址来模拟
	Sig string
}
type TXOutput struct {
	//转账金额
	Value float64
	//锁定脚本，用地址模拟
	PubKeyHash string
}

//设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

//提供创建交易的方法

//创建普通交易
/*
1.找到最合理的UTXO集合
2.将UTXO逐一转成inputs
3.创建outputs
4.有零钱要找零
*/
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	//找到合理的UTXO集合
	utxos, resValue := bc.FindNeedUTXOs(from, amount)
	if resValue < amount {
		fmt.Printf("余额不足%f", resValue)
		return nil
	}
	var inputs []TXInput
	var outputs []TXOutput
	//将这些UTXO逐一转成input
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{
				[]byte(id),
				int64(i),
				from,
			}
			inputs = append(inputs, input)
		}
	}
	output := TXOutput{amount, to}
	outputs = append(outputs, output)

	if resValue > amount {
		//找零
		outputs = append(outputs, TXOutput{resValue - amount, from})
	}

	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	return &tx
}

//判断是否是挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	// input只有一个
	// id为空
	// index为-1
	if len(tx.TXInputs) == 1 && tx.TXInputs[0].TXid == nil && tx.TXInputs[0].Index == -1 {
		return true
	}
	return false
}

//创建挖矿交易
func NewCoinBase(address string, data string) *Transaction {
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetHash()
	return &tx
}
