package app

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/stack"
)

// Chk checks if the error var is nil
// If not it prints the error and stops execution
func Chk(err error) {
	if nil != err {
		fmt.Println(stack.CallersMulti(0))
		log.Fatal(err)
	}
}

// PrintWelcome prints the welcome message
func PrintWelcome() {
	buf := new(bytes.Buffer)
	buf.Write([]byte{'\033', '[', '3', '2', 'm'})
	fmt.Fprint(buf, "\n\n\n")
	fmt.Fprint(buf, "     _     _                 _                   	 ")
	fmt.Fprint(buf, "    | |__ | | ___   ___   __| | ___ __ _ _ __ ___   ")
	fmt.Fprint(buf, "    | '_ \\| |/ _ \\ / _ \\ / _` |/ __/ _` | '__/ _ \\  ")
	fmt.Fprint(buf, "    | |_) | | (_) | (_) | (_| | (_| (_| | | |  __/  ")
	fmt.Fprint(buf, "    |_.__/|_|\\___/ \\___/ \\__,_|\\___\\__,_|_|  \\___|  ")
	fmt.Fprint(buf, "     \n")
	fmt.Fprint(buf, "     \n")
	buf.Write([]byte{'\033', '[', '0', 'm'})
	fmt.Fprint(buf, "\n\n")
	fmt.Println(buf.String())
}

// ParseErrors parses the error array and formats the error for TTY
func ParseErrors(errors []string) {
	if len(errors) == 0 {
		return
	}
	buf := new(bytes.Buffer)
	fmt.Fprint(buf, "\nInvalid environment variables/flags          \n")
	buf.Write([]byte{'\033', '[', '3', '1', 'm'})
	for _, err := range errors {
		fmt.Fprint(buf, "  ", err, "\n")
	}
	buf.Write([]byte{'\033', '[', '0', 'm'})
	fmt.Println(buf.String())
	os.Exit(0)
}

// Debug prints the data and stops program flow
func Debug(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Printf("val: %v", data)
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		fmt.Printf("val: %v", data)
	case string:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		fmt.Printf("val: %v", data)
	default:
		fmt.Println(data)
	}
	os.Exit(10)
}

//GetLogrMode takes the mode from commandline and returns logrMode
func GetLogrMode(mode string) logrus.Level {
	switch mode {
	case MODE_DEV:
		return logrus.Info
	case MODE_PROD:
		return logrus.Warn
	case MODE_DEBUG:
		return logrus.Debug
	}
	return logrus.Warn
}

func TruncWords(s string, n int) string {
	if n == 0 {
		return s
	}
	count := 0
	sep := " "
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] {
			count++
		}
		if count == n {
			return s[0:i] + "..."
		}
	}
	if count < n {
		return s
	}
	return ""
}
