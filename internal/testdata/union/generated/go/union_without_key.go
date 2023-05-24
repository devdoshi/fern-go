package api

import (
	json "encoding/json"
	fmt "fmt"
)

type UnionWithoutKey struct {
	Type string
	Foo  *Foo
	Bar  *Bar
}

func (x *UnionWithoutKey) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	x.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "foo":
		value := new(Foo)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Foo = value
	case "bar":
		value := new(Bar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Bar = value
	}
	return nil
}

type UnionWithoutKeyVisitor interface {
	VisitFoo(*Foo) error
	VisitBar(*Bar) error
}

func (x *UnionWithoutKey) Accept(v UnionWithoutKeyVisitor) error {
	switch x.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "foo":
		return v.VisitFoo(x.Foo)
	case "bar":
		return v.VisitBar(x.Bar)
	}
}