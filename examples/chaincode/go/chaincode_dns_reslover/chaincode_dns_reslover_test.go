package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args []string) {
	_, err := stub.MockInit("1", "init", args)
	if err != nil {
		fmt.Println("Init failed", err)
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, function string, name string, value string) {
	bytes, err := stub.MockQuery(function, []string{name})
	if err != nil {
		fmt.Println("Query", name, "failed", err)
		t.FailNow()
	}
	if bytes == nil {
		fmt.Println("Query", name, "failed to get value")
		if value != "" {
			t.FailNow()
		}
	}
	if string(bytes) != value {
		fmt.Println("Query value", name, "was not", value, "as expected", "the actual value is", string(bytes))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, function string, args []string) {
	_, err := stub.MockInvoke("1", function, args)
	if err != nil {
		fmt.Println("Invoke", args, "failed", err)
		t.FailNow()
	}
}

func TestExample02_Init(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex10", scc)

	//checkInit(t, stub, []string{"A", "", "B", ""})
	checkInit(t, stub, []string{"com:1.1.1.1", "cn:1.1.1.2"})

	checkState(t, stub, "com", "1.1.1.1:53")
	checkState(t, stub, "cn", "1.1.1.2:53")
}

func TestExample02_Query(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex10", scc)

	// Init A have no domain, B have no domain
	checkInit(t, stub, []string{"com:1.1.1.1:53", "cn:1.1.1.2:53"})
	// Query A
	checkQuery(t, stub, "TopLevelQuest", "google.com", "1.1.1.1:53")
	checkQuery(t, stub, "TopLevelQuest", "google.cn", "1.1.1.2:53")

}

// TestExample03_UpdateSimpleDomain 测试插入删除
func TestExample04_UpdateSimpleDomain(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("simpleDomainQuest Test", scc)

	// Init A have no domain, B have no domain
	checkInit(t, stub, []string{"com:127.0.0.1:30054", "cn:127.0.0.1:30054"})
	// Query A
	checkInvoke(t, stub, "delete", []string{"xxx.google.com"})
	checkInvoke(t, stub, "update", []string{"xxx.google.com", "2.2.2.2"})
	checkInvoke(t, stub, "update", []string{"xxx.google.com", "2.2.2.3"})
	checkQuery(t, stub, "resolve", "xxx.google.com", "2.2.2.2,2.2.2.3")
}

func TestExample05_DeleteSimpleDomain(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("simpleDomainDelete Test", scc)

	// Init A have no domain, B have no domain
	checkInit(t, stub, []string{"com:127.0.0.1:30054", "cn:127.0.0.1:30054"})
	// Query A
	checkInvoke(t, stub, "update", []string{"xxx.google.com", "2.2.2.2"})
	checkInvoke(t, stub, "update", []string{"xxx.google.com", "2.2.2.3"})
	checkInvoke(t, stub, "delete", []string{"xxx.google.com"})
	checkInvoke(t, stub, "update", []string{"xxx.google.com", "2.2.2.4"})
	checkQuery(t, stub, "resolve", "xxx.google.com", "2.2.2.4")
}

func TestExample02_Invoke(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex10", scc)

	// A and B have no domain
	checkInit(t, stub, []string{"com:1.1.1.1", "cn:1.1.1.2"})

	// A get domain google.com
	checkInvoke(t, stub, "add", []string{"google.kr", "1.1.1.3"})
	checkInvoke(t, stub, "add", []string{"google.jp", "1.1.1.4"})
	checkQuery(t, stub, "resolveDomain", "google.kr", "1.1.1.3")
	checkQuery(t, stub, "resolveDomain", "google.jp", "1.1.1.4")
}
