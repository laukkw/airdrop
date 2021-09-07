package service

import (
	"bufio"
	"github.com/rzry/airdrop/pkg/log"
	"go.uber.org/zap"
	"os"
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
