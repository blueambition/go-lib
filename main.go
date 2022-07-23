package main

import (
	"fmt"
	dc "github.com/shopspring/decimal"
	"math/big"
)

func main() {
	fmt.Println("还没开始就结束了")
	num, _ := new(big.Int).SetString("10000", 10)
	returnNum, _ := dc.NewFromBigInt(num, 0).Float64()
	fmt.Println(returnNum)
}
