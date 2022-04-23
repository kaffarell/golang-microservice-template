package lib

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

// Logger structure
type Logger struct {
	*logrus.Logger
}

type GinLogger struct {
	*Logger
}

type FxLogger struct {
	*Logger
}

var (
	globalLogger *Logger
	logrusLogger *logrus.Logger
)

// GetLogger get the logger
func GetLogger() Logger {
	if globalLogger == nil {
		logger := newLogger()
		globalLogger = &logger
	}
	return *globalLogger
}

// GetGinLogger get the gin logger
func (l Logger) GetGinLogger() GinLogger {
	return GinLogger{
		Logger: &Logger{Logger: logrusLogger},
	}
}

// GetFxLogger gets logger for go-fx
func (l *Logger) GetFxLogger() fxevent.Logger {
	return &FxLogger{Logger: &Logger{Logger: logrusLogger}}
}

// LogEvent log event for fx logger
func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.WithFields(logrus.Fields{
			"callee": e.FunctionName,
			"caller": e.CallerName,
		}).Debug("OnStart hook executing")
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.WithFields(logrus.Fields{
				"callee": e.FunctionName,
				"caller": e.CallerName,
				"error":  e.Err,
			}).Debug("OnStart hook failed")
		} else {
			l.Logger.WithFields(logrus.Fields{
				"callee":  e.FunctionName,
				"caller":  e.CallerName,
				"runtime": e.Runtime.String(),
			}).Debug("OnStart hook executed")
		}
	case *fxevent.OnStopExecuting:
		l.Logger.WithFields(logrus.Fields{
			"callee": e.FunctionName,
			"caller": e.CallerName,
		}).Debug("OnStop hook executing")
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.WithFields(logrus.Fields{
				"callee": e.FunctionName,
				"caller": e.CallerName,
				"error":  e.Err,
			}).Debug("OnStop hook failed")
		} else {
			l.Logger.WithFields(logrus.Fields{
				"callee":  e.FunctionName,
				"caller":  e.CallerName,
				"runtime": e.Runtime.String(),
			}).Debug("OnStop hook executed")
		}
	case *fxevent.Supplied:
		l.Logger.WithFields(logrus.Fields{
			"type":  e.TypeName,
			"error": e.Err,
		}).Debug("supplied: ")
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("provided: " + e.ConstructorName + " => " + rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.WithFields(logrus.Fields{
				"decorator": e.DecoratorName,
				"type":      rtype,
			}).Debug("decorated: ")
		}
	case *fxevent.Invoking:
		l.Logger.Debug("invoking: " + e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			l.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			//l.Logger.Debug("initialized: custom fxevent.Logger -> " + e.ConstructorName)
			l.Logger.Debug("testcool")
		}
	}
}

// newLogger sets up logger
func newLogger() Logger {
	logrusLogger := logrus.New()

	logger := &Logger{Logger: logrusLogger}

	return *logger
}

// Write interface implementation for gin-framework
func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Logger.Info(string(p))
	return len(p), nil
}

// Printf prints go-fx logs
func (l FxLogger) Printf(str string, args ...interface{}) {
	if len(args) > 0 {
		l.Logger.Debug(str, args)
	}
	l.Debug(str)
}
