package model

import (
	"encoding/json"
	"log"
)

type LanguageRegex struct {
	ExtensionRegex string   `json:"extension_regex"`
	FlagRegexes    []string `json:"flag_regexes"`
}

var LanguageRegexes = []LanguageRegex{
	{
		ExtensionRegex: `\.[jt]sx?$`,
		FlagRegexes: []string{
			`useFsFlag[(](?:\s*['"](.*)['"]\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                        // SDK React V3
			`['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*(['"].*['"]|[^\r\n\t\f\v,}]+).*[},]?`, // SDK JS V2 && SDK React V2
			`getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                          // SDK JS V3
		},
	},

	{
		ExtensionRegex: `\.go$`,
		FlagRegexes: []string{
			`\.GetModification(?:String|Number|Bool|Object|Array)\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK GO V2
		},
	},
	{
		ExtensionRegex: `\.py$`,
		FlagRegexes: []string{
			`\.get_modification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|True|False|\d+|"[^"]*"))?\s*\)`, // SDK PYTHON V2
		},
	},
	{
		ExtensionRegex: `\.java$`,
		FlagRegexes: []string{
			`\.getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK JAVA V2
			`\.getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                                                             // SDK JAVA V3
		},
	},
	{
		ExtensionRegex: `\.php$`,
		FlagRegexes: []string{
			`\-\>getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK PHP V1 && SDK PHP V2
			`\-\>getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                                                             // SDK PHP V3
		},
	},
	{
		ExtensionRegex: `\.kt$`,
		FlagRegexes: []string{
			`\.getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK ANDROID V2
			`\.getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                                                             // SDK ANDROID V3

		},
	},
	{
		ExtensionRegex: `\.swift$`,
		FlagRegexes: []string{
			`\.getModification\(\s*["'](\w+)['"]\s*,\s*default(?:String|Double|Float|Int|Bool|Json|Array)\s*:\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)\s*(?:,\s*activate\s*:\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK iOS V2
			`\.getFlag[(]\s*key\s*:\s*(?:\s*["'](.*)["']\s*,\s*defaultValue\s*:\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                                                                                                 // SDK iOS V3
		},
	},
	{
		ExtensionRegex: `\.m$`,
		FlagRegexes: []string{
			`getModification\s*:\s*@\s*['"](.+?)['"](?:\s*)default(?:String|Double|Bool|Float|Int|Json|Array):\@?\s*(['"].+?['"]|YES|NO|TRUE|FALSE|true|false|[+-]?(?:\d*[.])?\d+)?`, // SDK iOS V2
			`getFlagWithKey\s*:\s*\@['"](.+?)['"](?:\s*)['"]?defaultValue['"]?\s*\:\s*\@?\s*(.+?)\s*[\]]`,                                                                            // SDK iOS V3
		},
	},
	{
		ExtensionRegex: `\.[fc]s$`,
		FlagRegexes: []string{
			`\.GetModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK .NET V1
			`\.GetFlag\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,         // SDK .NET V3
		},
	},
	{
		ExtensionRegex: `\.vb$`,
		FlagRegexes: []string{
			`\.GetModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|True|false|False|\d+|"[^"]*"))?\s*\)`, // SDK .NET V1
			`\.GetFlag\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,                    // SDK .NET V3
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
