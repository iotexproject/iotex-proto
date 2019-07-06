package main

import (
	"log"
	"os"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-antenna-go/account"
	"github.com/iotexproject/iotex-antenna-go/iotx"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/pkg/errors"
)

// IotxProxy represents a proxy of iotex blockchain
type IotxProxy struct {
	sender string
	*iotx.Iotx
}

// NewIoProxy creates a new iotex proxy
func NewIotxProxy(server, pk string) (*IotxProxy, error) {
	service, err := iotx.NewIotx(server, true)
	if err != nil {
		return nil, err
	}
	if len(pk) == 0 {
		return nil, errors.New("empty private key")
	}
	sk, err := crypto.HexStringToPrivateKey(pk)
	if err != nil {
		return nil, err
	}
	acc, err := account.PrivateKeyToAccount(sk)
	if err != nil {
		return nil, err
	}
	if err := service.Accounts.AddAccount(acc); err != nil {
		return nil, err
	}

	return &IotxProxy{
		sender: acc.Address(),
		Iotx:   service,
	}, nil
}

// SendTransfer sends a signed transfer to IoTeX blockchain
// returns the hash of the pending transfer
func (p *IotxProxy) SendTransfer(amount, recipient string) (string, error) {
	req := &iotx.TransferRequest{
		From:     p.sender,
		To:       recipient,
		Value:    amount,
		GasLimit: "20000",
		GasPrice: "1",
	}
	return p.Iotx.SendTransfer(req)
}

// CheckTx checks whether a Tx has been committed to IoTeX blockchain
func (p *IotxProxy) CheckTx(hash string) error {
	req := &iotexapi.GetReceiptByActionRequest{
		ActionHash: hash,
	}
	_, err := p.Iotx.GetReceiptByAction(req)
	return err
}

func main() {
	// import private key string
	pk := os.Getenv("PRIVATE_KEY")

	// create an IoTeX proxy
	iotex, err := NewIotxProxy("api.iotex.one:443", pk)
	if err != nil {
		log.Fatalln(err)
	}
	defer iotex.Close()

	// send 2 IOTX token to io14s0vgnj0pjnazu4hsqlksdk7slah9vcfscn9ks
	// note amount is in unit of 10^-18
	amount := "2000000000000000000"
	recipient := "io14s0vgnj0pjnazu4hsqlksdk7slah9vcfscn9ks"
	tsf, err := iotex.SendTransfer(amount, recipient)
	if err != nil {
		log.Fatalln(err)
	}

	// note that our blockchain has a block time of 10 seconds
	// it would be best to wait 15 seconds before verifying the transaction

	// check the transfer success or not
	if err := iotex.CheckTx(tsf); err != nil {
		log.Fatalln(err)
	}
}
