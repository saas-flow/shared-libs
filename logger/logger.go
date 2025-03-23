package logger

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once

var Module = fx.Module("logger",
	fx.Provide(New),
)

type Config struct {
	LogPath      string
	LogFilename  string
	LogMaxSize   int
	LogMaxBackup int
	LogMaxAge    int
	LogCompress  bool
}

func New(config Config) (log *zap.Logger) {

	rollingFile := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", config.LogPath, config.LogFilename),
		MaxSize:    config.LogMaxSize,
		MaxBackups: config.LogMaxBackup,
		MaxAge:     config.LogMaxAge,
		Compress:   config.LogCompress,
	}

	fields := zap.Fields(
		zap.String("service_name", os.Getenv("SERVICE_NAME")),
		zap.String("service_version", os.Getenv("SERVICE_VERSION")),
		zap.String("service_namespace", os.Getenv("SERVICE_NAMESPACE")),
	)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	bufferedWriter := zapcore.AddSync(&BufferedWriteSyncer{
		writer: bufio.NewWriterSize(os.Stdout, 1024), // Buffer 1KB
	})

	fileCore := &MaskingCore{
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(rollingFile),
			zap.InfoLevel,
		),
	}

	consoleCore := &MaskingCore{
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(bufferedWriter),
			zap.InfoLevel,
		),
	}

	core := zapcore.NewTee(consoleCore, fileCore)
	log = zap.New(core, fields)
	defer log.Sync()

	zap.ReplaceGlobals(log)

	return
}
