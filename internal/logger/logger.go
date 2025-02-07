package logger

import (
	"go.uber.org/zap"
	"os"
)

func init() {
	os.MkdirAll("logs", os.ModePerm)
}

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "logs/app.log"} // Логи пишутся в stdout и файл app.log
	return config.Build()
}
