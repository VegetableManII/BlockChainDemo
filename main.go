package main

//重构代码
func main() {
	//这里的User-0即是该区块的生产者的公钥
	var bc = NewBlockChain("User-0")
	//取得存储区块的数据库指针将其传递到CLI中以进行上链操作
	cli := CLI{bc}
	cli.Run()
}
