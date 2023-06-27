package logger

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	aggregationSpotter "github.com/Entangle-Protocol/off-chain-apps/shared/contracts/aggregationSpotter"
)

type logger struct {
	logger *log.Logger
}

func Fatalf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	os.Exit(1)
}

func Fatal(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func Println(v ...interface{}) (n int, err error) {
	return fmt.Println(v...)
}

func Printf(format string, v ...interface{}) (n int, err error) {
	return fmt.Printf(format, v...)
}

func Errorf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}

func LogProposeOperation(data aggregationSpotter.AggregationSpotterOperationData) error {
	filePath := "logs/transactions.log"
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute file path: %v", err)
	}

	dirPath := filepath.Dir(absFilePath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	file, err := os.OpenFile(absFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// lggr := log.New(file, "", log.LstdFlags)
	l := &logger{
		logger: log.New(file, "", 0),
	}

	currentTime := time.Now()
	formattedTime := currentTime.Format("2006/01/02 15:04:05")
	fmt.Println(formattedTime)
	logString := fmt.Sprintf("{Timestamp: \"%s\", Contr: \"%s\", FunctionSelector: \"%x\", Params: \"%s\"},", formattedTime, data.Contr.Hex(), data.FunctionSelector, hex.EncodeToString(data.Params))

	l.Println(logString)

	return nil
}

func (l *logger) Println(v ...interface{}) {
	v = append([]interface{}{""}, v...)
	l.logger.Println(v...)
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(time.Now().Format("2006-01-02 15:04:05")+" "+format, v...)
}

// var (
// 	txLogger *log.Logger
// 	// allLoger *log.Logger
// )

// func InitCommonLoggers() {
// 	err := os.MkdirAll("logs", os.ModePerm)
// 	if err != nil {
// 		logger.Fatalf("Error creating directory: %v", err)
// 	}

// 	txLogFile, err := os.OpenFile("logs/transactions.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	// defer txLogFile.Close()

// 	txLogger = log.New(txLogFile, "[tx]\t", log.LstdFlags)
// }

// func createCustomLogger(loggerFile string) *log.Logger {
// 	if loggerFile == "" {
// 		return nil
// 	}

// 	err := os.MkdirAll("logs", os.ModePerm)
// 	if err != nil {
// 		logger.Fatalf("Error creating directory: %v", err)
// 	}

// 	logFile, err := os.OpenFile("logs/"+loggerFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	// defer logFile.Close()

// 	logger := log.New(logFile, "", log.LstdFlags)

// 	return logger
// }
