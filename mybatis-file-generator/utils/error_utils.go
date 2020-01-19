package utils

import (
	"log"
)

type PanicError struct {
	error error
}

func TryRecover() {

	//最后运行recover
	defer func() {
		r := recover()
		if err, ok := r.(PanicError); ok {
			//自动退出程序
			log.Panicln("Panic Error", err)
		} else if err, ok := r.(error); ok {
			log.Println("Error occurred", err)
		} else if err, ok := r.(string); ok {
			log.Println("Panic occurred", err)
		} else {
			log.Fatalln("other error occurred", r)
		}
	}()
}

func LogFatal(msg string, err error) {
	if err != nil {
		log.Fatalln("fatal error (", msg, ")", err)
	}
}

func LogPanic(msg string, err error) {
	if err != nil {
		log.Panicln("panic error (", msg, ")", err)
	}
}

func LogPrint(msg string, err error) {
	if err != nil {
		log.Println("error occurred (", msg, ")", err)
	}

}
