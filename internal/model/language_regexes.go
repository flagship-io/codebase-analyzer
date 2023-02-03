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
	FieldRegex      string `json:"key_regex"`
	HasMultipleKeys bool   `json:"has_multiple_keys"`
}

var LanguageRegexes = []LanguageRegex{
	{
		ExtensionRegex: `\.[jt]sx?$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex:   `(?s)useFsModifications\(.+?\)`, // SDK React V2
				FieldRegex:      `['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*['"]?(.+?)['"]?\s*[\"\,]`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)useFsFlag\(.+?\)`, // SDK React V3
				FieldRegex:      `useFsFlag\(['"]?\s*(.+?)['"](?:.\s*)['"]?(.+?)['"]?\s*[\"\)]`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)\.getModifications\(.+?\].+?\)`, // SDK JS V2
				FieldRegex:      `['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*['"]?(.+?)['"]?\s*[\"\,]`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)getFlag\(.+?\)`, // SDK JS V3
				FieldRegex:      `getFlag\(['"]?\s*(.+?)['"](?:.\s*)['"]?(.+?)['"]?\s*[\"\)\,]`,
				HasMultipleKeys: true,
			},
		},
	},
	{
		ExtensionRegex: `\.go$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification(String|Number|Bool|Object|Array)\(.+?\)`, // SDK GO V2
				FieldRegex:    `\s*['"](.+?)['"](?:,\s*)['"]?(.+?)['"]?\s*[\,]`,
			},
		},
	},
	{
		ExtensionRegex: `\.py$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.get_modification\(.+?\)`, // SDK PYTHON V2
				FieldRegex:    `\s*['"](.+?)['"](?:,\s*)['"]?(.+?)['"]?\s*[\)\,]`,
			},
		},
	},
	{
		ExtensionRegex: `\.java$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK JAVA V2
				FieldRegex:    `\s*['"](.+?)['"](?:,\s*)['"]?(.+?)['"]?\s*[\)\,]`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(.+?\)`, // SDK JAVA V3
				FieldRegex:    `(?s)\.getFlag\(['"](.+?)['"](?:.\s*)['"]?(.+?)['"]?\s*[\"\)\,]`,
			},
		},
	},
	{
		ExtensionRegex: `\.php$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\-\>getModification\(.+?\)`, // SDK PHP V2
				FieldRegex:    `\s*['"](.+?)['"](?:,\s*)['"]?(.+?)['"]?\s*[\)\,]`,
			},
			{
				FunctionRegex: `(?s)\-\>getFlag\(.+?\)`, // SDK PHP V3
				FieldRegex:    `(?s)\-\>getFlag\(['"](.+?)['"](?:.\s*)['"]?(.+?)['"]?\s*[\"\)\,]`,
			},
		},
	},
	{
		ExtensionRegex: `\.kt$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK ANDROID V2
				FieldRegex:    `\s*['"](.+?)['"](?:,\s*)['"]?(.+?)['"]?\s*[\)\,]`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(.+?\)`, // SDK ANDROID V3
				FieldRegex:    `(?s)\.getFlag\(['"](.+?)['"](?:.\s*)['"]?(.+?)['"]?\s*[\"\)\,]`,
			},
		},
	},
	{
		ExtensionRegex: `\.swift$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK iOS V2
				FieldRegex:    `\s*['"](.+?)['"](?:,\s*)['"]?default(?:String|Double|Bool|Float|Int|Json|Array)['"]?\s*\:\s*['"]?(.+?)['"]?\s*[\"\,]`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(key: ['"](.+?)['"]`, // SDK iOS V3
				FieldRegex:    `['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*['"]?(.+?)['"]?\s*[\)]`,
			},
		},
	},
	{
		ExtensionRegex: `\.m$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\]\s*getModification:@.+?\]`, // SDK iOS V2
				FieldRegex:    `\s*['"](.+?)['"]`,
			},
			{
				FunctionRegex: `(?s)\s*getFlagWithKey:@.+?\]`, // SDK iOS V3
				FieldRegex:    `\s*getFlagWithKey:@['"](.+?)['"](?:\s*)['"]?defaultValue['"]?\s*\:\@?\s*['"]?(.+?)['"]?\s*[\]]`,
			},
		},
	},
	{
		ExtensionRegex: `\.[fc]s$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification\(.+?\)`, // SDK .NET V1
				FieldRegex:    `(?s)\.GetModification\(['"](.+?)['"](?:,\s*)['"]?(.+?)['"]?\s*[\)\,]`,
			},
			{
				FunctionRegex: `(?s)\.GetFlag\(.+?\)`, // SDK .NET V3
				FieldRegex:    `(?s)\.GetFlag\(['"](.+?)['"](?:.\s*)['"]?(.+?)['"]?\s*[\"\)\,]`,
			},
		},
	},
	{
		ExtensionRegex: `\.vb$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification\(.+?\)`, // SDK .NET V1
				FieldRegex:    `(?s)\.GetModification\(['"](.+?)['"](?:,\s*)['"]?(.+?)['"]?\s*[\)\,]`,
			},
			{
				FunctionRegex: `(?s)\.GetFlag\(.+?\)`, // SDK .NET V3
				FieldRegex:    `(?s)\.GetFlag\(['"](.+?)['"](?:,\s*)['"]?(.+?)['"]?\s*[\)\,]`,
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
