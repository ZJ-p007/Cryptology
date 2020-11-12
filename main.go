package main

import (
	"CryptCode/des"
	"fmt"
)
//https://github.com/ZJ-p007/Cryptology.git
func main() {
	//一、使用DES进行加解密
	key := []byte("20201112")
	data := "窗含西岭千秋雪，门泊东吴万里船"
	cipherText,err :=des.DESEnCrypt([]byte(data),key)
    if err !=nil{
    	fmt.Println(err.Error())
		return
	}

	originText ,err :=des.DESDeCrypt(cipherText,key)
    if err != nil{
    	fmt.Println(err.Error())
		return
	}
	fmt.Println("解密后:",string(originText))

}









