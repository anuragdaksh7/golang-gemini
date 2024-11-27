package main

import (
	"gemini-go/internal/gemini"
	"gemini-go/router"
	"gemini-go/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	geminiSvc := gemini.NewService()
	geminiHandler := gemini.NewHandler(geminiSvc)

	router.InitRouter(geminiHandler)
	log.Fatal(router.Start("0.0.0.0:" + config.PORT))
}
