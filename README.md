[![Go Report Card](https://goreportcard.com/badge/github.com/aellwein/slf4go)](https://goreportcard.com/report/github.com/aellwein/slf4go)
[![Coverage Status](https://img.shields.io/coveralls/github/aellwein/slf4go/master.svg)](https://coveralls.io/github/aellwein/slf4go?branch=master)
[![Build Status](https://img.shields.io/travis/aellwein/slf4go/master.svg)](https://travis-ci.org/aellwein/slf4go) 

# SLF4GO
Simple Logging Facade for [Golang](https://www.golang.org), inspired by [SLF4J](https://www.slf4j.org).

# What is SLF4GO

SLF4J provides an abstraction over a particular logging framework used under the hood.

# How SLF4GO works

SLF4GO is actually a facade, the actual implementation of logging is hidden in and imported with 
SLF4GO adaptor code (see below for list of existing adaptors).

SLF4GO itself provides two interfaces, `LoggerFactory` and `Logger`.

`LoggerFactory` is adaptor-specific and is mostly the glue code to the underlying
logger implementation.

`Logger` provides the user-facing API, like `Trace`, `Debug`, `Info`, `Warn`, 
`Error`, `Fatal` and `Panic`.

# Usage

There are already some existing adaptors:
* [github.com/aellwein/slf4go-native-adaptor](https://github.com/aellwein/slf4go-native-adaptor) - logging implementation 
based upon standard Golang "log" package.
* [github.com/aellwein/slf4go-logrus-adaptor](https://github.com/aellwein/slf4go-logrus-adaptor) - an adaptor for 
[logrus](https://github.com/sirupsen/logrus) logging framework.
* [github.com/aellwein/slf4go-zap-adaptor](https://github.com/aellwein/slf4go-zap-adaptor) - an adaptor for 
[zap](https://github.com/uber-gp/zap) logging framework.



Now, in order to start using SLF4GO, you just need to do the following:
```sh
go get -u github.com/aellwein/slf4go

# now you need to get a particular adaptor, e.g. logrus:
go get -u github.com/aellwein/slf4go-logrus-adaptor

# or, for the native golang logging adaptor:
go get -u github.com/aellwein/slf4go-native-adaptor

```
The basic usage is always the same, you just need to import the correct
adaptor package you want to use with your application.

## Using native adaptor 

```go
package main

import (
    "github.com/aellwein/slf4go"
    _ "github.com/aellwein/slf4go-native-adaptor"
)

func main() {
    logger := slf4go.GetLogger("main")
    
    // our level is debug -> traces are hidden
    logger.SetLevel(slf4go.LevelDebug)
    
    logger.Debug("here goes some debug information")
    logger.Trace("this will not appear.")
    logger.Infof("Here is an extended form, logger '%s' with param %d.", logger.GetName(), 42)
    logger.Error("Some error occurred.")
    logger.Panic("This would print stack trace and cause panic.")
    logger.Fatal("This line would be logged and the program will terminate.")
}
```

## Use logrus adaptor

The same as above, only another adaptor is imported.
```go
package main

import (
    "github.com/aellwein/slf4go"
    _ "github.com/aellwein/slf4go-logrus-adaptor"
)

func main() {
    logger := slf4go.GetLogger("main")
    
    // our level is debug -> traces are hidden
    logger.SetLevel(slf4go.LevelDebug)
    
    logger.Debug("here goes some debug information")
    logger.Trace("this will not appear.")
    logger.Infof("Here is an extended form, logger '%s' with param %d.", logger.GetName(), 42)
    logger.Error("Some error occurred.")
    logger.Panic("This would print stack trace and cause panic.")
    logger.Fatal("This line would be logged and the program will terminate.")
}
```
You get the idea.


## Pass adaptor-specific parameters

To make the most of used logging framework, it's sometimes necessary to use its special features.
For this case, ``LoggingFactory.SetLoggingParameters()`` function can be called once before Logger
initialization, to pass a map of parameters to the logger adaptor.

Please note, that ``SetLoggingParameters()`` is **not threadsafe**, so you should called only once
**before** creating any ``Logger`` instance!

#### Example:

```go
package main

import (
	"github.com/aellwein/slf4go"
	_ "github.com/aellwein/slf4go-logrus-adaptor"
)

func main() {
	params := slf4go.LoggingParameters{
		"foo": 42,
		"bar": 1337,
		"baz": 0xDEADBEEF,
	}
	err := slf4go.GetLoggerFactory().SetLoggingParameters(params)
	if err != nil {
		panic(err)
	}
}

```


# Benefit

Separation of logging interface from its implementation is may be a good idea.
One day you will be able to change the underlying logging framework with another
one, just by changing the adaptor implementation (i.e. only changing the 
import statement).
