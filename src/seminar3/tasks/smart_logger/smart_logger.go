package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Level int

const (
	Info Level = iota
	Warn
	Error
)

func (l Level) String() string {
	switch l {
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type SmartLogger struct {
	output   io.Writer
	prefix   string
	level    Level
	logCount int
	isColor  bool
}

func NewSmartLogger(output io.Writer, prefix string) *SmartLogger {
	return &SmartLogger{
		output:   output,
		prefix:   prefix,
		level:    Info,
		logCount: 0,
		isColor:  false,
	}
}

func (sl *SmartLogger) SetLevel(level Level) {
	sl.level = level
}

func (sl *SmartLogger) EnableColor() {
	sl.isColor = true
}

func (sl *SmartLogger) Write(p []byte) (n int, err error) {
	message := strings.TrimSpace(string(p))
	return sl.output.Write([]byte(sl.formatLog(Info, message)))
}

func (sl *SmartLogger) String() string {
	return fmt.Sprintf("SmartLogger{prefix: '%s', level: %s, logs: %d}",
		sl.prefix, sl.level, sl.logCount)
}

func (sl *SmartLogger) GoString() string {
	return fmt.Sprintf("SmartLogger{prefix: %q, level: %v, logCount: %d, isColor: %t}",
		sl.prefix, sl.level, sl.logCount, sl.isColor)
}

func (sl *SmartLogger) Info(format string, args ...interface{}) {
	if sl.level <= Info {
		sl.log(Info, format, args...)
	}
}

func (sl *SmartLogger) Warn(format string, args ...interface{}) {
	if sl.level <= Warn {
		sl.log(Warn, format, args...)
	}
}

func (sl *SmartLogger) Error(format string, args ...interface{}) {
	if sl.level <= Error {
		sl.log(Error, format, args...)
	}
}

// Вспомогательные методы
func (sl *SmartLogger) log(level Level, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	formatted := sl.formatLog(level, message)

	if sl.output != nil {
		sl.output.Write([]byte(formatted))
	}
	sl.logCount++
}

func (sl *SmartLogger) formatLog(level Level, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	var levelStr string
	if sl.isColor {
		levelStr = sl.colorizeLevel(level)
	} else {
		levelStr = fmt.Sprintf("[%s]", level)
	}

	return fmt.Sprintf("%s %s %s: %s\n", timestamp, sl.prefix, levelStr, message)
}

func (sl *SmartLogger) colorizeLevel(level Level) string {
	colorCode := "37"
	switch level {
	case Info:
		colorCode = "32"
	case Warn:
		colorCode = "33"
	case Error:
		colorCode = "31"
	}

	return fmt.Sprintf("\033[%sm[%s]\033[0m", colorCode, level)
}

func (sl *SmartLogger) Close() error {
	if closer, ok := sl.output.(io.Closer); ok && sl.output != os.Stdout {
		return closer.Close()
	}
	return nil
}

func (sl *SmartLogger) Reset() {
	sl.logCount = 0
}

func (sl *SmartLogger) GetLogCount() int {
	return sl.logCount
}

func main() {
	fmt.Println("=== Демонстрация SmartLogger ===")

	// 1. Создаем логгер для консоли
	consoleLogger := NewSmartLogger(os.Stdout, "APP")
	consoleLogger.EnableColor()

	// Используем как обычный логгер
	consoleLogger.Info("Приложение запущено")
	consoleLogger.Warn("Нагрузка выше обычной: %.1f%%", 85.5)
	consoleLogger.Error("Ошибка подключения к БД")

	// 2. Используем как io.Writer
	fmt.Println("\n=== Использование как io.Writer ===")
	fmt.Fprintf(consoleLogger, "Это сообщение через fmt.Fprintf")

	// 3. Демонстрация интерфейсов Stringer и GoStringer
	fmt.Println("\n=== Stringer и GoStringer ===")
	fmt.Println("String():", consoleLogger.String())
	fmt.Printf("GoString(): %#v\n", consoleLogger)

	// 4. Логгер в буфер (удовлетворяет io.Writer)
	fmt.Println("\n=== Логгер в буфер ===")
	var buf strings.Builder
	bufferLogger := NewSmartLogger(&buf, "TEST")
	bufferLogger.Info("Тестовое сообщение 1")
	bufferLogger.Warn("Тестовое сообщение 2")

	fmt.Println("Логи в буфере:")
	fmt.Print(buf.String())
	fmt.Printf("Всего логов: %d\n", bufferLogger.GetLogCount())

	// 5. Фильтрация по уровню
	fmt.Println("\n=== Фильтрация по уровню ===")
	filteredLogger := NewSmartLogger(os.Stdout, "FILTERED")
	filteredLogger.SetLevel(Warn) // Только Warn и Error

	filteredLogger.Info("Это сообщение НЕ должно появиться") // Не появится
	filteredLogger.Warn("А это должно появиться")            // Появится
	filteredLogger.Error("И это тоже")                       // Появится

	// 6. Использование в функциях, принимающие io.Writer
	fmt.Println("\n=== Использование с стандартными функциями ===")
	writeToLogger(consoleLogger, "Сообщение через функцию")
}

// Функция, принимающая io.Writer - наш логгер подходит!
func writeToLogger(w io.Writer, message string) {
	fmt.Fprintf(w, "Пишем в io.Writer: %s", message)
}
