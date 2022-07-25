package main

import (
	"encoding/json"
	"fmt"
	"go-lib/node/account"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
)

type EIP712Domain struct {
	ChainId uint   `json:"chainId"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type EIP712Property struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type EIP712 struct {
	Domain      EIP712Domain `json:"domain"`
	Message     interface{}  `json:"message"`
	PrimaryType string       `json:"primaryType"` //MessagePrimary
	Types       struct {
		EIP712Domain []EIP712Property `json:"EIP712Domain"`
		Primary      []EIP712Property `json:"Primary"`
	} `json:"types"`
}

func ValidEIP712() {
	originData := EIP712{
		Domain: EIP712Domain{56, "Butterfly", "v1.0"},
		Message: struct {
			Content string `json:"content"`
		}{"Bind User"},
		PrimaryType: "Primary",
	}
	originData.Types.EIP712Domain = []EIP712Property{
		{"name", "string"},
		{"version", "string"},
		{"chainId", "uint256"},
	}
	originData.Types.Primary = []EIP712Property{
		{"content", "string"},
	}
	jsonData, _ := json.Marshal(originData)
	signData, _ := account.Signature("", string(jsonData))
	fmt.Println(signData)
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(jsonData), &typedData); err != nil {
		panic(err)
	}
	valid, err := sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress("0xaC39b311DCEb2A4b2f5d8461c1cdaF756F4F7Ae9"),
		typedData,
		"0xee0d9f9e63fa7183bea2ca2e614cf539464a4c120c8dfc1d5ccc367f242a2c5939d7f59ec2ab413b8a9047de5de2f1e5e97da4eba2ef0d6a89136464f992dae11c",
	)
	fmt.Println(valid, err) // true <nil>
}
