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
	for j := 0; j < txNum; j++ {
		txHash, txContent := genTransfer(admin, toAcc, 1, rpcClient, uint32(j))
		fmt.Print(txHash, ",", txContent, "\n")
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

