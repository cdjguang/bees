package bees

import (
	"os"
	"fmt"
	"io"
	"crypto/md5"
)

func GetMd5Str(fileName string) (n string, err error) {//获取文件md5值
	var Md5Str string
	if srcFile, err := os.Open(fileName); err != nil {
		fmt.Printf("源文件(%v) 打开失败(%v)\n", fileName, err)
	} else {
		md5h := md5.New()
		_, err = io.Copy(md5h, srcFile)
		if err != nil {
			fmt.Println("Copy错误：", err)
			return
		}
		Md5Str = fmt.Sprintf("%x", md5h.Sum(nil))
	}
	return Md5Str,err
}
