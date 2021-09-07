package airdrop_server

import (
	"github.com/rzry/airdrop/internal/airdrop_server/service"
	"testing"
)

func TestAirdrop(t *testing.T) {
	r := service.ReadFile("./1.txt")
	t.Log(r[1])
}
