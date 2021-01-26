package rac

type Logger interface {
	Errorf(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
}

type nullLogger struct{}

func (n nullLogger) Debugf(msg string, args ...interface{}) {}

func (n nullLogger) Errorf(msg string, args ...interface{}) {}

var _ Logger = (*nullLogger)(nil)
