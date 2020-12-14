package util

import "log"

// Panic prints a non-nil error and panics.
func Panic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
