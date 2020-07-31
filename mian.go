package main

//重构代码
func main() {
	var bc = NewBlockChain()
	cli := CLI{bc}
	cli.Run()
	//bc.AddBlock("转账100万")
	//bc.AddBlock("转账80万")
	/*

	 */

	//for i, block := range bc.B {
	//	fmt.Printf("当前区块高度:%d\n", i)
	//	fmt.Printf("PreHash:%X\n", block.PreHash)
	//	fmt.Printf("Hash:%X\n", block.Hash)
	//	fmt.Printf("Data:%s\n", block.Data)
	//}
}
