package service

import (
	"log"
)

func CatchErr(err error) {

	if err != nil {
		log.Fatal(err)
	}

}
