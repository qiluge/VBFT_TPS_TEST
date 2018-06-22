package main

import (
	"bytes"
	"fmt"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/rpc"
	"github.com/ontio/ontology/account"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"os"
	"strconv"
)

const ROUTING_NUM = 10

func main() {
	defer os.RemoveAll("./Log")

	log.InitLog(log.InfoLog)
	var wallet, _ = account.Open("wallet-admin.dat")
	admin, err := wallet.GetDefaultAccount([]byte("passwordtest"))
	if err != nil {
		fmt.Println("get admin err:", err)
	}
	wallet, _ = account.Open("wallet-account.dat")
	toAcc, err := wallet.GetAccountByAddress("AdTgdGPjahJjubZU19AwBu9F3oE4hncx4u", []byte("passwordtest"))
	if err != nil {
		fmt.Println("get account err:", err)
	}
	sdk := goSdk.NewOntologySdk()
	rpcClient := sdk.Rpc
	txNum, _ := strconv.Atoi(os.Args[1])
	txNum = txNum * 100000
	if txNum > 2<<32 {
		txNum = 2 << 32
	}
	exitChan := make(chan int)
	txNumPerRoutine := txNum / ROUTING_NUM
	for i := 0; i < ROUTING_NUM; i++ {
		go func(nonce uint32, routineIndex int) {
			for j := 0; j < txNumPerRoutine; j++ {
				txHash, txContent := genTransfer(admin, toAcc, 1, rpcClient, nonce)
				nonce++
				fmt.Print(txHash, ",", txContent, "\n")
			}
			exitChan <- 1
		}(uint32(txNumPerRoutine*i), i)
	}
	for i := 0; i < ROUTING_NUM; i++ {
		<-exitChan
	}
}

func genTransfer(from, to *account.Account, value uint64, rpc *rpc.RpcClient, nonce uint32) (string, string) {
	tx, err := rpc.NewTransferTransaction(0, 100000, "ont", from.Address, to.Address, value)
	if err != nil {
		return "", ""
	}
	tx.Nonce = nonce
	err = rpc.SignToTransaction(tx, from)
	if err != nil {
		return "", ""
	}

	txbf := new(bytes.Buffer)
	if err := tx.Serialize(txbf); err != nil {
		fmt.Println("Serialize transaction error.")
		os.Exit(1)
	}
	hash := tx.Hash()
	return hash.ToHexString(), common.ToHexString(txbf.Bytes())
}

