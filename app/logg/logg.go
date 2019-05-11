package logg

import "log"

func Info(v interface{}) {

	log.Println("INFO:", v)

}

func Warn(v interface{}) {

	log.Println("WARN:", v)

}

func Error(v interface{}) {

	log.Println("ERROR:", v)

}
