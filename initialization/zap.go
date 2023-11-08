package initialization

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var SeatLogger *zap.Logger

func InitZap() {
	// 创建自定义的时间格式
	timeFormat := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(Config.Zap.Prefix + "2006-01-02,15:04:05"))
	}
	//默认的生产配置
	log := zap.NewProductionConfig()
	//输出至控制台
	log.OutputPaths = []string{"stdout"}
	//输出至文件，先检查是否存在该目录和文件，不存在就创建
	filePath := Config.Zap.FilePath
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// 文件不存在，创建它
		err := os.MkdirAll("./log", os.ModePerm) // 创建目录 "./log"，如果已存在则忽略
		if err != nil {
			panic(fmt.Errorf("Fatal error zap file: %s \n", err))
		}
		// 创建文件
		file, err := os.Create(filePath)
		if err != nil {
			panic(fmt.Errorf("Fatal error zap file: %s \n", err))
		}
		defer file.Close()
	} else if err != nil {
		panic(fmt.Errorf("Fatal error zap file: %s \n", err))
	}
	log.OutputPaths = append(log.OutputPaths, filePath)

	log.EncoderConfig.EncodeTime = timeFormat
	log.Encoding = Config.Zap.Format
	SeatLogger, err = log.Build()
	if err != nil {
		panic("Failed to create logger: " + err.Error())
	}
	defer SeatLogger.Sync()
}
