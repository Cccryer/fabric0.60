package common

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const TABLE_NAME = "dns_record"
const DEAFULT_TTL = 86400
const DEAULT_TYPE = "A"
const DEAULT_OWNER = "admin"
const TableUserPrefix = "USER_"

const (
	INIT_ADMIN_NAME       = "admin"
	INIT_ADMIN_CERTFICATE = "-----BEGIN CERTIFICATE-----\nMIICCjCCAZGgAwIBAgIQAIvYoCKznwvKPCDx8ZyvtTAKBggqhkjOPQQDAjAXMRUw\nEwYDVQQDDAx3d3cudGFucy5mdW4wHhcNMjQwNTA4MDYxNzU1WhcNMjUwNTA4MDYx\nNzU1WjAXMRUwEwYDVQQDDAx3d3cudGFucy5mdW4wdjAQBgcqhkjOPQIBBgUrgQQA\nIgNiAASJ8KkFtmJVeUu30qei2lV/6ouCvDmu3+2IQDvxQz+b4+uEfd1jvBE2mH77\noTl/Lg9WGMrAOD64EdXQOlqK7UL53XESimQur9UFHJEW4IMq48ZYIdIKX9I3dMzE\nMSinf4WjgaEwgZ4wHQYDVR0OBBYEFJLmcANsF95a7AsSY5C6LyT7/fLZMA4GA1Ud\nDwEB/wQEAwIBhjAPBgNVHRMBAf8EBTADAQH/MDsGA1UdJQQ0MDIGCCsGAQUFBwMC\nBggrBgEFBQcDAQYIKwYBBQUHAwMGCCsGAQUFBwMEBggrBgEFBQcDCDAfBgNVHSME\nGDAWgBSS5nADbBfeWuwLEmOQui8k+/3y2TAKBggqhkjOPQQDAgNnADBkAjAHF4X5\nrhX8g1ZQAmz9pheV5wFFpjZOfM2jS9SVDbjNEw9vOOki8DAM/ripZMuOiT8CME6G\nbykxYmJJb3Rf3O2YKEqYmgOPkL3f0stA36cWNzp4C2PJqaQU2/ic16ZM1c63mg==\n-----END CERTIFICATE-----"

	INIT_JIM_NAME       = "jim"
	INIT_JIM_CERTFICATE = "-----BEGIN CERTIFICATE-----\nMIICCzCCAZGgAwIBAgIQAOpb0QCV/y0qdDtDHZEE7zAKBggqhkjOPQQDAjAXMRUw\nEwYDVQQDDAx3d3cudGFucy5mdW4wHhcNMjQwNTA4MDczODU2WhcNMjUwNTA4MDcz\nODU2WjAXMRUwEwYDVQQDDAx3d3cudGFucy5mdW4wdjAQBgcqhkjOPQIBBgUrgQQA\nIgNiAATCRfmQst/g22wAuSpRI9SOeeIiSHm6yFS/++d1FKdPC9I1VF5U2qjzvm5k\nJNUDBr7QSHqIcrtnuiZB+4xfVR5wIkir7mGx8kDq6yqUatZJhyI1mBvszrPGMWdL\n10LhxzijgaEwgZ4wHQYDVR0OBBYEFGEEKfoi8WRktgpNQ+5ZW1yWej0SMA4GA1Ud\nDwEB/wQEAwIBhjAPBgNVHRMBAf8EBTADAQH/MDsGA1UdJQQ0MDIGCCsGAQUFBwMC\nBggrBgEFBQcDAQYIKwYBBQUHAwMGCCsGAQUFBwMEBggrBgEFBQcDCDAfBgNVHSME\nGDAWgBRhBCn6IvFkZLYKTUPuWVtclno9EjAKBggqhkjOPQQDAgNoADBlAjAcdM3n\nsALhS5ksNd9h/XVXNFrNcrR22OKq81YLh3OU2GdWzAzqt8XU6UJM/UpudWECMQDt\nU/WJhQvaVAMr8XUrxjKdUoNThMh3J/zEAp3CZyS2vFfJa8cJDzV8j3s8a//8eVk=\n-----END CERTIFICATE-----"
)

type TableRecord struct {
	RecordName  string `json:"record_name"`
	RecordValue string `json:"record_value"`
	RecordType  string `json:"record_type"`
	RecordOwner string `json:"record_owner"`
	RecordTTL   int32  `json:"record_ttl"`
	CreateAt    uint64 `json:"create_at"`
	UpdateAt    uint64 `json:"update_at"`
}

func BuildRecordFromRawRow(row shim.Row) TableRecord {
	var record TableRecord
	if row.GetColumns() == nil {
		return TableRecord{}
	}
	if row.GetColumns()[0] != nil {
		record.RecordName = row.Columns[0].GetString_()
	}
	if row.GetColumns()[1] != nil {
		record.RecordValue = row.Columns[1].GetString_()
	}
	if row.GetColumns()[2] != nil {
		record.RecordType = row.Columns[2].GetString_()
	}
	if row.GetColumns()[3] != nil {
		record.RecordOwner = row.Columns[3].GetString_()
	}
	if row.GetColumns()[4] != nil {
		record.RecordTTL = row.Columns[4].GetInt32()
	}
	if row.GetColumns()[5] != nil {
		record.CreateAt = row.Columns[5].GetUint64()
	}
	if row.GetColumns()[6] != nil {
		record.UpdateAt = row.Columns[6].GetUint64()
	}
	return record
}

func BuildRowFromRecord(record TableRecord) shim.Row {
	var columns []*shim.Column
	columns = append(columns, &shim.Column{Value: &shim.Column_String_{String_: record.RecordName}})
	columns = append(columns, &shim.Column{Value: &shim.Column_String_{String_: record.RecordValue}})
	columns = append(columns, &shim.Column{Value: &shim.Column_String_{String_: record.RecordType}})
	columns = append(columns, &shim.Column{Value: &shim.Column_String_{String_: record.RecordOwner}})
	columns = append(columns, &shim.Column{Value: &shim.Column_Int32{Int32: record.RecordTTL}})
	columns = append(columns, &shim.Column{Value: &shim.Column_Uint64{Uint64: record.CreateAt}})
	columns = append(columns, &shim.Column{Value: &shim.Column_Uint64{Uint64: record.UpdateAt}})
	return shim.Row{Columns: columns}
}

// InsertRecord Insert a record
func InsertRecord(stub shim.ChaincodeStubInterface, record TableRecord) (bool, error) {
	row := BuildRowFromRecord(record)
	return stub.InsertRow(TABLE_NAME, row)
}

// UpdateRecord Query a record, if not exist, insert it
func UpdateRecord(stub shim.ChaincodeStubInterface, record TableRecord) (bool, error) {
	oldRecord, err := GetRecordByKey(stub, record.RecordName)
	if err != nil {
		return false, err
	}

	newRecord := TableRecord{
		RecordName:  record.RecordName,
		RecordValue: record.RecordValue,
		RecordType:  record.RecordType,
		RecordOwner: record.RecordOwner,
		RecordTTL:   record.RecordTTL,
		CreateAt:    oldRecord.CreateAt,
		UpdateAt:    0,
	}
	if oldRecord == (TableRecord{}) {
		newRecord.CreateAt = newRecord.UpdateAt
		return stub.InsertRow(TABLE_NAME, BuildRowFromRecord(newRecord))
	}
	return stub.ReplaceRow(TABLE_NAME, BuildRowFromRecord(newRecord))
}

// DeleteRecord Delete a record
func DeleteRecord(stub shim.ChaincodeStubInterface, record TableRecord) error {
	row := BuildRowFromRecord(record)
	return stub.DeleteRow(TABLE_NAME, []shim.Column{*row.Columns[0]})
}

// CheckRecordExist Check if a record exist
func CheckRecordExist(stub shim.ChaincodeStubInterface, record TableRecord) (bool, error) {
	row, err := stub.GetRow(TABLE_NAME, []shim.Column{{Value: &shim.Column_String_{String_: record.RecordName}}})
	if err != nil {
		return false, err
	}
	return row.Columns != nil, nil
}

// GetRecordByKey Get a record
func GetRecordByKey(stub shim.ChaincodeStubInterface, key string) (TableRecord, error) {
	row, err := stub.GetRow(TABLE_NAME, []shim.Column{{Value: &shim.Column_String_{String_: key}}})
	if err != nil {
		return TableRecord{}, err
	}
	return BuildRecordFromRawRow(row), nil
}

// GetAllRecords Get all records
func GetAllRecords(stub shim.ChaincodeStubInterface, table string) ([]TableRecord, error) {
	var records []TableRecord
	rowChannel, err := stub.GetRows(table, []shim.Column{})
	if err != nil {
		return nil, fmt.Errorf("get rows failed: %v", err)
	}

	for row := range rowChannel {
		record := BuildRecordFromRawRow(row)
		records = append(records, record)
	}
	return records, nil
}
