// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type ShapeType uint8

const (
	ShapeTypeEnum ShapeType = iota + 1
	ShapeTypeObject
	ShapeTypeUnion
	ShapeTypeUndiscriminatedUnion
)

func (x ShapeType) String() string {
	switch x {
	default:
		return strconv.Itoa(int(x))
	case ShapeTypeEnum:
		return "ENUM"
	case ShapeTypeObject:
		return "OBJECT"
	case ShapeTypeUnion:
		return "UNION"
	case ShapeTypeUndiscriminatedUnion:
		return "UNDISCRIMINATED_UNION"
	}
}

func (x ShapeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", x.String())), nil
}

func (x *ShapeType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "ENUM":
		value := ShapeTypeEnum
		*x = value
	case "OBJECT":
		value := ShapeTypeObject
		*x = value
	case "UNION":
		value := ShapeTypeUnion
		*x = value
	case "UNDISCRIMINATED_UNION":
		value := ShapeTypeUndiscriminatedUnion
		*x = value
	}
	return nil
}
