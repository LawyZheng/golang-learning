package log

import (
	"bufio"
	"io"
)

// Logger Object
type Logger struct {
	rd *bufio.Writer
}

// NewLogger new a obj
func NewLogger(rd io.Writer) *Logger {
	return &Logger{bufio.NewWriter(rd)}
}

// // NewLogger new a obj
// func NewLogger(rd interface{}) (l *Logger, err error) {
// 	err = error(nil)
// 	switch v := rd.(type) {
// 	case io.Writer:
// 		{
// 			l := &Logger{bufio.NewWriter(v)}
// 			return l, err
// 		}
// 	case string:
// 		{
// 			fileObj, err := os.OpenFile(v, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
// 			defer fileObj.Close()
// 			l := &Logger{bufio.NewWriter(fileObj)}
// 			return l, err

// 		}
// 		// default:
// 		// {
// 		// l := &Logger{bufio.NewWriter(v)}
// 		// }
// 		// }
// 	}
// 	return l, err
// }

// Info Log
func (l Logger) Info(msg string) {
	l.rd.WriteString("INFO: ")
	l.rd.WriteString(msg)
	l.rd.WriteString("\n")
	l.rd.Flush()
}
