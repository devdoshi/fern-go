package generatorexec

type LogUpdate struct {
	Level   LogLevel `json:"level"`
	Message string   `json:"message"`
}