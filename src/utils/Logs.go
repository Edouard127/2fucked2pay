package utils

import (
	"fmt"
	"os"
	"time"
)

var (
	date = time.Now().Format("2006-01-02-15-04-05")
)

func LogFile(s string) {
	f, err := os.OpenFile("logs/"+date+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		f, err = os.Create("logs/" + date + ".log")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	str := fmt.Sprintf("[%s]: %s", time.Now().Format("2006-01-02-15-04-05"), s)
	_, err = f.WriteString(str + "\n")
	if err != nil {
		return
	}

}
