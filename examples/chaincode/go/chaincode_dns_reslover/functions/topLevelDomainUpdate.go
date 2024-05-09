package functions

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/common"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/myutils"
	"strconv"
)

const IsDebug = false

// TopLevelDomainUpdate 添加权威域名映射
func TopLevelDomainUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 6 {
		return nil, errors.New("incorrect number of arguments, Expecting 6 args")
	}
	result := true
	domain := args[0]
	authorityServer := args[1]
	recordType := args[2]
	recordTtl := args[3]
	owner := args[4]
	signature := args[5]

	isValidDomain := myutils.CheckValidDomain(domain)
	if !isValidDomain {
		return nil, errors.New("input is invalid domain")
	}

	topLevelDomain, err := myutils.GetTopLevelDomain(domain)
	if err != nil || len(topLevelDomain) == 0 {
		return nil, errors.New("failed to get top level domain")
	}

	if recordType != "A" && recordType != "NS" {
		recordType = common.DEAULT_TYPE
	}

	ttl, err := strconv.ParseInt(recordTtl, 10, 32)
	if err != nil {
		return nil, err
	}
	if ttl < 0 {
		ttl = common.DEAFULT_TTL
	}

	// check the content
	isValidContent, err := verifyContent(stub, owner, domain, signature)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if !isValidContent {
		return nil, fmt.Errorf("the content is invalid, please check it")
	}

	isSameUser, err := verifySameUser(stub, owner, topLevelDomain)
	if err != nil {
		return nil, fmt.Errorf("error occur when get the history record %+v", err)
	}
	if !isSameUser {
		return nil, fmt.Errorf("error occur when check the history record %s", owner)
	}

	// update the ledger
	record := common.TableRecord{
		RecordName:  topLevelDomain,
		RecordValue: authorityServer,
		RecordType:  recordType,
		RecordOwner: owner,
		RecordTTL:   int32(ttl),
		CreateAt:    0,
		UpdateAt:    0,
	}
	if result, err = common.UpdateRecord(stub, record); err != nil {
		result = false
		return nil, fmt.Errorf("failed to update record, %+v", err)
	}
	return myutils.BuildResponse(result, "", nil), nil
}

// verifyContent 验证用户身份以及传输内容
func verifyContent(stub shim.ChaincodeStubInterface, owner string, content string, hash string) (bool, error) {
	cert, err := stub.GetState(common.TableUserPrefix + owner)
	if err != nil {
		return false, fmt.Errorf("error %+v occur when get user, %s", err, owner)
	}
	if cert == nil {
		return false, fmt.Errorf("error occur when get user's certificate, %s", owner)
	}
	// 校验证书
	if !myutils.Verify(cert, content, hash) {
		return false, fmt.Errorf("error occur when auth content by certificate, %s", owner)
	}
	return true, nil
}

// verifySameUser 验证是否是同一个用户
func verifySameUser(stub shim.ChaincodeStubInterface, owner string, key string) (bool, error) {
	record, err := common.GetRecordByKey(stub, key)
	if err != nil {
		return false, err
	}
	if len(record.RecordOwner) != 0 && record.RecordOwner != owner {
		return false, fmt.Errorf("the domain's owner is not you %s, now is %s", owner, record.RecordOwner)
	}
	return true, nil
}
