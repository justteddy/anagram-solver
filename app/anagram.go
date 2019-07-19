package app

import (
	"log"
	"sort"
	"strings"
	"sync"

	"golang.org/x/net/context"
)

// AnagramFinder интерфейс для взаимодействия со словарем для поиска анаграмм.
type AnagramFinder interface {
	// LoadWords загружает слова.
	LoadWords(ctx context.Context, words []string) error
	// SearchAnagrams ищет анаграммы для переданного слова в словаре.
	SearchAnagrams(ctx context.Context, word string) []string
}

// NewAnagramFinder инициализирует AnagramFinder.
func NewAnagramFinder() AnagramFinder {
	return &anagramFinder{
		dictionary: map[string]map[string]struct{}{},
	}
}

// anagramFinder реализация интерфейса AnagramFinder.
type anagramFinder struct {
	sync.RWMutex
	dictionary map[string]map[string]struct{}
}

// dictionaryData элемент данных для словаря.
type dictionaryData struct {
	// нормализованное представление слова.
	key string
	// оригинал слова.
	word string
}

// LoadWords реализация интерфейса AnagramFinder.
func (a *anagramFinder) LoadWords(ctx context.Context, words []string) error {
	dictElementsCh := prepareWords(ctx, words)

	a.Lock()
	defer a.Unlock()

	for dictData := range dictElementsCh {
		if _, ok := a.dictionary[dictData.key]; !ok {
			a.dictionary[dictData.key] = map[string]struct{}{
				dictData.word: {},
			}
			continue
		}
		a.dictionary[dictData.key][dictData.word] = struct{}{}
	}

	return nil
}

// SearchAnagrams реализация интерфейса AnagramFinder.
func (a *anagramFinder) SearchAnagrams(ctx context.Context, word string) []string {
	a.RLock()
	defer a.RUnlock()

	if wordMap, ok := a.dictionary[normalizeWord(word)]; ok {
		var i int
		anagrams := make([]string, len(wordMap))
		for anagram := range wordMap {
			anagrams[i] = anagram
			i++
		}
		sort.Slice(anagrams, func(i, j int) bool {
			return anagrams[i] < anagrams[j]
		})
		return anagrams
	}

	return nil
}

// prepareWords подготовка массива поступающих в словарь элементов.
func prepareWords(ctx context.Context, words []string) <-chan dictionaryData {
	dictElementsCh := make(chan dictionaryData, len(words))
	go func() {
		defer close(dictElementsCh)

		var wg sync.WaitGroup
		wg.Add(len(words))
		for _, word := range words {
			go prepareWord(ctx, dictElementsCh, &wg, word)
		}
		wg.Wait()
	}()

	return dictElementsCh
}

// prepareWord подготовка отдельного слова,поступающего в словарь элементов.
func prepareWord(ctx context.Context, dictElementsCh chan<- dictionaryData, wg *sync.WaitGroup, word string) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		log.Print("Context failed unexpectedly")
		return
	case dictElementsCh <- dictionaryData{key: normalizeWord(word), word: word}:
	}
}

// normalizeWord нормализует слово для словаря.
func normalizeWord(word string) string {
	runes := make([]string, 0)
	for _, r := range word {
		runes = append(runes, strings.ToLower(string(r)))
	}
	sort.Strings(runes)
	return strings.Join(runes, "")
}
