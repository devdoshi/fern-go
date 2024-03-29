// Generated by Fern. Do not edit.

package generatorexec

import (
	json "encoding/json"
	fmt "fmt"
)

type GeneratorUpdate struct {
	Type             string
	Init             *InitUpdate
	InitV2           *InitUpdateV2
	Log              *LogUpdate
	Publishing       *PackageCoordinate
	Published        *PackageCoordinate
	ExitStatusUpdate *ExitStatusUpdate
}

func (x *GeneratorUpdate) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	x.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "init":
		value := new(InitUpdate)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Init = value
	case "initV2":
		value := new(InitUpdateV2)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.InitV2 = value
	case "log":
		value := new(LogUpdate)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Log = value
	case "publishing":
		var valueUnmarshaler struct {
			Publishing *PackageCoordinate `json:"publishing"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		x.Publishing = valueUnmarshaler.Publishing
	case "published":
		var valueUnmarshaler struct {
			Published *PackageCoordinate `json:"published"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		x.Published = valueUnmarshaler.Published
	case "exitStatusUpdate":
		var valueUnmarshaler struct {
			ExitStatusUpdate *ExitStatusUpdate `json:"exitStatusUpdate"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		x.ExitStatusUpdate = valueUnmarshaler.ExitStatusUpdate
	}
	return nil
}

func (x GeneratorUpdate) MarshalJSON() ([]byte, error) {
	switch x.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "init":
		var marshaler = struct {
			Type string `json:"_type"`
			*InitUpdate
		}{
			Type:       x.Type,
			InitUpdate: x.Init,
		}
		return json.Marshal(marshaler)
	case "initV2":
		var marshaler = struct {
			Type string `json:"_type"`
			*InitUpdateV2
		}{
			Type:         x.Type,
			InitUpdateV2: x.InitV2,
		}
		return json.Marshal(marshaler)
	case "log":
		var marshaler = struct {
			Type string `json:"_type"`
			*LogUpdate
		}{
			Type:      x.Type,
			LogUpdate: x.Log,
		}
		return json.Marshal(marshaler)
	case "publishing":
		var marshaler = struct {
			Type       string             `json:"_type"`
			Publishing *PackageCoordinate `json:"publishing"`
		}{
			Type:       x.Type,
			Publishing: x.Publishing,
		}
		return json.Marshal(marshaler)
	case "published":
		var marshaler = struct {
			Type      string             `json:"_type"`
			Published *PackageCoordinate `json:"published"`
		}{
			Type:      x.Type,
			Published: x.Published,
		}
		return json.Marshal(marshaler)
	case "exitStatusUpdate":
		var marshaler = struct {
			Type             string            `json:"_type"`
			ExitStatusUpdate *ExitStatusUpdate `json:"exitStatusUpdate"`
		}{
			Type:             x.Type,
			ExitStatusUpdate: x.ExitStatusUpdate,
		}
		return json.Marshal(marshaler)
	}
}

type GeneratorUpdateVisitor interface {
	VisitInit(*InitUpdate) error
	VisitInitV2(*InitUpdateV2) error
	VisitLog(*LogUpdate) error
	VisitPublishing(*PackageCoordinate) error
	VisitPublished(*PackageCoordinate) error
	VisitExitStatusUpdate(*ExitStatusUpdate) error
}

func (x *GeneratorUpdate) Accept(v GeneratorUpdateVisitor) error {
	switch x.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "init":
		return v.VisitInit(x.Init)
	case "initV2":
		return v.VisitInitV2(x.InitV2)
	case "log":
		return v.VisitLog(x.Log)
	case "publishing":
		return v.VisitPublishing(x.Publishing)
	case "published":
		return v.VisitPublished(x.Published)
	case "exitStatusUpdate":
		return v.VisitExitStatusUpdate(x.ExitStatusUpdate)
	}
}
