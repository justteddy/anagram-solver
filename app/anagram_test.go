package app

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"golang.org/x/net/context"
)

func TestAnagramFinderSearch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cases := []struct {
		dictionary []string
		word       string

		expected []string
	}{
		{
			dictionary: nil,
			word:       "abba",
			expected:   nil,
		},
		{
			dictionary: []string{"foobar"},
			word:       "",
			expected:   nil,
		},
		{
			dictionary: []string{"foobar"},
			word:       "foo",
			expected:   nil,
		},
		{
			dictionary: []string{"FoObaR", "barfoo"},
			word:       "OoFaRb",
			expected:   []string{"FoObaR", "barfoo"},
		},
		{
			dictionary: []string{"foobar", "aabb", "baba", "boofar", "test"},
			word:       "abba",
			expected:   []string{"aabb", "baba"},
		},
		{
			dictionary: []string{"lena", "elan", "lane", "neal", "lean"},
			word:       "ealn",
			expected:   []string{"elan", "lane", "lean", "lena", "neal"},
		},
		{
			dictionary: []string{"вижу", "живу"},
			word:       "ужви",
			expected:   []string{"вижу", "живу"},
		},
	}

	for _, c := range cases {
		af := NewAnagramFinder()
		err := af.LoadWords(ctx, c.dictionary)
		assert.Equal(t, err, nil, "loading dictionary should not return error")

		anagrams := af.SearchAnagrams(ctx, c.word)
		assert.Equal(t, anagrams, c.expected)
	}
}
