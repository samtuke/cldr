package ru_KG

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "ru_KG",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar: calendar,
}

func init() {
	cldr.RegisterLocale(Locale)
}
