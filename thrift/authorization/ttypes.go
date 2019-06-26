// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package authorization

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type GrantType int64

const (
	GrantType_DEVELOPER GrantType = 1
	GrantType_APP_ROOT  GrantType = 2
	GrantType_APP_USER  GrantType = 3
	GrantType_GUEST     GrantType = 4
)

func (p GrantType) String() string {
	switch p {
	case GrantType_DEVELOPER:
		return "GrantType_DEVELOPER"
	case GrantType_APP_ROOT:
		return "GrantType_APP_ROOT"
	case GrantType_APP_USER:
		return "GrantType_APP_USER"
	case GrantType_GUEST:
		return "GrantType_GUEST"
	}
	return "<UNSET>"
}

func GrantTypeFromString(s string) (GrantType, error) {
	switch s {
	case "GrantType_DEVELOPER":
		return GrantType_DEVELOPER, nil
	case "GrantType_APP_ROOT":
		return GrantType_APP_ROOT, nil
	case "GrantType_APP_USER":
		return GrantType_APP_USER, nil
	case "GrantType_GUEST":
		return GrantType_GUEST, nil
	}
	return GrantType(0), fmt.Errorf("not a valid GrantType string")
}

func GrantTypePtr(v GrantType) *GrantType { return &v }

type Grantee struct {
	TypeA1     *GrantType `thrift:"type,1" json:"type"`
	Identifier *string    `thrift:"identifier,2" json:"identifier"`
}

func NewGrantee() *Grantee {
	return &Grantee{}
}

var Grantee_TypeA1_DEFAULT GrantType

func (p *Grantee) GetTypeA1() GrantType {
	if !p.IsSetTypeA1() {
		return Grantee_TypeA1_DEFAULT
	}
	return *p.TypeA1
}

var Grantee_Identifier_DEFAULT string

func (p *Grantee) GetIdentifier() string {
	if !p.IsSetIdentifier() {
		return Grantee_Identifier_DEFAULT
	}
	return *p.Identifier
}
func (p *Grantee) IsSetTypeA1() bool {
	return p.TypeA1 != nil
}

func (p *Grantee) IsSetIdentifier() bool {
	return p.Identifier != nil
}

func (p *Grantee) Read(iprot thrift.TProtocol) error {
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

func (p *Grantee) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		temp := GrantType(v)
		p.TypeA1 = &temp
	}
	return nil
}

func (p *Grantee) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Identifier = &v
	}
	return nil
}

func (p *Grantee) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Grantee"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *Grantee) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetTypeA1() {
		if err := oprot.WriteFieldBegin("type", thrift.I32, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:type: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.TypeA1)); err != nil {
			return fmt.Errorf("%T.type (1) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:type: %s", p, err)
		}
	}
	return err
}

func (p *Grantee) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetIdentifier() {
		if err := oprot.WriteFieldBegin("identifier", thrift.STRING, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:identifier: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.Identifier)); err != nil {
			return fmt.Errorf("%T.identifier (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:identifier: %s", p, err)
		}
	}
	return err
}

func (p *Grantee) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Grantee(%+v)", *p)
}
