package logger

// Defer logger defer call
type Defer func()

type LogConfigure struct {
	Level      string  `json:"level"`
	RotateHour float64 `json:"rotateHour"`
	MaxHour    float64 `json:"maxHour"`
}
