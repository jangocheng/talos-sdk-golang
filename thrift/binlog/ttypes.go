// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package binlog

import (
	"bytes"
	"fmt"

	"github.com/XiaoMi/talos-sdk-golang/thrift/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type ColumnType int64

const (
	ColumnType_BOOLEAN           ColumnType = 0
	ColumnType_TINYINT           ColumnType = 1
	ColumnType_SMALLINT          ColumnType = 2
	ColumnType_MEDIUMINT         ColumnType = 3
	ColumnType_INT               ColumnType = 4
	ColumnType_BIGINT            ColumnType = 5
	ColumnType_REAL              ColumnType = 6
	ColumnType_DOUBLE            ColumnType = 7
	ColumnType_FLOAT             ColumnType = 8
	ColumnType_DECIMAL           ColumnType = 9
	ColumnType_NUMERIC           ColumnType = 10
	ColumnType_CHAR              ColumnType = 11
	ColumnType_VARCHAR           ColumnType = 12
	ColumnType_TIMESTAMP         ColumnType = 13
	ColumnType_UNSIGNEDINT       ColumnType = 14
	ColumnType_ENUM              ColumnType = 15
	ColumnType_UNSIGNEDTINYINT   ColumnType = 16
	ColumnType_UNSIGNEDSMALLINT  ColumnType = 17
	ColumnType_UNSIGNEDMEDIUMINT ColumnType = 18
	ColumnType_UNSIGNEDBIGINT    ColumnType = 19
	ColumnType_DATE              ColumnType = 20
	ColumnType_DATETIME          ColumnType = 21
	ColumnType_TIME              ColumnType = 22
	ColumnType_JSON              ColumnType = 23
	ColumnType_YEAR              ColumnType = 24
	ColumnType_SET               ColumnType = 25
)

func (p ColumnType) String() string {
	switch p {
	case ColumnType_BOOLEAN:
		return "ColumnType_BOOLEAN"
	case ColumnType_TINYINT:
		return "ColumnType_TINYINT"
	case ColumnType_SMALLINT:
		return "ColumnType_SMALLINT"
	case ColumnType_MEDIUMINT:
		return "ColumnType_MEDIUMINT"
	case ColumnType_INT:
		return "ColumnType_INT"
	case ColumnType_BIGINT:
		return "ColumnType_BIGINT"
	case ColumnType_REAL:
		return "ColumnType_REAL"
	case ColumnType_DOUBLE:
		return "ColumnType_DOUBLE"
	case ColumnType_FLOAT:
		return "ColumnType_FLOAT"
	case ColumnType_DECIMAL:
		return "ColumnType_DECIMAL"
	case ColumnType_NUMERIC:
		return "ColumnType_NUMERIC"
	case ColumnType_CHAR:
		return "ColumnType_CHAR"
	case ColumnType_VARCHAR:
		return "ColumnType_VARCHAR"
	case ColumnType_TIMESTAMP:
		return "ColumnType_TIMESTAMP"
	case ColumnType_UNSIGNEDINT:
		return "ColumnType_UNSIGNEDINT"
	case ColumnType_ENUM:
		return "ColumnType_ENUM"
	case ColumnType_UNSIGNEDTINYINT:
		return "ColumnType_UNSIGNEDTINYINT"
	case ColumnType_UNSIGNEDSMALLINT:
		return "ColumnType_UNSIGNEDSMALLINT"
	case ColumnType_UNSIGNEDMEDIUMINT:
		return "ColumnType_UNSIGNEDMEDIUMINT"
	case ColumnType_UNSIGNEDBIGINT:
		return "ColumnType_UNSIGNEDBIGINT"
	case ColumnType_DATE:
		return "ColumnType_DATE"
	case ColumnType_DATETIME:
		return "ColumnType_DATETIME"
	case ColumnType_TIME:
		return "ColumnType_TIME"
	case ColumnType_JSON:
		return "ColumnType_JSON"
	case ColumnType_YEAR:
		return "ColumnType_YEAR"
	case ColumnType_SET:
		return "ColumnType_SET"
	}
	return "<UNSET>"
}

func ColumnTypeFromString(s string) (ColumnType, error) {
	switch s {
	case "ColumnType_BOOLEAN":
		return ColumnType_BOOLEAN, nil
	case "ColumnType_TINYINT":
		return ColumnType_TINYINT, nil
	case "ColumnType_SMALLINT":
		return ColumnType_SMALLINT, nil
	case "ColumnType_MEDIUMINT":
		return ColumnType_MEDIUMINT, nil
	case "ColumnType_INT":
		return ColumnType_INT, nil
	case "ColumnType_BIGINT":
		return ColumnType_BIGINT, nil
	case "ColumnType_REAL":
		return ColumnType_REAL, nil
	case "ColumnType_DOUBLE":
		return ColumnType_DOUBLE, nil
	case "ColumnType_FLOAT":
		return ColumnType_FLOAT, nil
	case "ColumnType_DECIMAL":
		return ColumnType_DECIMAL, nil
	case "ColumnType_NUMERIC":
		return ColumnType_NUMERIC, nil
	case "ColumnType_CHAR":
		return ColumnType_CHAR, nil
	case "ColumnType_VARCHAR":
		return ColumnType_VARCHAR, nil
	case "ColumnType_TIMESTAMP":
		return ColumnType_TIMESTAMP, nil
	case "ColumnType_UNSIGNEDINT":
		return ColumnType_UNSIGNEDINT, nil
	case "ColumnType_ENUM":
		return ColumnType_ENUM, nil
	case "ColumnType_UNSIGNEDTINYINT":
		return ColumnType_UNSIGNEDTINYINT, nil
	case "ColumnType_UNSIGNEDSMALLINT":
		return ColumnType_UNSIGNEDSMALLINT, nil
	case "ColumnType_UNSIGNEDMEDIUMINT":
		return ColumnType_UNSIGNEDMEDIUMINT, nil
	case "ColumnType_UNSIGNEDBIGINT":
		return ColumnType_UNSIGNEDBIGINT, nil
	case "ColumnType_DATE":
		return ColumnType_DATE, nil
	case "ColumnType_DATETIME":
		return ColumnType_DATETIME, nil
	case "ColumnType_TIME":
		return ColumnType_TIME, nil
	case "ColumnType_JSON":
		return ColumnType_JSON, nil
	case "ColumnType_YEAR":
		return ColumnType_YEAR, nil
	case "ColumnType_SET":
		return ColumnType_SET, nil
	}
	return ColumnType(0), fmt.Errorf("not a valid ColumnType string")
}

func ColumnTypePtr(v ColumnType) *ColumnType { return &v }

type Operation int64

const (
	Operation_INSERT       Operation = 0
	Operation_UPDATE       Operation = 1
	Operation_DELETE       Operation = 2
	Operation_DUMP         Operation = 3
	Operation_ADD_COLUMNS  Operation = 4
	Operation_DROP_COLUMNS Operation = 5
)

func (p Operation) String() string {
	switch p {
	case Operation_INSERT:
		return "Operation_INSERT"
	case Operation_UPDATE:
		return "Operation_UPDATE"
	case Operation_DELETE:
		return "Operation_DELETE"
	case Operation_DUMP:
		return "Operation_DUMP"
	case Operation_ADD_COLUMNS:
		return "Operation_ADD_COLUMNS"
	case Operation_DROP_COLUMNS:
		return "Operation_DROP_COLUMNS"
	}
	return "<UNSET>"
}

func OperationFromString(s string) (Operation, error) {
	switch s {
	case "Operation_INSERT":
		return Operation_INSERT, nil
	case "Operation_UPDATE":
		return Operation_UPDATE, nil
	case "Operation_DELETE":
		return Operation_DELETE, nil
	case "Operation_DUMP":
		return Operation_DUMP, nil
	case "Operation_ADD_COLUMNS":
		return Operation_ADD_COLUMNS, nil
	case "Operation_DROP_COLUMNS":
		return Operation_DROP_COLUMNS, nil
	}
	return Operation(0), fmt.Errorf("not a valid Operation string")
}

func OperationPtr(v Operation) *Operation { return &v }

type Column struct {
	BoolValue   *bool    `thrift:"boolValue,1" json:"boolValue"`
	I32Value    *int32   `thrift:"i32Value,2" json:"i32Value"`
	I64Value    *int64   `thrift:"i64Value,3" json:"i64Value"`
	DoubleValue *float64 `thrift:"doubleValue,4" json:"doubleValue"`
	StringValue *string  `thrift:"stringValue,5" json:"stringValue"`
}

func NewColumn() *Column {
	return &Column{}
}

var Column_BoolValue_DEFAULT bool

func (p *Column) GetBoolValue() bool {
	if !p.IsSetBoolValue() {
		return Column_BoolValue_DEFAULT
	}
	return *p.BoolValue
}

var Column_I32Value_DEFAULT int32

func (p *Column) GetI32Value() int32 {
	if !p.IsSetI32Value() {
		return Column_I32Value_DEFAULT
	}
	return *p.I32Value
}

var Column_I64Value_DEFAULT int64

func (p *Column) GetI64Value() int64 {
	if !p.IsSetI64Value() {
		return Column_I64Value_DEFAULT
	}
	return *p.I64Value
}

var Column_DoubleValue_DEFAULT float64

func (p *Column) GetDoubleValue() float64 {
	if !p.IsSetDoubleValue() {
		return Column_DoubleValue_DEFAULT
	}
	return *p.DoubleValue
}

var Column_StringValue_DEFAULT string

func (p *Column) GetStringValue() string {
	if !p.IsSetStringValue() {
		return Column_StringValue_DEFAULT
	}
	return *p.StringValue
}
func (p *Column) IsSetBoolValue() bool {
	return p.BoolValue != nil
}

func (p *Column) IsSetI32Value() bool {
	return p.I32Value != nil
}

func (p *Column) IsSetI64Value() bool {
	return p.I64Value != nil
}

func (p *Column) IsSetDoubleValue() bool {
	return p.DoubleValue != nil
}

func (p *Column) IsSetStringValue() bool {
	return p.StringValue != nil
}

func (p *Column) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Column) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.BoolValue = &v
	}
	return nil
}

func (p *Column) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.I32Value = &v
	}
	return nil
}

func (p *Column) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.I64Value = &v
	}
	return nil
}

func (p *Column) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.DoubleValue = &v
	}
	return nil
}

func (p *Column) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.StringValue = &v
	}
	return nil
}

func (p *Column) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Column"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Column) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetBoolValue() {
		if err := oprot.WriteFieldBegin("boolValue", thrift.BOOL, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:boolValue: %s", p, err)
		}
		if err := oprot.WriteBool(bool(*p.BoolValue)); err != nil {
			return fmt.Errorf("%T.boolValue (1) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:boolValue: %s", p, err)
		}
	}
	return err
}

func (p *Column) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetI32Value() {
		if err := oprot.WriteFieldBegin("i32Value", thrift.I32, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:i32Value: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.I32Value)); err != nil {
			return fmt.Errorf("%T.i32Value (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:i32Value: %s", p, err)
		}
	}
	return err
}

func (p *Column) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetI64Value() {
		if err := oprot.WriteFieldBegin("i64Value", thrift.I64, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:i64Value: %s", p, err)
		}
		if err := oprot.WriteI64(int64(*p.I64Value)); err != nil {
			return fmt.Errorf("%T.i64Value (3) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:i64Value: %s", p, err)
		}
	}
	return err
}

func (p *Column) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetDoubleValue() {
		if err := oprot.WriteFieldBegin("doubleValue", thrift.DOUBLE, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:doubleValue: %s", p, err)
		}
		if err := oprot.WriteDouble(float64(*p.DoubleValue)); err != nil {
			return fmt.Errorf("%T.doubleValue (4) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:doubleValue: %s", p, err)
		}
	}
	return err
}

func (p *Column) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetStringValue() {
		if err := oprot.WriteFieldBegin("stringValue", thrift.STRING, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:stringValue: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.StringValue)); err != nil {
			return fmt.Errorf("%T.stringValue (5) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:stringValue: %s", p, err)
		}
	}
	return err
}

func (p *Column) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Column(%+v)", *p)
}

type ColumnSchema struct {
	Name         string      `thrift:"name,1,required" json:"name"`
	TypeA1       *ColumnType `thrift:"type,2" json:"type"`
	DefaultValue *Column     `thrift:"defaultValue,3" json:"defaultValue"`
}

func NewColumnSchema() *ColumnSchema {
	return &ColumnSchema{}
}

func (p *ColumnSchema) GetName() string {
	return p.Name
}

var ColumnSchema_TypeA1_DEFAULT ColumnType

func (p *ColumnSchema) GetTypeA1() ColumnType {
	if !p.IsSetTypeA1() {
		return ColumnSchema_TypeA1_DEFAULT
	}
	return *p.TypeA1
}

var ColumnSchema_DefaultValue_DEFAULT *Column

func (p *ColumnSchema) GetDefaultValue() *Column {
	if !p.IsSetDefaultValue() {
		return ColumnSchema_DefaultValue_DEFAULT
	}
	return p.DefaultValue
}
func (p *ColumnSchema) IsSetTypeA1() bool {
	return p.TypeA1 != nil
}

func (p *ColumnSchema) IsSetDefaultValue() bool {
	return p.DefaultValue != nil
}

func (p *ColumnSchema) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ColumnSchema) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *ColumnSchema) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		temp := ColumnType(v)
		p.TypeA1 = &temp
	}
	return nil
}

func (p *ColumnSchema) ReadField3(iprot thrift.TProtocol) error {
	p.DefaultValue = &Column{}
	if err := p.DefaultValue.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.DefaultValue, err)
	}
	return nil
}

func (p *ColumnSchema) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ColumnSchema"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ColumnSchema) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return fmt.Errorf("%T.name (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:name: %s", p, err)
	}
	return err
}

func (p *ColumnSchema) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetTypeA1() {
		if err := oprot.WriteFieldBegin("type", thrift.I32, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:type: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.TypeA1)); err != nil {
			return fmt.Errorf("%T.type (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:type: %s", p, err)
		}
	}
	return err
}

func (p *ColumnSchema) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetDefaultValue() {
		if err := oprot.WriteFieldBegin("defaultValue", thrift.STRUCT, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:defaultValue: %s", p, err)
		}
		if err := p.DefaultValue.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.DefaultValue, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:defaultValue: %s", p, err)
		}
	}
	return err
}

func (p *ColumnSchema) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ColumnSchema(%+v)", *p)
}

type Row struct {
	Row []*Column `thrift:"row,1,required" json:"row"`
}

func NewRow() *Row {
	return &Row{}
}

func (p *Row) GetRow() []*Column {
	return p.Row
}
func (p *Row) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Row) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*Column, 0, size)
	p.Row = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &Column{}
		if err := _elem0.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem0, err)
		}
		p.Row = append(p.Row, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Row) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Row"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Row) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.LIST, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:row: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Row)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.Row {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:row: %s", p, err)
	}
	return err
}

func (p *Row) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Row(%+v)", *p)
}

type Binlog struct {
	Timestamp     int64           `thrift:"timestamp,1,required" json:"timestamp"`
	OperationType Operation       `thrift:"operationType,2,required" json:"operationType"`
	Schema        []*ColumnSchema `thrift:"schema,3,required" json:"schema"`
	Rows          []*Row          `thrift:"rows,4" json:"rows"`
	Database      *string         `thrift:"database,5" json:"database"`
	Table         *string         `thrift:"table,6" json:"table"`
	PrimaryKey    []string        `thrift:"primaryKey,7" json:"primaryKey"`
	OldRows       []*Row          `thrift:"oldRows,8" json:"oldRows"`
}

func NewBinlog() *Binlog {
	return &Binlog{}
}

func (p *Binlog) GetTimestamp() int64 {
	return p.Timestamp
}

func (p *Binlog) GetOperationType() Operation {
	return p.OperationType
}

func (p *Binlog) GetSchema() []*ColumnSchema {
	return p.Schema
}

var Binlog_Rows_DEFAULT []*Row

func (p *Binlog) GetRows() []*Row {
	return p.Rows
}

var Binlog_Database_DEFAULT string

func (p *Binlog) GetDatabase() string {
	if !p.IsSetDatabase() {
		return Binlog_Database_DEFAULT
	}
	return *p.Database
}

var Binlog_Table_DEFAULT string

func (p *Binlog) GetTable() string {
	if !p.IsSetTable() {
		return Binlog_Table_DEFAULT
	}
	return *p.Table
}

var Binlog_PrimaryKey_DEFAULT []string

func (p *Binlog) GetPrimaryKey() []string {
	return p.PrimaryKey
}

var Binlog_OldRows_DEFAULT []*Row

func (p *Binlog) GetOldRows() []*Row {
	return p.OldRows
}
func (p *Binlog) IsSetRows() bool {
	return p.Rows != nil
}

func (p *Binlog) IsSetDatabase() bool {
	return p.Database != nil
}

func (p *Binlog) IsSetTable() bool {
	return p.Table != nil
}

func (p *Binlog) IsSetPrimaryKey() bool {
	return p.PrimaryKey != nil
}

func (p *Binlog) IsSetOldRows() bool {
	return p.OldRows != nil
}

func (p *Binlog) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.ReadField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.ReadField8(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Binlog) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Timestamp = v
	}
	return nil
}

func (p *Binlog) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		temp := Operation(v)
		p.OperationType = temp
	}
	return nil
}

func (p *Binlog) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*ColumnSchema, 0, size)
	p.Schema = tSlice
	for i := 0; i < size; i++ {
		_elem1 := &ColumnSchema{}
		if err := _elem1.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem1, err)
		}
		p.Schema = append(p.Schema, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Binlog) ReadField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*Row, 0, size)
	p.Rows = tSlice
	for i := 0; i < size; i++ {
		_elem2 := &Row{}
		if err := _elem2.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem2, err)
		}
		p.Rows = append(p.Rows, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Binlog) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.Database = &v
	}
	return nil
}

func (p *Binlog) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 6: %s", err)
	} else {
		p.Table = &v
	}
	return nil
}

func (p *Binlog) ReadField7(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]string, 0, size)
	p.PrimaryKey = tSlice
	for i := 0; i < size; i++ {
		var _elem3 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_elem3 = v
		}
		p.PrimaryKey = append(p.PrimaryKey, _elem3)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Binlog) ReadField8(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*Row, 0, size)
	p.OldRows = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &Row{}
		if err := _elem4.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem4, err)
		}
		p.OldRows = append(p.OldRows, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Binlog) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Binlog"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Binlog) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:timestamp: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Timestamp)); err != nil {
		return fmt.Errorf("%T.timestamp (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:timestamp: %s", p, err)
	}
	return err
}

func (p *Binlog) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("operationType", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:operationType: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.OperationType)); err != nil {
		return fmt.Errorf("%T.operationType (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:operationType: %s", p, err)
	}
	return err
}

func (p *Binlog) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("schema", thrift.LIST, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:schema: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Schema)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.Schema {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:schema: %s", p, err)
	}
	return err
}

func (p *Binlog) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetRows() {
		if err := oprot.WriteFieldBegin("rows", thrift.LIST, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:rows: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Rows)); err != nil {
			return fmt.Errorf("error writing list begin: %s", err)
		}
		for _, v := range p.Rows {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v, err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:rows: %s", p, err)
		}
	}
	return err
}

func (p *Binlog) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetDatabase() {
		if err := oprot.WriteFieldBegin("database", thrift.STRING, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:database: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.Database)); err != nil {
			return fmt.Errorf("%T.database (5) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:database: %s", p, err)
		}
	}
	return err
}

func (p *Binlog) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetTable() {
		if err := oprot.WriteFieldBegin("table", thrift.STRING, 6); err != nil {
			return fmt.Errorf("%T write field begin error 6:table: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.Table)); err != nil {
			return fmt.Errorf("%T.table (6) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 6:table: %s", p, err)
		}
	}
	return err
}

func (p *Binlog) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetPrimaryKey() {
		if err := oprot.WriteFieldBegin("primaryKey", thrift.LIST, 7); err != nil {
			return fmt.Errorf("%T write field begin error 7:primaryKey: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.PrimaryKey)); err != nil {
			return fmt.Errorf("error writing list begin: %s", err)
		}
		for _, v := range p.PrimaryKey {
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 7:primaryKey: %s", p, err)
		}
	}
	return err
}

func (p *Binlog) writeField8(oprot thrift.TProtocol) (err error) {
	if p.IsSetOldRows() {
		if err := oprot.WriteFieldBegin("oldRows", thrift.LIST, 8); err != nil {
			return fmt.Errorf("%T write field begin error 8:oldRows: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.OldRows)); err != nil {
			return fmt.Errorf("error writing list begin: %s", err)
		}
		for _, v := range p.OldRows {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v, err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 8:oldRows: %s", p, err)
		}
	}
	return err
}

func (p *Binlog) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Binlog(%+v)", *p)
}
