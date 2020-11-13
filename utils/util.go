package utils

import "bytes"

//该函数将对明文进行尾部填充，采用PKCS5方式
func PKCS5EndPadding(text []byte,size int) []byte {
	//如果是刚好是整租，则不用进行尾部添加
	/*if  len(text)%size == 1{
		return text
	}*/
	paddingSize := size - len(text)%size
	paddingText := bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	return append(text,paddingText...)
}

//该函数将对明文进行尾部填充，采用zero方式
func ZeroEndPadding(text []byte,size int) []byte {
	paddingSize := size - len(text)%size
	paddingText := bytes.Repeat([]byte{byte(0)},paddingSize)
	return append(text,paddingText...)
}

//该函数用于对解密后的数据进行尾部去除
func ClearPKCS5Padding(text []byte,size int) []byte {
	/*if  len(text)%size == 1{//明文刚刚充满一组，则不用去除
		return text
	}*/
	lastEle :=int(text [len(text) - 1])
	return text[0 :len(text) - lastEle]
}
