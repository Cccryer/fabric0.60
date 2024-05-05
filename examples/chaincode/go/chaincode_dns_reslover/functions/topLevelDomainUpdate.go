package functions

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/common"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/myutils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// TopLevelDomainUpdate 添加权威域名映射
func TopLevelDomainUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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
	if err != nil || len(topLevelDomain) == 0 {
		return nil, errors.New("failed to get top level domain")
	}

	// update the ledger
	record := common.TableRecord{
		RecordName:  topLevelDomain,
		RecordValue: authorityServer,
		RecordType:  common.DEAULT_TYPE,
		RecordOwner: common.DEAULT_OWNER,
		RecordTTL:   common.DEAFULT_TTL,
		CreateAt:    0,
		UpdateAt:    uint64(time.Now().Unix()),
	}
	if result, err = common.UpdateRecord(stub, record); err != nil {
		result = false
		return nil, fmt.Errorf("failed to update record, %v", err)
	}
	return myutils.BuildResponse(result, "", nil), nil
}
