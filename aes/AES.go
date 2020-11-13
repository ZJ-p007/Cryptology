package aes

import (
	"CryptCode/utils"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)
//https://github.com/ZJ-p007/Cryptology.git
func AessEnCrypt(origin []byte,key []byte) ([]byte,error) {
//3元素:key,data,mode
	block,err :=aes.NewCipher(key)
	if err != nil{
		return nil,err
	}
   //2.对明文数据进行尾部填充
   cryptData :=utils.PKCS5EndPadding(origin,block.BlockSize())
   //3.实例化mode
   blockMode :=cipher.NewCBCEncrypter(block,key)
   //4.加密
   ciperData := make([]byte,len(cryptData))
   blockMode.CryptBlocks(ciperData,cryptData)
   return cryptData,nil
}

//AES解密
func AessDecrypt(cipherBytes []byte,key []byte) ([]byte,error) {
	block,err :=aes.NewCipher(key)
	if err != nil{
		fmt.Println(err.Error())
	}
	blockMode :=cipher.NewCBCDecrypter(block,key)
    originBytes := make([]byte,len(cipherBytes))
    blockMode.CryptBlocks(originBytes,cipherBytes)
    originBytes = utils.PKCS5EndPadding(originBytes,block.BlockSize())
	//originBytes = utils.ClearPKCS5Padding(originBytes,block.BlockSize())
    return originBytes,nil

}
