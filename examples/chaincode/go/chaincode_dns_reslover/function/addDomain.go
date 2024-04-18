package function

import (
	"errors"
	"fabric/fabric-pbft-domain-system/examples/chaincode/go/chaincode_dns_reslover/myutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

// AddDomain 添加权威域名映射
func AddDomain(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("incorrect number of arguments, Expecting 2 args")
	}
	result := true
	domain := args[0]
	authorityServer := args[1]
	isValidDomain := myutils.CheckValidDomain(domain)
	if !isValidDomain {
		return nil, errors.New("input is invalid domain")
	}
	topLevelDomain, err := myutils.GetTopLevelDomain(domain)
	if err != nil {
		return nil, errors.New("failed to get top level domain")
	}
	if err := stub.PutState(topLevelDomain, []byte(authorityServer)); err != nil {
		result = false
		return nil, errors.New("failed to update the domain-owner relation")
	}
	return []byte(strconv.FormatBool(result)), nil
}
