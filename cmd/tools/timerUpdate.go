package tools

import "boot/gen/log"

func NewTimerTool() TimerTool {
	return TimerTool{}
}

type TimerTool struct {
	logger *log.Logger
}
