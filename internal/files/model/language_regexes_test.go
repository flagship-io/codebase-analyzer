package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
)

func TestCustomRegex(t *testing.T) {
	AddCustomRegexes(`[{"extension_regex":".tsx?","flag_regexes":[{"function_regex":"(?s)useFlagShipActivation\\(.+?\\)","key_regex":"\\s*['\"](.+?)['\"]","has_multiple_keys":true}]}]`)

	found := funk.Find(LanguageRegexes, func(languageRegex LanguageRegex) bool {
		return languageRegex.ExtensionRegex == ".tsx?"
	})

	assert.NotNil(t, found)
	foundLanguageRegex, ok := found.(LanguageRegex)

	assert.True(t, ok)
	assert.NotNil(t, foundLanguageRegex)

	assert.Equal(t, ".tsx?", foundLanguageRegex.ExtensionRegex)
	assert.Equal(t, 1, len(foundLanguageRegex.FlagRegexes))
	assert.Equal(t, "(?s)useFlagShipActivation\\(.+?\\)", foundLanguageRegex.FlagRegexes[0].FunctionRegex)
	assert.Equal(t, "\\s*['\"](.+?)['\"]", foundLanguageRegex.FlagRegexes[0].KeyRegex)
	assert.Equal(t, true, foundLanguageRegex.FlagRegexes[0].HasMultipleKeys)
}
