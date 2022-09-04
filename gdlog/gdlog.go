package gdlog

import "log"

func Debug(v ...any) {
	log.Println("DEBUG", v)
}

func Info(v ...any) {
	log.Println("INFO", v)
}

func Error(v ...any) {
	log.Println("ERROR", v)
}
