package common

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

const TABLE_NAME = "dns_record"
const DEAFULT_TTL = 86400
const DEAULT_TYPE = "A"
const DEAULT_OWNER = "admin"

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
	record.RecordName = row.Columns[0].GetString_()
	record.RecordValue = row.Columns[1].GetString_()
	record.RecordType = row.Columns[2].GetString_()
	record.RecordOwner = row.Columns[3].GetString_()
	record.RecordTTL = row.Columns[4].GetInt32()
	record.CreateAt = row.Columns[5].GetUint64()
	record.UpdateAt = row.Columns[6].GetUint64()
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
		UpdateAt:    uint64(time.Now().Unix()),
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
