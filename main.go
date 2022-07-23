package main

import (
	"fmt"
	"go-lib/node/account"
)

func main() {
	priKey := "e50bc4b30036a662addecfcb7a9fc4337430f758518f977dc5e354701a697f14"
	sigData, _ := account.Signature(priKey, "Hello world")
	addr, err := account.SignatureAccount("Hello world", sigData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(addr)
	flag, err := account.VerifySignature(priKey, "Hello world", sigData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(flag)
}
