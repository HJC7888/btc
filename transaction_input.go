package main

//交易输入
type TXInput struct{
	Txid []byte
	Vout int
	Signature []byte
	PubKey []byte
}
