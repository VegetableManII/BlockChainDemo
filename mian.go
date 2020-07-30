package main

import (
	"fmt"
)

//重构代码
func main() {
	var bc = NewBlockChain()
	bc.AddBlock("转账100万")
	for i, block := range bc.Blocks {
		fmt.Printf("当前区块高度:%d\n", i)
		fmt.Printf("PreHash:%X\n", block.PreHash)
		fmt.Printf("Hash:%X\n", block.Hash)
		fmt.Printf("Data:%s\n", block.Data)
	}
}
