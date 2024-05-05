package functions

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/common"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/myutils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// TopLevelDomainDelete deletes the top level domain from the state in ledger
func TopLevelDomainDelete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("incorrect number of arguments. Expecting 1")
	}

	requestDomain := args[0]
	topLevelDomain, err := myutils.GetTopLevelDomain(requestDomain)
	record := common.TableRecord{RecordName: topLevelDomain}
	if err != nil {
		return nil, errors.New("failed to get top level domain")
	}
	// Delete the key from the state in ledger
	if err := common.DeleteRecord(stub, record); err != nil {
		return nil, fmt.Errorf("failed to delete state + %v", err)
	}
	return myutils.BuildResponse(true, "", nil), nil
}
