package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/golangcrypto/ripemd160"
	"github.com/shengdoushi/base58"
	"log"
)

type Wallet struct {
	//私钥
	Private *ecdsa.PrivateKey
	//公钥
	//Publick *ecdsa.PublicKey
	//这里不存储原始公钥而是存储X和Y的凭借字符串
	//在接收端分割X和Y
	Public []byte
}

//创建钱包
func NewWallet() *Wallet {
	//创建曲线
	curve := elliptic.P256()
	//生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic("私钥生成失败")
	}

	//生成公钥
	pubKeyOrig := privateKey.PublicKey
	//拼接X,Y
	pubKey := append(pubKeyOrig.X.Bytes(), pubKeyOrig.Y.Bytes()...)
	return &Wallet{
		privateKey,
		pubKey,
	}
}

//生成地址
func (w *Wallet) NewAddress() string {
	pubKey := w.Public

	ripHaValue := ripe160HashValue(pubKey)

	version := byte(00)
	payload := append([]byte{version}, ripHaValue...)

	checkCode := checkSum(payload)
	//25字节数据
	payload = append(payload, checkCode...)
	myAlphabet := base58.BitcoinAlphabet
	address := base58.Encode(payload, myAlphabet)

	return address
}
func ripe160HashValue(data []byte) []byte {
	hash := sha256.Sum256(data)
	//编码器
	ripe160hasher := ripemd160.New()

	_, err := ripe160hasher.Write(hash[:])
	if err != nil {
		log.Panic(err)
	}
	//返回ripe160的哈希结果
	ripeHashValue := ripe160hasher.Sum(nil)
	return ripeHashValue
}
func checkSum(data []byte) []byte {
	//检查Sum
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	//前四字节校验码
	checkCode := hash2[:4]
	return checkCode
}