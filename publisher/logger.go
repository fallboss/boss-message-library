package publisher

import (
	"github.com/rendis/abslog"
	"sync"
)

var (
	logger *abslog.AbsLog
	once   sync.Once
)

func GetLogger() *abslog.AbsLog {
	once.Do(func() {
		logger = abslog.GetAbsLogBuilder().LogLevel(abslog.InfoLevel).Build()
	})
	return logger
}
