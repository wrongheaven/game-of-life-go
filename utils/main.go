package utils

import "log"

func Bekind(err error) {
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}
