package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"time"
)

type config struct {
	OsGame OsGameCfg `toml:"osgame"`
}

type OsGameCfg struct {
	Language    []string   `toml:"language"`     //解析到的语言配置
	OsGameTypes [][]string `toml:"osgame_types"` // 游戏大小模式
}

func main() {
	file := "test.toml"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	var cfg config
	_, err = toml.Decode(string(content), &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("cfg=%+v", cfg)
	fmt.Println(time.Now().Unix())
}
