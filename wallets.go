package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
)

type Wallets struct {
	//map[钱包地址]钱包
	WalletsMap map[string]*Wallet
}

func NewWallets() *Wallets {
	var wallets Wallets
	wallets.WalletsMap = make(map[string]*Wallet)
	return &wallets
}
func (ws *Wallets) CreatWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()

	ws.WalletsMap[address] = wallet
	ws.saveToFile()
	return address
}

func (ws *Wallets) saveToFile() {
	var buffer bytes.Buffer
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}
	_ = ioutil.WriteFile("./wallets/wallet.dat", buffer.Bytes(), 0600)
}
