package bees

import (
	"path"
	"fmt"
)

func GetFileSuffix(nameStr string) string {//获取文件后缀名
	var fileName string
	fileName = path.Base(nameStr) //获取文件名带后缀
	fmt.Println("fileName =", fileName)
	var fileSuffix string
	fileSuffix = path.Ext(fileName) //获取文件后缀
	return fileSuffix
}
