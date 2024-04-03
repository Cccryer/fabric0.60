package main

import (
	"crypto/sha256"
	"fmt"
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

func VerifyPowNonce() bool {
	data1 := "helloworld!"
	hash := GetHash([]byte(data1))
	return CheckProofOfWork(hash, 0x1d00ffff)
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

func GetNonce(nbits uint32, data []byte) {
	target := nbits2target(nbits)
	var nonce uint32

}

func main() {
	VerifyPowNonce()
	result := nbits2targetStr(0x1d00ffff)
	fmt.Println(result)
}
