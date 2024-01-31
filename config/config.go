package config

var (
	logger *Logger
)

func GetLogger(p string) *Logger {
	// Initialize logger
	logger := NewLogger(p)
	return logger

}
