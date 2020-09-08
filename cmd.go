package main

import (
	"fmt"
)

func (cli *CLI) AddBlock(data string) {
	//cli.bc.AddBlock(data)
	//todo
	fmt.Println("添加成功!")
}

func (cli *CLI) PrintBlockChain() {
	it := cli.bc.NewIterator()
	//调用迭代器
	for {
		block := it.Next() //指针左移

		fmt.Printf("PreHash:%X\n", block.PreHash)
		fmt.Printf("Hash:%X\n", block.Hash)
		fmt.Printf("Data:%v\n", block.Data)
		fmt.Printf("Sig:%s\n", block.Data[0].TXInputs[0].Sig)

		if len(block.PreHash) == 0 {
			fmt.Printf("遍历区块链完成\n")
			break
		}
	}
}
func (cli *CLI) GetBalance(address string) {
	utxo := cli.bc.FindUTXOs(address)
	total := 0.0
	for _, txo := range utxo {
		total += txo.Value
	}
	fmt.Printf("\"%s\"的余额为%f", address, total)
}
func (cli *CLI) Send(from, to string, amount float64, miner, data string) {
	//创建挖矿交易
	coinbase := NewCoinBase(miner, data)
	//创建普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx != nil {
		//添加到区块
		cli.bc.AddBlock([]*Transaction{coinbase, tx})
	}
}
