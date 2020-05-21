package errors

import (
	"log"
	"runtime"
)

var _, fn, line, _ = runtime.Caller(1)

func CheckError(e error) {
	if e != nil {
		log.Printf(`[ERRO]: %s:%d => %s`, fn, line, e)
	}
}

func CheckErrorMsg(e error, msg string) {
	if e != nil {
		log.Printf(`[ERRO]: %s:%d %s => %s`, fn, line, msg, e)
		// log.Println(msg, " => ", e)
	}
}
