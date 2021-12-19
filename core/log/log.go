package log

import (
	"bitrade/core/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	zapLog   *zap.SugaredLogger // 简易版日志文件
	logLevel = zap.NewAtomicLevel()
	prefix   = ""
)

type Level int8

//日志级别
const (
	DebugLevel Level = iota - 1

	InfoLevel

	WarnLevel

	ErrorLevel

	DPanicLevel

	PanicLevel

	FatalLevel
)

//日志输出目标类型
const (
	OUT_PRINT_File = iota
	OOUT_PRINT_CONSOLE
	OUT_PRINT_ALL
)

// LogConf 本地日志的配置
type Config struct {
	Path  string `ini:"path"`
	Level string `ini:"level"`
	Out   int8   `ini:"out"`
}

// GetLogConf 获取日志的配置
func (conf *Config) GetLogConf() *Config {
	return conf
}

// InitLog 初始化日志文件
func InitLog() (err error) {
	var conf Config
	err = config.GetConfig().Section("Log").MapTo(&conf)
	if err != nil {
		return
	}
	cores := make([]zapcore.Core, 0)

	loglevel := InfoLevel
	switch conf.Level {
	case "INFO":
		loglevel = InfoLevel
	case "ERROR":
		loglevel = ErrorLevel
	}
	setLevel(loglevel)

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	//输入日志到文件
	if conf.Out == OUT_PRINT_File || conf.Out == OUT_PRINT_ALL {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.Path,
			MaxSize:    128, // MB
			LocalTime:  true,
			Compress:   true,
			MaxBackups: 8, // 最多保留 n 个备份
		})

		file := zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			logLevel,
		)
		cores = append(cores, file)
	}

	//输入日志到控制台
	if conf.Out == OOUT_PRINT_CONSOLE || conf.Out == OUT_PRINT_ALL {
		console := zapcore.NewCore(
			zapcore.NewConsoleEncoder(config),
			zapcore.AddSync(os.Stdout),
			logLevel,
		)
		cores = append(cores, console)
	}

	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
	zapLog = logger.Sugar()

	return nil
}

func setLevel(level Level) {
	logLevel.SetLevel(zapcore.Level(level))
}

func Info(args ...interface{}) {
	args = append([]interface{}{prefix}, args...)
	zapLog.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zapLog.Infof(prefix+template, args...)
}

func Warn(args ...interface{}) {
	args = append([]interface{}{prefix}, args...)
	zapLog.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zapLog.Warnf(prefix+template, args...)
}

func Error(args ...interface{}) {
	args = append([]interface{}{prefix}, args...)
	zapLog.Error(args...)
}

func Errorf(template string, args ...interface{}) {

	zapLog.Errorf(prefix+template, args...)
}
