package pkg

import (
	"github.com/manifoldco/promptui"
	"os"
)

func YorN() bool {
	prompt := promptui.Select{
		Label: "!!!请查看终端参数,确认输入参数是否正确!!!,继续请选择yes",
		Items: []string{"yes", "no"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		os.Exit(3)
	}
	switch result {
	case "yes":
		return true
	default:
		return false

	}
}
