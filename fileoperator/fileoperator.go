package fileoperator

import (
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Write(fn string, content string) error {
	var err error
	var f *os.File

	if checkFileIsExist(fn) { //如果文件存在
		f, err = os.OpenFile(fn, os.O_WRONLY, 0666) //打开文件
	} else {
		f, err = os.Create(fn) //创建文件
	}
	defer f.Close()

	if err != nil {
		return err
	}

	check(err)
	_, err = io.WriteString(f, content) //写入文件(字符串)
	check(err)

	return err
}

func Read(fn string) (string, error) {
	var content string

	data, err := ioutil.ReadFile(fn)
	//check(err)
	if err != nil {
		return "", err
	}

	content = string(data)

	return content, err
}
