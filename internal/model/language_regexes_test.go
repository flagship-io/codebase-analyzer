package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
)

func TestCustomRegex(t *testing.T) {
	AddCustomRegexes(`[{"file_extension":".tsx?","regexes":["\\s*['\"](.+?)['\"]"]}]`)

	found := funk.Find(LanguageRegexes, func(languageRegex LanguageRegex) bool {
		return languageRegex.FileExtension == ".tsx?"
	})

	assert.NotNil(t, found)
	foundLanguageRegex, ok := found.(LanguageRegex)

	assert.True(t, ok)
	assert.NotNil(t, foundLanguageRegex)

	assert.Equal(t, ".tsx?", foundLanguageRegex.FileExtension)
	assert.Equal(t, 1, len(foundLanguageRegex.Regexes))
	assert.Equal(t, "\\s*['\"](.+?)['\"]", foundLanguageRegex.Regexes[0])
}
