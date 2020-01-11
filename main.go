package main

import (
	"flag"
	"fmt"
	"github.com/guotie/config"
	"github.com/guotie/deferinit"
	"os"
	"runtime"
	"sync"
)

var (
	configFn      = flag.String("config", "./config.json", "config file path")
	masterDirPath string
	readLock      sync.WaitGroup //获取锁
	dirArray      []string       //整合所有文件夹
)

/**
服务运行
创建人:邵炜
创建时间:2017年2月8日18:01:18
输入参数:配置文件路径 是否为调试模式(d表示为调试模式,否则为正式运行模式)
*/
func serverRun(cfn string) {
	config.ReadCfg(cfn)
	configRead()

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("set many cpu successfully!")

	deferinit.InitAll()
	fmt.Println("init all module successfully!")

	deferinit.RunRoutines()
	fmt.Println("init all run successfully!")

}

/**
服务停止
创建人:邵炜
创建时间:2017年2月9日14:06:27
*/
func serverExit() {
	deferinit.StopRoutines()
	fmt.Println("stop routine successfully!")

	deferinit.FiniAll()
	fmt.Println("stop all modules successfully!")
}

func main() {
	serverRun(*configFn)
	readRootDir(masterDirPath)

	for index, dir := range dirArray {
		if index%10 == 0 {
			readLock.Wait()
		}
		readLock.Add(1)
		go syncExecCommand(dir)
	}
	readLock.Wait()
	//execCommand("go version")
	serverExit()
	os.Exit(0)
	fmt.Println("run over!")
}
