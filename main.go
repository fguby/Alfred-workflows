package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	ftp "github.com/dutchcoders/goftp"
	"github.com/go-ini/ini"
	"github.com/labstack/gommon/log"
)

/**
	使用多个goroutine来完成上传工作
**/

type Server struct {
	Server   string
	User     string
	Password string
}

var (
	fm  = fmt.Sprintf
	cfg *ini.File
	e   = func(err error) {
		if err != nil {
			fmt.Printf("error:%v", err)
		}
	}
	Setting = &Server{}
)

//读取配置文件
func init() {
	var err error
	cfg, err = ini.Load("./server.conf")
	if err != nil {
		fmt.Println("加载文件失败: %s", err)
		//非正常运行，退出整个程序
		os.Exit(1)
	}
	mapTo("server", Setting)
	if Setting.Server == "" || Setting.User == "" ||
		Setting.Password == "" {
		exec_shell("open ./server.conf")
		os.Exit(1)
	}
}

func main() {
	var err error
	var client *ftp.FTP

	if client, err = ftp.Connect(Setting.Server); err != nil {
		str := fmt.Sprintf("osascript -e 'display notification \"%s连接失败\" with title \"FTP\"'", Setting.Server)
		exec_shell(str)
		panic(err)
	}

	if err := client.Login(Setting.User, Setting.Password); err != nil {
		str := fmt.Sprintf("osascript -e 'display notification \"%s登陆失败\" with title \"FTP\"'", Setting.Server)
		exec_shell(str)
		panic(err)
	}

	defer client.Close()

	//命令行里获取路径参数
	//path := os.Args[1]
	path := os.Args[1]
	//path := "/Users/wushaoqiang/Downloads/videos"
	ok := IsFile(path)
	//如果是文件直接上传
	if ok {
		status := UploadFile(client, path)
		if status {
			//调用通知
			exec_shell("osascript -e 'display notification \"上传完成\" with title \"FTP\"'")
		}
	} else {
		//如果是目录，循环遍历上传
		files, err := ioutil.ReadDir(path)
		e(err)
		for _, file := range files {
			if !file.IsDir() {
				localFile := path + "/" + file.Name()
				UploadFile(client, localFile)
			}
		}
		str := fmt.Sprintf("osascript -e 'display notification \"%d个文件上传完成\" with title \"FTP\"'", len(files))
		exec_shell(str)
	}
}

func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

//上传文件至ftp服务器
func UploadFile(client *ftp.FTP, f string) bool {
	file, err := os.Open(f)
	if err != nil {
		e(err)
		return false
	}

	_, fileName := filepath.Split(f)

	if err := client.Stor(fileName, file); err != nil {
		e(err)
		return false
	}
	return true
}

func exec_shell(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Error(err.Error(), stderr.String())
	} else {
		log.Info(out.String())
	}
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
