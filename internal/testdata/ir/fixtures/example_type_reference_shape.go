// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type ExampleTypeReferenceShape struct {
	Type      string
	Primitive *ExamplePrimitive
	Container *ExampleContainer
	Unknown   any
	Named     *ExampleNamedType
}

func (x *ExampleTypeReferenceShape) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	x.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "primitive":
		var valueUnmarshaler struct {
			Primitive *ExamplePrimitive `json:"primitive"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		x.Primitive = valueUnmarshaler.Primitive
	case "container":
		var valueUnmarshaler struct {
			Container *ExampleContainer `json:"container"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		x.Container = valueUnmarshaler.Container
	case "unknown":
		var valueUnmarshaler struct {
			Unknown any `json:"unknown"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		x.Unknown = valueUnmarshaler.Unknown
	case "named":
		value := new(ExampleNamedType)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Named = value
	}
	return nil
}

func (x ExampleTypeReferenceShape) MarshalJSON() ([]byte, error) {
	switch x.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "primitive":
		var marshaler = struct {
			Type      string            `json:"type"`
			Primitive *ExamplePrimitive `json:"primitive"`
		}{
			Type:      x.Type,
			Primitive: x.Primitive,
		}
		return json.Marshal(marshaler)
	case "container":
		var marshaler = struct {
			Type      string            `json:"type"`
			Container *ExampleContainer `json:"container"`
		}{
			Type:      x.Type,
			Container: x.Container,
		}
		return json.Marshal(marshaler)
	case "unknown":
		var marshaler = struct {
			Type    string `json:"type"`
			Unknown any    `json:"unknown"`
		}{
			Type:    x.Type,
			Unknown: x.Unknown,
		}
		return json.Marshal(marshaler)
	case "named":
		var marshaler = struct {
			Type string `json:"type"`
			*ExampleNamedType
		}{
			Type:             x.Type,
			ExampleNamedType: x.Named,
		}
		return json.Marshal(marshaler)
	}
}

type ExampleTypeReferenceShapeVisitor interface {
	VisitPrimitive(*ExamplePrimitive) error
	VisitContainer(*ExampleContainer) error
	VisitUnknown(any) error
	VisitNamed(*ExampleNamedType) error
}

func (x *ExampleTypeReferenceShape) Accept(v ExampleTypeReferenceShapeVisitor) error {
	switch x.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "primitive":
		return v.VisitPrimitive(x.Primitive)
	case "container":
		return v.VisitContainer(x.Container)
	case "unknown":
		return v.VisitUnknown(x.Unknown)
	case "named":
		return v.VisitNamed(x.Named)
	}
}