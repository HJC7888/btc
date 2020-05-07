package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type TXOutput struct {
	Value int
	PubKeyHash []byte
}

//创建输出
func NewTXOutput(value int , address string) *TXOutput{
	out := &TXOutput{value,nil}
	out.Lock([]byte(address))
	return out
}

//锁定地址
func (out *TXOutput) Lock(address []byte){
	pubKey160 := Base58Decode(address)
	pubKey160 = pubKey160[1:len(pubKey160)-4]
	out.PubKeyHash = pubKey160
}

//解锁地址
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash,pubKeyHash) == 0
}

type TXOutputs struct{
	Outputs []TXOutput
}
//序列化交易输出集合
func (outs TXOutputs) Serialize() []byte{
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(outs)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
//反序列化交易输出集合
func UnSerializeTXOutputs(txOutputsBytes []byte) TXOutputs{
	var outputs TXOutputs
	decoder := gob.NewDecoder( bytes.NewReader(txOutputsBytes) )
	err := decoder.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}
	return outputs
}