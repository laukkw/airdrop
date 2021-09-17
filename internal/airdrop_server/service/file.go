package service

import (
	"bufio"
	"github.com/rzry/airdrop/pkg/log"
	"go.uber.org/zap"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	todo := make([]string, 0)
	fp, err := os.Open(path)
	if err != nil {
		log.Error("open file err ", zap.Error(err))
		return nil
	}
	buf := bufio.NewScanner(fp)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		todo = append(todo, line)
	}
	return todo
}

func ReadAddFile(path string) []string {
	todo := make([]string, 0)
	fp, err := os.Open(path)
	if err != nil {
		log.Error("open file err ", zap.Error(err))
		return nil
	}
	buf := bufio.NewScanner(fp)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		if strings.Contains(line, "转账失败"){
			comma := strings.Index(line, "user\": \"")
			temp := strings.Index(line[comma:], "0x")
			if len(line[comma+temp:len(line)-2]) == 42{
				todo = append(todo, line[comma+temp:len(line)-2])
			}
		}
	}
	return todo
}