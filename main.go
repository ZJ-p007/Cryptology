package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	key := []byte("C1906041")
	data := "遇贵人先立业遇良人先成家"
	//1.得到cipher
	block,_:= des.NewCipher(key)
	//2.对数据明文进行结尾填充
	paddingData := EndPadding([]byte(data),block.BlockSize())
	//3.选择模式
	mode := cipher.NewCBCDecrypter(block,key)
	//4.加密
	dstData := make([]byte,len(paddingData))
	mode.CryptBlocks(dstData,paddingData)
	fmt.Println("加密后的内容:",dstData)

	/*解密
	DES三元素:key,data,mode
	 */
	key1 := []byte("c1906041")
	data1 := dstData//待解密的数据
	block1,err := des.NewCipher(key1)
	if err != nil{
		panic(err.Error())
	}
	deMode :=cipher.NewCBCDecrypter(block1,key1)
	orininalData := make([]byte,len(data1))
	//分组解密
	deMode.CryptBlocks(orininalData,data1)
    fmt.Println("解密后:",string(orininalData))



}

//明文尾部填充
func EndPadding(text []byte, blockSize int) []byte {
	//计算要填充的块内容的大小
	paddingSize := blockSize - len(text)%blockSize
	//Repeat平铺
    paddingText := bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	//fmt.Println("明文尾部追加",paddingText)
    return append(text,paddingText...)
}








