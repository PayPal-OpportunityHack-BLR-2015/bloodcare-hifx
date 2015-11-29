package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Prase() parses the environment variables based on the following rules and
// sets the flags
//		- "prefix" is prefixed before the variable name
//		- dots are converted to underscores
//		- dashes are converted to underscores
//		- resulting variable name is converted to uppercase
func EnvParse(prefix string) {
	if err := parse(prefix); err != nil {
		log.Fatalln(err)
	}
}

// ParseAll() does the same functionality as Prase()
// except that it omits the "prefix" string
func ParseAll() {
	if err := parse(""); err != nil {
		log.Fatalln(err)
	}
}

// parse() iterates through the flags and populate the
// values from environment variables if present
func parse(prefix string) (errp error) {
	cmdline := make(map[string]bool)

	flag.Visit(func(f *flag.Flag) {
		cmdline[f.Name] = true
	})

	flag.VisitAll(func(f *flag.Flag) {
		if _, ok := cmdline[f.Name]; !ok {
			val := os.Getenv(getEnvName(prefix, f.Name))
			if val != "" {
				err := f.Value.Set(val)
				if err != nil {
					if errp != nil {
						errp = fmt.Errorf(errp.Error()+", value:\"%s\" for flag:\"%s\"", val, f.Name)
					} else {
						errp = fmt.Errorf(":envtoflag: error setting  value:\"%s\" for flag:\"%s\"", val, f.Name)
					}

				}
			}
		}
	})

	return
}

// getEnvName() takes the
func getEnvName(prefix string, flag string) string {
	name := strings.ToUpper(prefix + "_" + flag)
	name = strings.Replace(name, ".", "_", -1)
	name = strings.Replace(name, "-", "_", -1)
	return name
}
