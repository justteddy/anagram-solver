package app

import (
	"log"

	"github.com/go-openapi/runtime/middleware"

	"anagram-solver/generated/restapi/operations"
)

// NewService иницивализация объекта, обработчика запросов.
func NewService() *Service {
	return &Service{
		anagram: NewAnagramFinder(),
	}
}

// Service объект обработчик запросов сервиса.
type Service struct {
	anagram AnagramFinder
}

// HandleLoadDictionary POST /load
func (s *Service) HandleLoadDictionary(params operations.LoadDictionaryParams) middleware.Responder {
	log.Printf("Starting to load dictionary: %v", params.Body)
	defer log.Print("Finished to load")

	ctx := params.HTTPRequest.Context()
	if err := s.anagram.LoadWords(ctx, params.Body); err != nil {
		log.Printf("Failed to load dictionary with error: %s", err)
		return operations.NewLoadDictionaryInternalServerError()
	}

	return operations.NewLoadDictionaryOK()
}

// HandleSearchAnagrams GET /get?word={word}
func (s *Service) HandleSearchAnagrams(params operations.SearchAnagramsParams) middleware.Responder {
	log.Printf("Starting to search anagrams for '%s'", params.Word)
	defer log.Print("Finished to search")

	ctx := params.HTTPRequest.Context()
	anagrams := s.anagram.SearchAnagrams(ctx, params.Word)

	return operations.NewSearchAnagramsOK().WithPayload(anagrams)
}
