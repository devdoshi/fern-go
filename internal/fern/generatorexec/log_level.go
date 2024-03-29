// Generated by Fern. Do not edit.

package generatorexec

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type LogLevel uint8

const (
	LogLevelDebug LogLevel = iota + 1
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

func (x LogLevel) String() string {
	switch x {
	default:
		return strconv.Itoa(int(x))
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	}
}

func (x LogLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", x.String())), nil
}

func (x *LogLevel) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "DEBUG":
		value := LogLevelDebug
		*x = value
	case "INFO":
		value := LogLevelInfo
		*x = value
	case "WARN":
		value := LogLevelWarn
		*x = value
	case "ERROR":
		value := LogLevelError
		*x = value
	}
	return nil
}
