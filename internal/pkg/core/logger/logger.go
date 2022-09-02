package logger

import (
	"os"
	"tasks/internal/config"

	gelf "github.com/snovichkov/zap-gelf"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	var (
		err  error
		host string
		core zapcore.Core
	)

	if host, err = os.Hostname(); err != nil {
		panic(err)
	}

	if core, err = gelf.NewCore(
		gelf.Addr(config.Config.Graylog.Gelf.Address),
		gelf.Host(host),
	); err != nil {
		panic(err)
	}

	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return core.Enabled(l)
		})),
	)
}
