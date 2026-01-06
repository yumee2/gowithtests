package mocking

import (
	"fmt"
	"io"
	"time"
)

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(writer, "Go!")
}
