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
	FieldRegex string `json:"field_regex"`
}

var LanguageRegexes = []LanguageRegex{
	{
		ExtensionRegex: `\.[jt]sx?$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `useFsFlag[(](?:\s*['"](.*)['"]\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`, // SDK React V3
			},
			{
				FieldRegex: `['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*(['"].*['"]|[^\r\n\t\f\v,}]+).*[},]?`, // SDK JS V2 && SDK React V2
			},
			{
				FieldRegex: `getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`, // SDK JS V3
			},
		},
	},
	{
		ExtensionRegex: `\.go$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `\.GetModification(?:String|Number|Bool|Object|Array)\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK GO V2
			},
		},
	},
	{
		ExtensionRegex: `\.py$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `\.get_modification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|True|False|\d+|"[^"]*"))?\s*\)`, // SDK PYTHON V2
			},
		},
	},
	{
		ExtensionRegex: `\.java$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `\.getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK JAVA V2
			},
			{
				FieldRegex: `\.getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`, // SDK JAVA V3
			},
		},
	},
	{
		ExtensionRegex: `\.php$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `\-\>getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK PHP V1 && SDK PHP V2
			},
			{
				FieldRegex: `\-\>getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`, // SDK PHP V3
			},
		},
	},
	{
		ExtensionRegex: `\.kt$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `\.getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK ANDROID V2
			},
			{
				FieldRegex: `\.getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`, // SDK ANDROID V3
			},
		},
	},
	{
		ExtensionRegex: `\.swift$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `\.getModification\(\s*["'](\w+)['"]\s*,\s*default(?:String|Double|Float|Int|Bool|Json|Array)\s*:\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)\s*(?:,\s*activate\s*:\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK iOS V2
			},
			{
				FieldRegex: `\.getFlag[(]\s*key\s*:\s*(?:\s*["'](.*)["']\s*,\s*defaultValue\s*:\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`, // SDK iOS V3
			},
		},
	},
	{
		ExtensionRegex: `\.m$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `getModification\s*:\s*@\s*['"](.+?)['"](?:\s*)default(?:String|Double|Bool|Float|Int|Json|Array):\@?\s*(['"].+?['"]|YES|NO|TRUE|FALSE|true|false|[+-]?(?:\d*[.])?\d+)?`, // SDK iOS V2
			},
			{
				FieldRegex: `getFlagWithKey\s*:\s*\@['"](.+?)['"](?:\s*)['"]?defaultValue['"]?\s*\:\s*\@?\s*(.+?)\s*[\]]`, // SDK iOS V3
			},
		},
	},
	{
		ExtensionRegex: `\.[fc]s$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `\.GetModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK .NET V1
			},
			{
				FieldRegex: `\.GetFlag\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK .NET V3
			},
		},
	},
	{
		ExtensionRegex: `\.vb$`,
		FlagRegexes: []FlagRegex{
			{
				FieldRegex: `\.GetModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|True|false|False|\d+|"[^"]*"))?\s*\)`, // SDK .NET V1
			},
			{
				FieldRegex: `\.GetFlag\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK .NET V3
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
