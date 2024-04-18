package function

import (
	"errors"
	"fabric/fabric-pbft-domain-system/examples/chaincode/go/chaincode_dns_reslover/myutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Delete 删除某一个权威域名映射
func DeleteDomain(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("incorrect number of arguments. Expecting 1")
	}

	requestDomain := args[0]
	topLevelDomain, err := myutils.GetTopLevelDomain(requestDomain)
	if err != nil {
		return nil, errors.New("failed to get top level domain")
	}
	// Delete the key from the state in ledger
	if err := stub.DelState(topLevelDomain); err != nil {
		return nil, errors.New("failed to delete state")
	}
	return nil, nil
}
