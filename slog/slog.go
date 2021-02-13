package slog

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//Config logging configuration
type Config struct {
	Pretty bool
	Level  zerolog.Level
}

//Log Log object
type Log struct {
	logger zerolog.Logger
}

//New Creates a new log
func New(config *Config, source string) *Log {
	if config == nil {
		config = &Config{
			Pretty: true,
			Level:  zerolog.InfoLevel,
		}
	}
	logger := log.Level(config.Level)
	if source != "" {
		logger = logger.With().Str("source", source).Logger()
	}
	if config.Pretty {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	return &Log{logger}
}

//Trace log Trace message
func (m *Log) Trace(msg string) {
	m.logger.Trace().Msg(msg)
}

//Tracef log Trace message formatted
func (m *Log) Tracef(msg string, v ...interface{}) {
	m.logger.Trace().Msgf(msg, v...)
}

//Debug log Debug message
func (m *Log) Debug(msg string) {
	m.logger.Debug().Msg(msg)
}

//Debugf log Debug message formatted
func (m *Log) Debugf(msg string, v ...interface{}) {
	m.logger.Debug().Msgf(msg, v...)
}

//Info log info message
func (m *Log) Info(msg string) {
	m.logger.Info().Msg(msg)
}

//Infof log info message formatted
func (m *Log) Infof(msg string, v ...interface{}) {
	m.logger.Info().Msgf(msg, v...)
}

//Warn log Warn message
func (m *Log) Warn(msg string) {
	m.logger.Warn().Msg(msg)
}

//Warnf log Warn message formatted
func (m *Log) Warnf(msg string, v ...interface{}) {
	m.logger.Warn().Msgf(msg, v...)
}

//Error log error message
func (m *Log) Error(err error, msg string) {
	m.error(err).Msg(msg)
}

//Errorf log Error message formatted
func (m *Log) Errorf(err error, msg string, v ...interface{}) {
	m.error(err).Msgf(msg, v...)
}

func (m *Log) error(err error) *zerolog.Event {
	e := m.logger.Error()
	if err != nil {
		e = e.Stack().Err(err)
	}
	return e
}

//Fatal log Fatal message
func (m *Log) Fatal(err error, msg string) {
	m.fatal(err).Msg(msg)
}

//Fatalf log Fatal message formatted
func (m *Log) Fatalf(err error, msg string, v ...interface{}) {
	m.fatal(err).Msgf(msg, v...)
}

func (m *Log) fatal(err error) *zerolog.Event {
	e := m.logger.Fatal()
	if err != nil {
		e = e.Stack().Err(err)
	}
	return e
}

//Panic log Panic message
func (m *Log) Panic(err error, msg string) {
	m.panic(err).Msg(msg)
}

//Panicf log Panic message formatted
func (m *Log) Panicf(err error, msg string, v ...interface{}) {
	m.panic(err).Msgf(msg, v...)
}

func (m *Log) panic(err error) *zerolog.Event {
	e := m.logger.Panic()
	if err != nil {
		e = e.Stack().Err(err)
	}
	return e
}
