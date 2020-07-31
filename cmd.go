package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("添加成功!")
}

func (cli *CLI) PrintBlockChain() {
	it := cli.bc.NewIterator()
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
}
