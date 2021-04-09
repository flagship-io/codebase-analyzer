package model

import (
	"encoding/json"
	"log"
)

type LanguageRegex struct {
	ExtensionRegex string      `json:"extension_regexp"`
	FlagRegexes    []FlagRegex `json:"flag_regexes"`
}

type FlagRegex struct {
	FunctionRegex   string `json:"function_regex"`
	KeyRegex        string `json:"key_regex"`
	HasMultipleKeys bool   `json:"has_multiple_keys"`
}

var LanguageRegexes = []LanguageRegex{
	{
		ExtensionRegex: `\.[jt]sx?$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex:   `(?s)useFsModifications\(.+?\)`,
				KeyRegex:        `['"]?key['"]?\s*\:\s*['"](.+?)['"]`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)\.getModifications\(.+?\].+?\)`,
				KeyRegex:        `['"]?key['"]?\s*\:\s*['"](.+?)['"]`,
				HasMultipleKeys: true,
			},
		},
	},
	{
		ExtensionRegex: `\.go$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification(String|Number|Bool|Object|Array)\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.py$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.get_modification\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.java$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.kt$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.swift$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.m$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\]\s*getModification:@.+?\]`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
}

func AddCustomRegexes(customRegexJSON string) {
	customRegexes := []LanguageRegex{}
	err := json.Unmarshal([]byte(customRegexJSON), &customRegexes)

	if err != nil {
		log.Printf("Error when parsing custom regexes : %v", err)
		return
	}

	LanguageRegexes = append(LanguageRegexes, customRegexes...)
}
