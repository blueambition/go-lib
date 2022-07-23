package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	dc "github.com/shopspring/decimal"
	"go-lib/node/custom"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

//创建账户
func Create(path string, password string) string {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return ""
	}
	return account.Address.Hex()
}

//获取私钥
func GetPriKeyByFile(keyFile, password string) (string, error) {
	keyJson, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return "", err
	}
	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		return "", err
	}
	privateKey := fmt.Sprintf("%x", key.PrivateKey.D.Bytes())
	return privateKey, nil
}

//获取私钥
//ksPath keystore文件夹目录地址
//address eth address
//password 密码
//返回私钥
func GetPriKey(ksPath, address, password string) (string, error) {
	var (
		keyFile  string
		files    []os.FileInfo
		err      error
		fileName string
		flag     bool
	)
	files, err = ioutil.ReadDir(ksPath)
	if err != nil || len(files) == 0 {
		return "", errors.New("此目录不存在文件")
	}
	address = strings.Replace(address, "0x", "", 1)
	address = strings.ToLower(address)
	for _, v := range files {
		fileName = strings.ToLower(v.Name())
		//fmt.Println("文件名："+fileName,"keyfile:",keyFile,"address:",address)
		flag = strings.Contains(fileName, address)
		if flag {
			keyFile = ksPath + "/" + v.Name()
			break
		}
	}
	if keyFile == "" {
		return "", errors.New("未找到账号keystore文件")
	}
	keyJson, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return "", errors.New("keyJson有误：" + err.Error() + "文件路径：" + keyFile + "keyJson :" + string(keyJson))
	}
	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		return "", errors.New("不能解锁文件：" + err.Error() + "文件路径：" + keyFile)
	}
	privateKey := fmt.Sprintf("%x", key.PrivateKey.D.Bytes())
	return privateKey, nil
}

//导入私钥
func ImportPriKey(ksPath, pkStr, password string) (string, error) {
	ks := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	privateKey, err := crypto.HexToECDSA(pkStr)
	if err != nil {
		return "", err
	}
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}

//获取地址对应私钥文件
//ksPath keystore文件夹目录地址
//address eth address
//password 密码
//返回KeyJson
func KeyStoreJson(ksPath, address, password string) string {
	var (
		path    string //文件路径
		flag    bool
		account accounts.Account
	)
	ks := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	address = strings.Replace(address, "0x", "", 1)
	address = strings.ToLower(address)
	list := ks.Accounts()
	for _, v := range list {
		path = v.URL.String()
		path = strings.ToLower(path)
		flag = strings.Contains(path, address)
		if flag {
			account = v
			break
		}
	}
	if account.Address.String() != "" {
		keyJSON, err := ks.Export(account, password, password)
		if err == nil {
			return string(keyJSON)
		}
	}
	return ""
}

//平台币余额
func Balance(client *ethclient.Client, address string) float64 {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return 0
	}
	amountDC, _ := dc.NewFromString(balance.String())
	ratioDC := dc.NewFromFloat(math.Pow(10, 18))
	realAmount, _ := amountDC.Div(ratioDC).RoundFloor(6).Float64()
	return realAmount
}

//平台币是否存在Pending交易
func PendingBalance(client *ethclient.Client, address string) float64 {
	account := common.HexToAddress(address)
	balance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		return 0
	}
	ratio := math.Pow(10, 18)
	amount, err := strconv.ParseFloat(balance.String(), 64)
	if err != nil {
		return 0
	}
	totalAmount := amount / ratio
	latestAmount := Balance(client, address)
	return totalAmount - latestAmount
}

//判断是否是以太坊地址
func IsErc20Address(address string) bool {
	if len(address) != 42 {
		return false
	}
	if address[0:2] != "0x" {
		return false
	}
	return true
}

//认证信息
func AuthOpts(client *ethclient.Client, authData custom.AuthData) (*bind.TransactOpts, error) {
	var (
		gasPrice *big.Int
		chainId  *big.Int
		auth     *bind.TransactOpts
		err      error
	)
	chainId, err = client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	privateKey := KeyStoreJson(authData.KsPath, authData.Account, authData.Password)
	auth, err = bind.NewTransactorWithChainID(strings.NewReader(privateKey), authData.Password, chainId)
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(authData.Account))
	if err != nil {
		return nil, err
	}
	auth.Nonce = new(big.Int).SetUint64(nonce)
	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	if authData.GasMultiple > 1 {
		gasPriceStr := gasPrice.String()
		gasPriceFloat, _ := strconv.ParseFloat(gasPriceStr, 64)
		gasPriceFloat = gasPriceFloat * authData.GasMultiple
		gasPriceStr = strconv.FormatFloat(gasPriceFloat, 'f', -1, 64)
		gasPrice, _ = new(big.Int).SetString(gasPriceStr, 10)
	}
	//设置成预估的限制
	auth.GasLimit = authData.GasLimit
	auth.GasPrice = gasPrice
	auth.Value = authData.ETHValue
	return auth, nil
}
