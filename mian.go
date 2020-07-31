package main

import "fmt"

//重构代码
func main() {
	var bc = NewBlockChain()
	bc.AddBlock("转账100万")
	bc.AddBlock("转账80万")
	it := bc.NewIterator()
	//调用迭代器
	for {
		block := it.Next() //指针左移

		fmt.Printf("PreHash:%X\n", block.PreHash)
		fmt.Printf("Hash:%X\n", block.Hash)
		fmt.Printf("Data:%s\n", block.Data)

		if len(block.PreHash) == 0 {
			fmt.Printf("遍历区块链完成\n")
			break
		}
	}

	//for i, block := range bc.B {
	//	fmt.Printf("当前区块高度:%d\n", i)
	//	fmt.Printf("PreHash:%X\n", block.PreHash)
	//	fmt.Printf("Hash:%X\n", block.Hash)
	//	fmt.Printf("Data:%s\n", block.Data)
	//}
}
