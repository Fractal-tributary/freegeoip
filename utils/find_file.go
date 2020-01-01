package utils

import (
	"os"
	"path/filepath"
	"regexp"
)

func FindFile(path, fileName string) (filePath string,err error) {
	err =filepath.Walk(path,func (path string,f os.FileInfo,err error) error {
		if (f==nil){return err}
		if f.IsDir(){return nil}
		//正则匹配路径名和需要查找的文件名
		ok,_:=regexp.MatchString(fileName,path)
		if ok{
			filePath = path
		}
		return nil
	})
	if err!=nil{
		return
	}
	return
}