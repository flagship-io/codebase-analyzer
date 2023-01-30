package model

import (
	"encoding/json"
	"log"
)

type LanguageRegex struct {
	ExtensionRegex string      `json:"extension_regex"`
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
				FunctionRegex:   `(?s)useFsFlag\(.+?\)`, // SDK React V3
				KeyRegex:        `useFsFlag\(['"]?\s*(.+?)['"]\,`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)useFsModifications\(.+?\)`, // SDK React V2
				KeyRegex:        `['"]?key['"]?\s*\:\s*['"](.+?)['"]`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)getFlag\(.+?\)`, // SDK JS V3
				KeyRegex:        `getFlag\(['"]?\s*(.+?)['"]\,`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)\.getModifications\(.+?\].+?\)`, // SDK JS V2
				KeyRegex:        `['"]?key['"]?\s*\:\s*['"](.+?)['"]`,
				HasMultipleKeys: true,
			},
		},
	},
	{
		ExtensionRegex: `\.go$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification(String|Number|Bool|Object|Array)\(.+?\)`, // SDK GO V2
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.py$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.get_modification\(.+?\)`, // SDK PYTHON V2
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.java$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK JAVA V2
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(.+?\)`, // SDK JAVA V3
				KeyRegex:      `(?s)\.getFlag\(['"](.+?)['"],`,
			},
		},
	},
	{
		ExtensionRegex: `\.php$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\-\>getModification\(.+?\)`, // SDK JAVA V2
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
			{
				FunctionRegex: `(?s)\-\>getFlag\(.+?\)`, // SDK JAVA V3
				KeyRegex:      `(?s)\-\>getFlag\(['"](.+?)['"],`,
			},
		},
	},
	{
		ExtensionRegex: `\.kt$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK ANDROID V2
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(.+?\)`, // SDK ANDROID V3
				KeyRegex:      `(?s)\.getFlag\(['"](.+?)['"],`,
			},
		},
	},
	{
		ExtensionRegex: `\.swift$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK iOS V2
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(key: ['"](.+?)['"]`, // SDK iOS V3
				KeyRegex:      `['"]?key['"]?\s*\:\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.m$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\]\s*getModification:@.+?\]`, // SDK iOS V2
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
			{
				FunctionRegex: `(?s)\s*getFlagWithKey:@.+?\]`, // SDK iOS V3
				KeyRegex:      `\s*getFlagWithKey:@['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegex: `\.[fc]s$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification\(.+?\)`, // SDK .NET V1
				KeyRegex:      `(?s)\.GetModification\(['"](.+?)['"],`,
			},
			{
				FunctionRegex: `(?s)\.GetFlag\(.+?\)`, // SDK .NET V3
				KeyRegex:      `(?s)\.GetFlag\(['"](.+?)['"],`,
			},
		},
	},
	{
		ExtensionRegex: `\.vb$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification\(.+?\)`, // SDK .NET V1
				KeyRegex:      `(?s)\.GetModification\(['"](.+?)['"],`,
			},
			{
				FunctionRegex: `(?s)\.GetFlag\(.+?\)`, // SDK .NET V3
				KeyRegex:      `(?s)\.GetFlag\(['"](.+?)['"],`,
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
