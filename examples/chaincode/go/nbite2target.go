package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"math/big"
)

func nbits2target(nBits uint32) *big.Int {
	exponent := nBits >> 24
	mantissa := nBits & 0x007fffff

	var rtn *big.Int

	if exponent <= 3 {
		mantissa >>= uint(8 * (3 - exponent))
		rtn = new(big.Int).SetUint64(uint64(mantissa))
	} else {
		rtn = new(big.Int).SetUint64(uint64(mantissa))
		rtn.Lsh(rtn, uint(8*(exponent-3)))
	}

	//判断负数和溢出
	//pfNegative := mantissa != 0 && (nBits&0x00800000) != 0
	//
	//pfOverflow := mantissa != 0 && ((exponent > 34) ||
	//	(mantissa > 0xff && exponent > 33) ||
	//	(mantissa > 0xffff && exponent > 32))

	return rtn
}

func nbits2targetStr(nBits uint32) string {
	target := nbits2target(nBits)
	targetStr := fmt.Sprintf("%064x", target)
	return "0x" + targetStr
}

func VerifyPowNonce(nonce uint32, data []byte, nbits uint32) bool {

	compact := fmt.Sprintf("%d%s", nonce, data)
	hash := GetHash([]byte(compact))
	return CheckProofOfWork(hash, nbits)
}

func CheckProofOfWork(hash *big.Int, nbits uint32) bool {
	target := nbits2target(nbits)
	//if(pfNegative || target == 0 || pfOverflow || target > ?)
	result := hash.Cmp(target)
	if result < 1 {
		return true
	}
	return false
}

func GetHash(data []byte) *big.Int {
	hash1 := sha256.Sum256(data)
	hash := sha256.Sum256([]byte(hash1[:]))
	hash256 := new(big.Int)
	hash256.SetBytes(hash[:])

	hash256str := fmt.Sprintf("%064x", hash256)
	fmt.Printf("0x" + hash256str + "\n")
	return hash256
}

func GetNonce(nbits uint32, data []byte) uint32 {
	target := nbits2target(nbits)
	fmt.Printf("target = 0x" + fmt.Sprintf("%064x", target) + "\n")
	var nonce uint32
	nonce = 0
	compact := fmt.Sprintf("%d%s", nonce, data)
	for GetHash([]byte(compact)).Cmp(target) > 0 {
		fmt.Println(compact)
		nonce++
		compact = fmt.Sprintf("%d%s", nonce, data)
		if nonce > 200 { //真要挖矿？
			break
		}
	}
	return nonce
}

type ChaincodeInput struct {
	Args [][]byte `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

var chaincodeCtorJSON string
var testcmd = &cobra.Command{
	Use:   "test",
	Short: fmt.Sprintf("test"),
	Run: func(cmd *cobra.Command, args []string) {
		// 在这里访问chaincodeCtorJSON变量
		fmt.Println("Constructor message:", chaincodeCtorJSON)
		var input ChaincodeInput
		if err := json.Unmarshal([]byte(chaincodeCtorJSON), &input); err != nil {
			fmt.Printf("unmarshal error \n")
		}
		fmt.Println(input.Args)
	},
}

func Cmd() *cobra.Command {
	flags := testcmd.PersistentFlags()
	flags.StringVarP(&chaincodeCtorJSON, "ctor", "c", "{}",
		fmt.Sprintf("Constructor message for the %s in JSON format", "test"))
	return testcmd
}

func main() {
	data1 := "helloworld!"
	nonce := GetNonce(0x1d00ffff, []byte(data1))
	flag := VerifyPowNonce(nonce, []byte(data1), 0x1d00ffff)
	if flag {
		fmt.Println("yes!")
	}
	fmt.Println("==============================================================")
	//result := nbits2targetStr(0x1d00ffff)
	//fmt.Println(result)

	rootCmd := &cobra.Command{Use: "myprogram"}
	rootCmd.AddCommand(Cmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	//input := &pb.ChaincodeInput{}
	//if err := json.Unmarshal([]byte(chaincodeCtorJSON), &input); err != nil {
	//	fmt.Printf("unmarshal error \n")
	//}

}
