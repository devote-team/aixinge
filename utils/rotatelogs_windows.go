package utils

import (
	"aixinge/global"
	zapRotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zapRotateLogs.New(
		path.Join(global.CONFIG.Zap.Director, "%Y-%m-%d.log"),
		zapRotateLogs.WithMaxAge(7*24*time.Hour),
		zapRotateLogs.WithRotationTime(24*time.Hour),
	)
	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
