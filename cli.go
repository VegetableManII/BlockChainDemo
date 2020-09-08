package main

import (
	"fmt"
	"os"
	"strconv"
)

//接受命令行参数并控制区块链操作

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA  "add data to blockchain"
	printChain   "print the blockchain"
	getBalance --address ADDRESS "get the balance of address"
	sendCoin FROM TO AMOUNT MINER DATA "FROM转账AMOUNT给TO由MINER进行挖矿"
`

//接受参数的动作
func (cli *CLI) Run() {
	//得到命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	//分析命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		//添加区块
		//确保命令有效
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		}
	case "printChain":
		//打印区块链
		cli.PrintBlockChain()
	case "getBalance":
		//查找指定地址的余额
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cli.GetBalance(address)
		}
	case "sendCoin":
		if len(args) == 7 {
			from := args[2]
			to := args[3]
			amount, _ := strconv.ParseFloat(args[4], 64)
			miner := args[5]
			data := args[6]
			cli.Send(from, to, amount, miner, data)
		} else {
			fmt.Printf(Usage)
		}

	default:
		fmt.Printf(Usage)

	}
	//根据命令控制
}
