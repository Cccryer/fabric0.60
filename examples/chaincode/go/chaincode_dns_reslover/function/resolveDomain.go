package function

import (
	"errors"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/myutils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ResolveDomain 通过域名获取权威域名服务器地址
func ResolveDomain(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("incorrect number of arguments, Expecting 1 args")
	}
	domain := args[0]
	isValidDomain := myutils.CheckValidDomain(domain)
	if !isValidDomain {
		return nil, errors.New("input is invalid domain")
	}
	topLevelDomain, err := myutils.GetTopLevelDomain(domain)
	if err != nil {
		return nil, errors.New("failed to get top level domain")

	}
	authorityServer, err := stub.GetState(topLevelDomain)
	if err != nil {
		jsonResp := "{\"Error\":\"failed to resolver domain list by owner" + domain + "\"}"
		return nil, errors.New(jsonResp)
	}
	return authorityServer, nil
}
