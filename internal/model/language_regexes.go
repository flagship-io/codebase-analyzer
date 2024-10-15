package model

import (
	"encoding/json"
	"log"
)

type LanguageRegex struct {
	FileExtension   string   `json:"file_extension"`
	Regexes         []string `json:"regexes"`
	Split           bool     `json:"split"`
	ForFlag         bool     `json:"search_flag"`
	ForDefaultValue bool     `json:"search_default_value"`
}

var LanguageRegexes = []LanguageRegex{
	{
		FileExtension: `\.[jt]sx?$`,
		Regexes: []string{
			`useFsFlag[(](?:(?:\s*['"](.*)['"]\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                   // SDK React V3
			`useFsFlag[(](?:(?:\s*["'](.*)["']\s*[)]\s*.getValue[(](["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                     // SDK React V4
			`['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*(['"].*['"]|[^\r\n\t\f\v,}]+).*[},]?`, // SDK JS V2 && SDK React V2
			`getFlag[(](?:(?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                     // SDK JS V3
			`getFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*.getValue[(](["']\w*\s*[^"]*["']|[^)]*))\s*[)])?`,                     // SDK JS V4
		},
	},
	{
		FileExtension: `\.[jt]sx?$`,
		Regexes: []string{
			`(?:(\w+)\s*[?]?\s*[:]?\s*(?:number|string|boolean|any|void| never|null|undefined|bigint|symbol|object|IFSFlag|FSFlag)?)\s*[=]\s*(?:.*)useFsFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*))[^\.]`, // SDK React V4 Key
			`(?:(\w+)\s*[?]?\s*[:]?\s*(?:number|string|boolean|any|void| never|null|undefined|bigint|symbol|object|IFSFlag|FSFlag)?)\s*[=]\s*(?:.*)getFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*))[^\.]`,   // SDK JS V4 Key
		},
		Split:   true,
		ForFlag: true,
	},
	{
		FileExtension: `\.[jt]sx?$`,
		Regexes: []string{
			`\s*(\w*)[\.]getValue[(](["']?\w*["']?)[)]`, // SDK JS & React V4 Default value
		},
		Split:           true,
		ForDefaultValue: true,
	},
	{
		FileExtension: `\.go$`,
		Regexes: []string{
			`\.GetModification(?:String|Number|Bool|Object|Array)\((?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`, // SDK GO V2
		},
	},
	{
		FileExtension: `\.py$`,
		Regexes: []string{
			`\.get_modification\((?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|True|False|\d+|"[^"]*"))?\s*\))?`, // SDK PYTHON V2
		},
	},
	{
		FileExtension: `\.java$`,
		Regexes: []string{
			`\.getModification\(\s*(?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`, // SDK JAVA V2
			`\.getFlag[(](?:\s*(?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                                             // SDK JAVA V3
		},
	},
	{
		FileExtension: `\.kt$`,
		Regexes: []string{
			`\.getModification\(\s*(?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`, // SDK ANDROID V2
			`\.getFlag[(](?:\s*(?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                                             // SDK ANDROID V3
			`getFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*.value[(](["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                                      // SDK ANDROID V4

		},
	},
	{
		FileExtension: `\.kt$`,
		Regexes: []string{
			`(?:(\w+))\s*[=]\s*(?:.*)getFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*))[^\.]`, // SDK ANDROID V4 Key
		},
		Split:   true,
		ForFlag: true,
	},
	{
		FileExtension: `\.kt$`,
		Regexes: []string{
			`\s*(\w*)[\.]value[(](["']?\w*["']?)[)]`, // SDK ANDROID V4 Default value
		},
		Split:           true,
		ForDefaultValue: true,
	},
	{
		FileExtension: `\.php$`,
		Regexes: []string{
			`\-\>getModification\(\s*(?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`, // SDK PHP V1 && SDK PHP V2
			`\-\>getFlag[(](?:\s*(?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                                             // SDK PHP V3
			`\-\>getFlag[(](?:\s*(?:\s*["'](\w*)["']\s*[)]\s*\-\>\s*getValue[(](["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                        // SDK PHP V4
		},
	},
	{
		FileExtension: `\.php$`,
		Regexes: []string{
			`(?:(\w+))\s*[=]\s*(?:.*)getFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*))[^\-]`, // SDK PHP V4 Key
		},
		Split:   true,
		ForFlag: true,
	},
	{
		FileExtension: `\.php$`,
		Regexes: []string{
			`\s*(\w*)\s*\-\>\s*getValue[(](["']?\w*["']?)[)]`, // SDK PHP V4 Default value
		},
		Split:           true,
		ForDefaultValue: true,
	},
	{
		FileExtension: `\.swift$`,
		Regexes: []string{
			`\.getModification\((?:\s*["'](\w+)['"]\s*,\s*default(?:String|Double|Float|Int|Bool|Json|Array)\s*:\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)\s*(?:,\s*activate\s*:\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`, // SDK iOS V2
			`\.getFlag[(](?:\s*key\s*:\s*(?:\s*["'](.*)["']\s*,\s*defaultValue\s*:\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                                                                                 // SDK iOS V3
			`getFlag[(](?:\s*key\s*[:]\s*(?:\s*["'](\w*)["']\s*[)]\s*\.value[(]\s*defaultValue\s*[:]\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                                                               // SDK iOS V4
		},
	},
	{
		FileExtension: `\.swift$`,
		Regexes: []string{
			`(?:(\w+))\s*[=]\s*(?:.*)getFlag[(](?:(?:\s*key\s*[:]\s*["'](\w*)["']\s*[)]\s*))[^\.]`, // SDK iOS V4 Key
		},
		Split:   true,
		ForFlag: true,
	},
	{
		FileExtension: `\.swift$`,
		Regexes: []string{
			`\s*(\w*)[\.]value[(]\s*defaultValue\s*[:]\s*(["']?\w*["']?)[)]`, // SDK iOS V4 Default value
		},
		Split:           true,
		ForDefaultValue: true,
	},
	{
		FileExtension: `\.m$`,
		Regexes: []string{
			`getModification\s*:(?:\s*\@?\s*['"](.+?)['"](?:\s*)default(?:String|Double|Bool|Float|Int|Json|Array):\@?\s*(['"].+?['"]|YES|NO|TRUE|FALSE|true|false|[+-]?(?:\d*[.])?\d+)?)?`, // SDK iOS V2
			`getFlagWithKey\s*:(?:\s*\@?['"](.+?)['"](?:\s*)['"]?defaultValue['"]?\s*\:\s*\@?\s*(.+?)\s*[\]])?`,                                                                             // SDK iOS V3
		},
	},
	{
		FileExtension: `\.[fc]s$`,
		Regexes: []string{
			`\.GetModification\((?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`, // SDK .NET V1
			`\.GetFlag\((?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`,         // SDK .NET V3
			`GetFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*.GetValue[(](["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                                // SDK .NET V4
		},
	},
	{
		FileExtension: `\.[fc]s$`,
		Regexes: []string{
			`(?:(\w+))\s*[=]\s*(?:.*)GetFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*))[^\.]`, // SDK .NET V4 Key
		},
		Split:   true,
		ForFlag: true,
	},
	{
		FileExtension: `\.[fc]s$`,
		Regexes: []string{
			`\s*(\w*)[\.]GetValue[(](["']?\w*["']?)[)]`, // SDK .NET V4 Default value
		},
		Split:           true,
		ForDefaultValue: true,
	},
	{
		FileExtension: `\.vb$`,
		Regexes: []string{
			`\.GetModification\((?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|True|false|False|\d+|"[^"]*"))?\s*\))?`, // SDK .NET V1
			`\.GetFlag\((?:\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`,                    // SDK .NET V3
			`GetFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*.GetValue[(](["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                                                           // SDK .NET V4
		},
	},
	{
		FileExtension: `\.dart$`,
		Regexes: []string{
			`getFlag[(](?:(?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,             // SDK Flutter V1-V2-V3
			`getFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*.value[(](["'].*\s*[^"]*["']|[^)]*))\s*[)])?`, // SDK Flutter V4
		},
	},
	{
		FileExtension: `\.dart$`,
		Regexes: []string{
			`(?:(\w+))\s*[=]\s*(?:.*)getFlag[(](?:(?:\s*["'](\w*)["']\s*[)]\s*))[^\.]`, // SDK Flutter V4 Key
		},
		Split:   true,
		ForFlag: true,
	},
	{
		FileExtension: `\.dart$`,
		Regexes: []string{
			`\s*(\w*)[\.]value[(](["']?\w*["']?)[)]`, // SDK Flutter V4 Default value
		},
		Split:           true,
		ForDefaultValue: true,
	},
}

func AddCustomRegexes(customRegexJSON string) {
	regexes := []LanguageRegex{}
	err := json.Unmarshal([]byte(customRegexJSON), &regexes)

	if err != nil {
		log.Printf("Error when parsing custom regexes : %v", err)
		return
	}

	LanguageRegexes = append(LanguageRegexes, regexes...)
}
