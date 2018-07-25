# Go Log [![Go Doc](https://godoc.org/github.com/Noah-Huppert/golog?status.svg)](http://godoc.org/github.com/Noah-Huppert/golog)
A simple Go logging package.

# Table Of Contents
- [Overview](#overview)
- [Usage](#usage)
- [Advanced Usage](#advanced-usage)

# Overview
Go Log is a simple logging library inspired by the standard `log` package.

It was created because there was a lack of simple easy to use logging libraries 
in the Go ecosystem.  

Features:

- 1 line setup
- Log levels: Debug, Info, Warn, Error, Fatal
- Log output formatting via Go templates
- Function names similar to `log` package
- Advanced customizability if needed

# Usage
Import Go Log into your project:

```go
import "github.com/Noah-Huppert/golog"
```

Create a logger:

```go
// logger will print normal messages to stdout and errors to stderr
logger := golog.NewStdLogger("example")
```

Log:

```go
logger.Debug("hello debug")
logger.Debugf("hello %s", "debug")

logger.Info("hello info")
logger.Infof("hello %s", "info")

logger.Warn("hello warn")
logger.Warnf("hello %s", "warn")

logger.Error("hello error")
logger.Errorf("hello %s", "error")

logger.Fatal("hello fatal")
logger.Fatalf("hello %s", "fatal")
```

# Advanced Usage
Set logging level:

```go
logger.SetLevel(golog.DebugLevel)

logger.SetLevel(golog.InfoLevel)

logger.SetLevel(golog.WarnLevel)

logger.SetLevel(golog.ErrorLevel)

logger.SetLevel(golog.FatalLevel)
```

Customize output format:

```go
logger.SetFormatTmpl("name={{ .Name }} level={{ .Level }} msg={{ .Msg }}")
```

Customize output location:

```go
import (
	"io"
	"os"

	"github.com/Noah-Huppert/golog"
)

func openFile(name string) *io.File {
	f, err := io.OpenFile(name, os.O_RDONLY|os.O_CREATE, 666)
	if err != nil {
		panic(err)
	}

	return f
}

fatalLogF := openFile("fatal.log")
defer fatalLogF.Close()

errorLogF := openFile("error.log")
defer errorLogF.Close()

warnLogF := openFile("warn.log")
defer warnLogF.Close()

infoLogF := openFile("info.log")
defer infoLogF.Close()

debugLogF := openFile("debug.log")
defer debugLogF.Close()

// logger will write each log level's output to a different file
logger := golog.NewWriterLogger("example", fatalLogF, errorLogF, warnLogF,
	infoLogF, debugLogF)
```
