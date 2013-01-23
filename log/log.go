package log

import (
    "fmt"
//    "runtime"
)

type logger struct {
    level int
}

var l logger

const (
    all = iota
    debug
    msg
    info
    error
    none
)

func init() {

    // create a global logger object
    l.level = debug
}

func SetLevel(level string) {
    switch level {
    case "ALL": l.level = all
    case "DEBUG": l.level = debug
    case "ERROR": l.level = error
    case "MESSAGE": l.level = msg
    case "INFO": l.level = info
    case "NONE": l.level = none
    default: l.level = none
    }
}

func print(level int, format string, v ...interface{}) {

    if level >= l.level {
        tmp := fmt.Sprintf(format, v...)

        var msgType string

        switch level {
        case all: msgType = "ALL"
        case debug:
            msgType = "DEBUG"
//            pc, file, line, ok := runtime.Caller(3)
//            f := runtime.FuncForPC(pc)
//            
//            if ok {
//                fmt.Printf("%s %s (%s:%d)\n", msgType, f.Name(), file, line)
//            }
//            return

        case error: msgType = "ERROR"
        case msg: msgType = "MESSAGE"
        case info: msgType = "INFO"
        }

        fmt.Printf("%s: %v\n", msgType, tmp)

    }
}

func All(format string, v ...interface{}) {
    print(all, format, v...)
}

func Debug(format string, v ...interface{}) {
    print(debug, format, v...)
}

func Error(format string, v ...interface{}) {
    print(error, format, v...)
}

func Msg(format string, v ...interface{}) {
    print(msg, format, v...)
}

func Info(format string, v ...interface{}) {
    print(info, format, v...)
}

func None(format string, v ...interface{}) {
    print(none, format, v...)
}
