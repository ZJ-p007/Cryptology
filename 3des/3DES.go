package _des

import (
	"CryptCode/utils"
	"crypto/cipher"
	"crypto/des"
)

//用于实现3DES算法加密
func TripleDesEncrypt(orgitext []byte,key []byte ) ([]byte,error) {
	//Key,data,mode
	//1.实例化一个Cipher
	block,err := des.NewTripleDESCipher(key)
	if err != nil{
		return nil,err
	}
	//对明文进行尾部填充

	//填充后的明文
    //cryptData := PKCS5EndPadding(orgitext,block.BlockSize())
	cryptData := utils.ZeroEndPadding(orgitext,block.BlockSize())
    //加密模式
    mode := cipher.NewCBCEncrypter(block,key)
    //对填充后的明文分组加密
    cipherText := make([]byte,len(cryptData))
    mode.CryptBlocks(cipherText,cryptData)
    return cipherText,nil
}


//该函数用于对密文进行解密，返回明文
func TripeDesDecrypt(ciphertext []byte,key []byte) ([]byte,error) {
//三元素：key，data,mode
	//1.实例化一个cipher
	block,err:=des.NewTripleDESCipher(key)
    if err != nil{
    	return nil,err
	}
	//不需要对密文进行填充，可直接使用，实例化一个mode
	blockMode :=cipher.NewCBCDecrypter(block,key)
	originText := make([]byte,len(ciphertext))//明文
	blockMode.CryptBlocks(originText,ciphertext)
	originText = utils.ClearPKCS5Padding(originText,block.BlockSize())
	return originText,nil
}

