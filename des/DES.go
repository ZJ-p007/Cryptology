package des

import (
	"CryptCode/utils"
	"crypto/cipher"
	"crypto/des"
)

//使用DES算法对明文进行加密
func DESEnCrypt(data []byte, key []byte) ([]byte,error) {
	//三要d素：key,data,mode
	block,err :=des.NewCipher(key)
	if err != nil{
		return nil,err
	}
	//对数据进行尾部填充
	originData :=utils.PKCS5EndPadding(data,block.BlockSize())//填充后明文
	blockMode :=cipher.NewCBCEncrypter(block,key)
	cipherData := make([]byte, len(originData))
    blockMode.CryptBlocks(cipherData,originData)
   return cipherData,nil
}

//使用DES算法对密文进行解密
func DESDeCrypt(data []byte, key []byte) ([]byte,error) {
	block,err :=des.NewCipher(key)
	if err != nil{
		return nil,err
	}
	//解密
	blockMode:=cipher.NewCBCDecrypter(block,key)
	originData := make([]byte,len(data))
	blockMode.CryptBlocks(originData, data)
	originData = utils.ClearPKCS5Padding(originData,block.BlockSize())
	return originData,nil
}
