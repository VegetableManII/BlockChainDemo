package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func (cli *CLI) PrintBlockChain() {
	it := cli.bc.NewIterator()
	//调用迭代器
	for {
		block := it.Next() //指针左移

		fmt.Printf("PreHash:%X\n", block.PreHash)
		fmt.Printf("Hash:%X\n", block.Hash)
		fmt.Printf("Sig:%s\n", block.Data[0].TXInputs[0].Sig)

		if len(block.PreHash) == 0 {
			fmt.Printf("遍历区块链完成\n")
			break
		}
	}
}

func (cli *CLI) PrintTransactionDetails(hash string) {
	it := cli.bc.NewIterator()
	for true {
		block := it.Next()

		encodeStr := hex.EncodeToString(block.Hash)
		//fmt.Printf("当前区块Hash:%s", encodeStr)

		if strings.EqualFold(hash, encodeStr) {
			txs := block.Data
			for i, tx := range txs {
				fmt.Printf("第%d个交易信息,交易ID:%X\n", i, tx.TXID)
				for j, input := range tx.TXInputs {
					fmt.Printf("\t第%d个交易输入信息:\n", j)
					fmt.Printf("\t \t引用交易ID:%X\n\t \t引用交易索引:%d\n\t \t签名:%s\n", input.TXid, input.Index, input.Sig)
				}
				for k, input := range tx.TXOutputs {
					fmt.Printf("\t第%d个交易输出信息:\n", k)
					fmt.Printf("\t \t交易金额%f\n\t \t收款人%s\n", input.Value, input.PubKeyHash)
				}
			}
		}
		if len(block.PreHash) == 0 {
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
func (cli *CLI) NewWallet() {
	//ws := NewWallets()
	//for address := range ws.WalletsMap {
	//	fmt.Printf("地址：%s\n", address)
	//}
	ws := NewWallets()
	s := ws.CreatWallet()
	fmt.Printf("%s\n", s)
}
