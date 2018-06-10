package main

import (
	"bytes"
	"fmt"
	sdkcom "github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology/account"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
	nstates "github.com/ontio/ontology/smartcontract/service/native/ont"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
	"github.com/ontio/ontology/smartcontract/states"
	vmtypes "github.com/ontio/ontology/smartcontract/types"
	"github.com/satori/go.uuid"
	"os"
)

const (
	TX_PER_WORKER_NUM = 2000
	WORKER_NUM        = 50
)

var (
	Admin *account.Account
)

func main() {
	defer os.RemoveAll("./Log")

	log.InitLog(log.InfoLog)
	var wallet, _ = account.Open("wallet-admin.dat")
	Admin, _ = wallet.GetDefaultAccount([]byte("passwordtest"))
	wallet, _ = account.Open("wallet-account.dat")
	var toAcc, _ = wallet.GetAccountByAddress("TA5SpWVGtvQbqnB2DMoV2woBq31mixCSFv", []byte("passwordtest"))
	endChannel := make(chan int)
	for j := 0; j < WORKER_NUM; j++ {
		go func() {
			for i := 0; i < TX_PER_WORKER_NUM; i++ {
				txHash, txContent := genTransfer(Admin, toAcc, 1)
				fmt.Print(txHash, ",", txContent, "\n")
			}
			endChannel <- 1
		}()
	}
	for j := 0; j < WORKER_NUM; j++ {
		<-endChannel
	}
}

func genTransfer(from, to *account.Account, value uint64) (string, string) {
	var sts []*nstates.State
	sts = append(sts, &nstates.State{
		From:  from.Address,
		To:    to.Address,
		Value: value,
	})
	transfers := &nstates.Transfers{
		States: sts,
	}
	bf := new(bytes.Buffer)

	if err := transfers.Serialize(bf); err != nil {
		fmt.Println("Serialize transfers struct error.")
		os.Exit(1)
	}

	cont := &states.Contract{
		Address: utils.OntContractAddress,
		Method:  "transfer",
		Args:    bf.Bytes(),
	}

	ff := new(bytes.Buffer)

	if err := cont.Serialize(ff); err != nil {
		fmt.Println("Serialize contract struct error.")
		os.Exit(1)
	}

	tx := sdkcom.NewInvokeTransaction(0, 0, vmtypes.Native, ff.Bytes())
	attribute := new(types.TxAttribute)
	uid, _ := uuid.NewV4()
	attribute.Data = make([]byte, len(uid))
	copy(attribute.Data[:], uid[:])
	tx.Attributes = append(tx.Attributes, attribute)
	err := sdkcom.SignTransaction(tx, from)
	if err != nil {
		fmt.Println("signTransaction error:", err)
		os.Exit(1)
	}

	txbf := new(bytes.Buffer)
	if err := tx.Serialize(txbf); err != nil {
		fmt.Println("Serialize transaction error.")
		os.Exit(1)
	}
	hash := tx.Hash()
	return common.ToHexString(hash.ToArray()), common.ToHexString(txbf.Bytes())
}
