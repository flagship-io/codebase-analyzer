package model

type LanguageRegex struct {
	ExtensionRegexp string
	FlagRegexes     []FlagRegex
}

type FlagRegex struct {
	FunctionRegex   string
	KeyRegex        string
	HasMultipleKeys bool
}

var LanguageRegexes = []LanguageRegex{
	{
		ExtensionRegexp: `\.jsx?$`,
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
		ExtensionRegexp: `\.go$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.GetModification(String|Number|Bool|Object|Array)\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegexp: `\.py$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.get_modification\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegexp: `\.java$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegexp: `\.kt$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegexp: `\.swift$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\.getModification\(.+?\)`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
	{
		ExtensionRegexp: `\.m$`,
		FlagRegexes: []FlagRegex{
			{
				FunctionRegex: `(?s)\]\s*getModification:@.+?\]`,
				KeyRegex:      `\s*['"](.+?)['"]`,
			},
		},
	},
}
