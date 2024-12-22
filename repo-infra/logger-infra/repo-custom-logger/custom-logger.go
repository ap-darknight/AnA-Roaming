package repo_custom_logger

import (
	"fmt"
	"go.uber.org/zap"
)

type CustomLogger struct {
	Logger       *zap.SugaredLogger
	ReqID        string
	RootFunction string
}

// NewCustomLogger creates an instance of CustomLogger with a zap logger
func NewRepoCustomLogger(logger *zap.SugaredLogger) *CustomLogger {
	return &CustomLogger{Logger: logger}
}

// Prefix with pre-added data
func (c *CustomLogger) addPrefix(message string) string {
	prefix := fmt.Sprintf("[%s, %s] ", c.ReqID, c.RootFunction)
	return prefix + message
}

// ErrorwWrapper wraps the Errorw method with prefixed data in the log message
func (c *CustomLogger) Errorw(msg string, additionalFields ...interface{}) {
	msgWithPrefix := c.addPrefix(msg)
	c.Logger.Errorw(msgWithPrefix, append(additionalFields)...)
}

// InfofWrapper wraps the Infof method with prefixed data in the log message
func (c *CustomLogger) Infof(format string, args ...interface{}) {
	msgWithPrefix := c.addPrefix(fmt.Sprintf(format, args...))
	c.Logger.Info(msgWithPrefix)
}

// ErrorfWrapper wraps the Errorf method with prefixed data in the log message
func (c *CustomLogger) Errorf(format string, args ...interface{}) {
	msgWithPrefix := c.addPrefix(fmt.Sprintf(format, args...))
	c.Logger.Error(msgWithPrefix)
}

// InfowWrapper wraps the Infow method with prefixed data in the log message
func (c *CustomLogger) Infow(msg string, additionalFields ...interface{}) {
	msgWithPrefix := c.addPrefix(msg)
	c.Logger.Infow(msgWithPrefix, additionalFields...)
}
