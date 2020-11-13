package main

import (
	_des "CryptCode/3des"
	"CryptCode/aes"
	"CryptCode/des"
	"fmt"
)

//https://github.com/ZJ-p007/Cryptology.git
func main() {
	//一、使用DES进行加解密
	key := []byte("20201112")
	//fmt.Println(key)
	data := "窗含西岭千秋雪，门泊东吴万里船"
	cipherText, err := des.DESEnCrypt([]byte(data), key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	originText, err := des.DESDeCrypt(cipherText, key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("DES解密后:", string(originText))
	fmt.Println("=====================================")

	//3DES加解密
	key1 := []byte("202011122020111220201112") //3des的密钥长度24字节
	//fmt.Println(key1)
	/*key2 := make([]byte,15)
	key2 = append(key2,[]byte("20201112")...)*/
	data1 := "穷在闹市无人问，富在深山有远亲"
	cipherText1, err := _des.TripleDesEncrypt([]byte(data1), key1)
	if err != nil {
		fmt.Println("3DES加密失败", err.Error())
		return
	}
	originalText1, err := _des.TripeDesDecrypt(cipherText1, key1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("3DES解密后:", string(originalText1))

	fmt.Println("=====================================")
	//AES加解密
	key2 := []byte("2020111220201112")
	data2 := "人生在世不称意，明诏散发弄扁舟"
	cipherText2, err := aes.AessEnCrypt([]byte(data2), key2)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(cipherText2)

	//AES解密
	originalText2,err := aes.AessDecrypt(cipherText2,key2)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("AES解密后:",string(originalText2))

}
