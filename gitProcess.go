package main

import (
	"bufio"
	"fmt"
	"github.com/swgloomy/gutil"
	"io"
	"os/exec"
	"strings"
)

//读取根目录下所有的文件夹
//create by gloomy 2018-2-25 15:27:22
func readRootDir(dirPath string) {
	rootDirIn, err := gutil.GetMyAllDirByDir(dirPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, value := range *rootDirIn {
		dirArray = append(dirArray, value)
		readRootDir(fmt.Sprintf("%s/%s", dirPath, value))
	}
}

func syncExecCommand(value string) {
	defer readLock.Done()
	contentArrayIn, err := execCommand("cmd", "/C", fmt.Sprintf("git -C %s pull", strings.Join([]string{masterDirPath, value}, "\\")))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, item := range *contentArrayIn {
			fmt.Println(item)
		}
	}
}

//执行函数命令 commandName命令 params命令参数
//create by gloomy 2018-2-25 15:36:26
func execCommand(commandName string, params ...string) (*[]string, error) {
	var resultContentArray []string
	cmd := exec.Command(commandName, params...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return &resultContentArray, err
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		resultContentArray = append(resultContentArray, line)
	}
	cmd.Wait()
	fmt.Println(resultContentArray)
	return &resultContentArray, nil
}
