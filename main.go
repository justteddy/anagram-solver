package main

import (
	"log"

	"github.com/go-openapi/loads"

	"anagram-solver/generated/restapi"
	"anagram-solver/generated/restapi/operations"

	"anagram-solver/app"
)

const defaultPort = 17001

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewAnagramAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	// обработчики эндпоинтов
	service := app.NewService()
	api.LoadDictionaryHandler = operations.LoadDictionaryHandlerFunc(service.HandleLoadDictionary)
	api.SearchAnagramsHandler = operations.SearchAnagramsHandlerFunc(service.HandleSearchAnagrams)

	server.Port = defaultPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
