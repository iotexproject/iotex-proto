# iotex-proto
Protobuf and utility package for IoTeX blockchain transaction and gRPC API

- `\proto` includes protobuf definition for all core data objects and gRPC API used by IoTeX blockchain

- `\golang` includes the generated protobuf files for go language

# Getting Started
## Installing
Install the Google protocol buffers compiler `protoc` v3.0.0 or above from https://github.com/protocolbuffers/protobuf/releases

Enable go mod. Install grpc-gateway https://github.com/grpc-ecosystem/grpc-gateway. Basically this is what you need:

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Compiling
```
make gogen
```
This generates the protobuf files and put into \golang directory

## Sign IoTeX blockchain transaction
secp256k1 ECDSA algorithm is used by IoTeX blockchain to sign and verify transaction. The signature of an IoTeX transaction is computed as the secp256k1 signature of hash of raw transaction
```
signature = secp256k1.Sign(hash of raw transaction)
```
The signature is in 65-byte [R, S, V] format where the last byte V is the recovery id for public key recovery

The following guide used sender address `io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m` and recipient address `io187wzp08vnhjjpkydnr97qlh8kh0dpkkytfam8j` as example. Replace your actual address and recipient address when creating and signing the transaction

### Create raw transaction
1. construct a message Transfer as defined in `\proto\type\action.proto`, amount is in unit of 10^-18 IOTX token

For example, to transfer 1.2 IOTX token, set amount = “1200000000000000000”

Set recipient = "io187wzp08vnhjjpkydnr97qlh8kh0dpkkytfam8j"

payload = hex-bytes of message you want to attach to transaction, can be nil/NULL

2. construct a message ActionCore as defined in `\proto\type\action.proto`, with action = transfer message in 1

Set version = 1, gasLimit = 10000, gasPrice = 1000000000000, that is 0.000001 IOTX

For nonce, issue a gRPC request GetAccount(GetAccountRequest) as defined in `\proto\api\api.proto` use the value of "pendingNonce" field in the reply

### Sign raw transaction
1. serialize the ActionCore message using protobuf
```
bytes = proto.Serialize(ActionCore message above)
```
2. hash of raw transaction is computed as the 32-byte Keccak256 hash of the bytes
```
hash = Keccak256(bytes)
```
3. sign the hash using sender's private key
```
sig = secp256k1.Sign(hash)
```

### Send signed transaction to IoTeX blockchain
1. construct a message Action as defined in \proto\type\action.proto

Set action = ActionCore above, senderPubKey = bytes representation of sender's public key, signature = sig above

2. issue a gRPC request SendAction(SendActionRequest) to IoTeX blockchain endpoint

### Go exmaple
Here's an example in golang:

Return value of SendTransfer() is the hash of signed transaction, you can query it on https://www.iotexscan.io to confirm that it has been committed to IoTeX blockchain
```
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
```