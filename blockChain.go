package main

//引入区块链
type BlockChain struct {
	//定义区块链数组
	Blocks []*Block
}

//定义一个区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		Blocks: []*Block{
			genesisBlock,
		},
	}
}

//添加区块
func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	preHash := lastBlock.Hash

	block := NewBlock(data, preHash)

	bc.Blocks = append(bc.Blocks, block)
}
