// Go Log is a simple to use logging library inspired by the standard `log`
// package.
//
// It provides the: Debug, Info, Warn, Error, and Fatal log levels.
//
// Go Log is easy to setup:
//
// ```
// // logger will print normal messages to stdout and errors to stderr
// logger := golog.NewStdLogger("example")
// ```
//
// The logging API should be familiar to those who have used the standard
// `fmt` and `log` packages.
//
// ```
// logger.Debug("hello debug")
// logger.Debugf("hello %s", "debug")
//
// logger.Info("hello info")
// logger.Infof("hello %s", "info")
//
// logger.Warn("hello warn")
// logger.Warnf("hello %s", "warn")
//
// logger.Error("hello error")
// logger.Errorf("hello %s", "error")
//
// logger.Fatal("hello fatal")
// logger.Fatalf("hello %s", "fatal")
// ```
package golog
