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
	FieldRegex      string `json:"field_regex"`
	HasMultipleKeys bool   `json:"has_multiple_keys"`
}

var LanguageRegexes = []LanguageRegex{
	{
		ExtensionRegex: `\.[jt]sx?$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex:   `(?s)useFsModifications\(.+?\)`, // SDK React V2
				FieldRegex:      `['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*(['"].*['"]|[^\r\n\t\f\v ,}]+).*[},]`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)useFsFlag\(.+?\)`, // SDK React V3
				FieldRegex:      `useFsFlag[(](?:\s*['"](.*)['"]\s*,\s*(".*\s*[^"]*"|[^)]*))\s*[)]`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)\.getModifications\(.+?\].+?\)`, // SDK JS V2
				FieldRegex:      `['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*(['"].*['"]|[^\r\n\t\f\v ,}]+).*[},]`,
				HasMultipleKeys: true,
			},
			{
				FunctionRegex:   `(?s)getFlag\(.+?\)`, // SDK JS V3
				FieldRegex:      `getFlag[(](?:\s*["'](.*)["']\s*,\s*(".*\s*[^"]*"|[^)]*))\s*[)]`,
				HasMultipleKeys: true,
			},
		},
	},
	{
		ExtensionRegex: `\.go$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification(String|Number|Bool|Object|Array)\(.+?\)`, // SDK GO V2
				FieldRegex:    `\.GetModification(?:String|Number|Bool|Object|Array)\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,
			},
		},
	},
	{
		ExtensionRegex: `\.py$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.get_modification\(.+?\)`, // SDK PYTHON V2
				FieldRegex:    `\.get_modification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|True|False|\d+|"[^"]*"))?\s*\)`,
			},
		},
	},
	{
		ExtensionRegex: `\.java$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK JAVA V2
				FieldRegex:    `\.getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(.+?\)`, // SDK JAVA V3
				FieldRegex:    `\.getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,
			},
		},
	},
	{
		ExtensionRegex: `\.php$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\-\>getModification\(.+?\)`, // SDK PHP V1 && SDK PHP V2
				FieldRegex:    `\-\>getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,
			},
			{
				FunctionRegex: `(?s)\-\>getFlag\(.+?\)`, // SDK PHP V3
				FieldRegex:    `\-\>getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,
			},
		},
	},
	{
		ExtensionRegex: `\.kt$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK ANDROID V2
				FieldRegex:    `\.getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(.+?\)`, // SDK ANDROID V3
				FieldRegex:    `\.getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,
			},
		},
	},
	{
		ExtensionRegex: `\.swift$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`, // SDK iOS V2
				FieldRegex:    `\.getModification\(\s*["'](\w+)['"]\s*,\s*default(?:String|Double|Float|Int|Bool|Json|Array)\s*:\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)\s*(?:,\s*activate\s*:\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,
			},
			{
				FunctionRegex: `(?s)\.getFlag\(key: ['"](.+?)['"]`, // SDK iOS V3
				FieldRegex:    `\.getFlag[(]\s*key\s*:\s*(?:\s*["'](.*)["']\s*,\s*defaultValue\s*:\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,
			},
		},
	},
	{
		ExtensionRegex: `\.m$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\]\s*getModification:@.+?\]`, // SDK iOS V2
				FieldRegex:    `getModification\s*:\s*@\s*['"](.+?)['"](?:\s*)default(?:String|Double|Bool|Float|Int|Json|Array):\@?\s*(['"].+?['"]|YES|NO|TRUE|FALSE|true|false|[+-]?(?:\d*[.])?\d+)?`,
			},
			{
				FunctionRegex: `(?s)\s*getFlagWithKey:@.+?\]`, // SDK iOS V3
				FieldRegex:    `getFlagWithKey\s*:\s*\@['"](.+?)['"](?:\s*)['"]?defaultValue['"]?\s*\:\s*\@?\s*(.+?)\s*[\]]`,
			},
		},
	},
	{
		ExtensionRegex: `\.[fc]s$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification\(.+?\)`, // SDK .NET V1
				FieldRegex:    `\.GetModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,
			},
			{
				FunctionRegex: `(?s)\.GetFlag\(.+?\)`, // SDK .NET V3
				FieldRegex:    `\.GetFlag\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,
			},
		},
	},
	{
		ExtensionRegex: `\.vb$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification\(.+?\)`, // SDK .NET V1
				FieldRegex:    `\.GetModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|True|false|False|\d+|"[^"]*"))?\s*\)`,
			},
			{
				FunctionRegex: `(?s)\.GetFlag\(.+?\)`, // SDK .NET V3
				FieldRegex:    `\.GetFlag\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,
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
