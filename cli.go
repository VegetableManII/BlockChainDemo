package main

import (
	"fmt"
	"os"
)

//接受命令行参数并控制区块链操作

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA  "add data to blockchain"
	printChain   "print the blockchain"
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
	default:
		fmt.Printf(Usage)

	}
	//根据命令控制
}
