package app

import (
	"fmt"
	"os"
)

func ErrorHappened(msg ...interface{}) {
	fmt.Print("Error")
	fmt.Print(msg...)
	os.Exit(1)
}

func print(msg ...interface{}) {
	fmt.Print(msg...)

}
