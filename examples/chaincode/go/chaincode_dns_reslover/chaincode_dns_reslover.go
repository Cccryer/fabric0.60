package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/functions"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/myutils"
	"strings"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct{}
type void struct{}
type ResponseCode int

// strategies 用于存储查询策略
var strategies map[string]func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error)

const SimpleDomainQuest = "resolve"
const SimpleDomainUpdate = "update"
const SimpleDomainDelete = "delete"

const TopLevelUpdate = "TopLevelUpdate"
const TopLevelQuest = "TopLevelQuest"
const TopLevelDelete = "TopLevelDelete"

func init() {
	strategies = make(map[string]func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error))

	// 非顶级域名操作
	strategies[SimpleDomainQuest] = functions.SimpleDomainResolve
	strategies[SimpleDomainDelete] = functions.SimpleDomainDelete
	strategies[SimpleDomainUpdate] = functions.SimpleDomainUpdate

	// 顶级域名操作
	strategies[TopLevelQuest] = functions.TopLevelDomainResolve
	strategies[TopLevelUpdate] = functions.TopLevelDomainUpdate
	strategies[TopLevelDelete] = functions.TopLevelDomainDelete

}

// Init init the domain-ip relation
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var topLevelDomain, serverIp string
	var err error
	for _, arg := range args {
		pairs := strings.Split(arg, ":")
		if len(pairs) != 3 && len(pairs) != 2 {
			return nil, errors.New("incorrect number of arguments. Expecting 2")
		}
		topLevelDomain = pairs[0]
		serverIp = pairs[1]
		if len(pairs) == 2 {
			serverIp = pairs[1] + ":53"
		} else {
			serverIp = pairs[1] + ":" + pairs[2]
		}
		if topLevelDomain == "" || !myutils.CheckValidIp(serverIp) {
			return nil, errors.New("input is invalid domain or ip")
		}
		// Write the state to the ledger
		//record := myutils.BuildNewRecord(myutils.A, serverIp, 60, 0)
		//recordJson, _ := json.Marshal(record)
		err = stub.PutState(topLevelDomain, []byte(serverIp))
		if err != nil {
			return nil, errors.New("failed to update the domain-owner relation")
		}
	}
	return myutils.BuildResponse(true, "", nil), nil
}

// Invoke update the domain-ip relation
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, functionName string, args []string) ([]byte, error) {
	if len(functionName) != 0 && strategies[functionName] != nil {
		return strategies[functionName](stub, args)
	}
	return myutils.BuildWrongResponse("unknown functions name"), nil
}

// Query query the domain-ip relation
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(function) != 0 && strategies[function] != nil {
		return strategies[function](stub, args)
	}
	return myutils.BuildWrongResponse("unknown functions name"), nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
