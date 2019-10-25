package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			//过滤指定格式
			//ok := strings.HasSuffix(fi.Name(),"")
			ok := strings.HasPrefix(fi.Name(),"README.")
			if !ok {
				fullName := pathname + "/" + fi.Name()
				s = append(s, fullName)
			}
		}
	}
	return s, nil
}

func main() {
	//遍历目标文件夹下所有的文件名
	var fileWithDir []string
	fileWithDir, _ = GetAllFile("./data", fileWithDir)
	fmt.Printf("list of all files: %v\n", fileWithDir)

	fileEn,err := ioutil.ReadFile(fileWithDir[0])
	fileZh,err := ioutil.ReadFile(fileWithDir[1])
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	dataEn := strings.Split(string(fileEn),"\n\n")
	dataZh := strings.Split(string(fileZh),"\n\n")

	//创建双语合成后的目标文件
	fileName := "testFile.rst"
	currFile,err:=os.Create(fileName)
	for i,paragEn := range dataEn{
		fmt.Println("Contents of file",+i,":", paragEn)
		fmt.Println("Contents of file",+i,":", dataZh[i])
		currFile.WriteString(paragEn+"\n\n")
		currFile.WriteString(dataZh[i]+"\n\n")
	}
}