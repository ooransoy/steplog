package steplog

import (
	"fmt"
)

var StripedPrint bool = false

type Steplogger struct {
	Index int
	Depth int
}

func NewSteplogger() *Steplogger {
	return &Steplogger{}
}

func (l *Steplogger) Open(str string) {
	indent := ""
	for i := 0; i < l.Depth; i++ {
		indent += "  "
	}
	if l.Index&1 == 1 && StripedPrint {
		fmt.Println("\x1b[48;5;8m" + indent + str + "\x1b[49m")
	} else {
		fmt.Println(indent + str)
	}
	l.Index++
	l.Depth++
}

func (l *Steplogger) Close() {
	l.Depth--
}

func (l *Steplogger) Print(str string) {
	l.Open(str)
	l.Close()
}
