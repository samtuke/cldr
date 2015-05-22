package ta_SG

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "ta_SG",
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
