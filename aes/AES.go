package aes

import (
	"CryptCode/utils"
	"crypto/aes"
	"crypto/cipher"
)

func AesEnCrypt(origin []byte,key []byte) ([]byte,error) {
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