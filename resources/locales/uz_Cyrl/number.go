package uz_Cyrl

import "github.com/theplant/i18n/cldr"

var (
	symbols = cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "-", Percent: "٪", PerMille: "؉"}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0%"}
)
