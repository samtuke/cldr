package bn

import "github.com/theplant/i18n/cldr"

var (
	symbols = cldr.Symbols{}
	formats = cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "#,##,##0.00¤", Percent: "#,##,##0%"}
)
