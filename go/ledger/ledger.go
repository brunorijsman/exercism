// Go exercise ledger
package ledger

// Refactoring steps:
// 1. Move sorting of entries into separate function sortEntries (as is)
// 2. Move conversion of Entry to string into a separate function entryToString (as is)
// 3. Remove useless parallelism and channel
// 4. Move formating of header into separate function localizeHeader (as is)
// 5. Use sort.Slices and comparison function entryLessThan to sort entries.
// 6. Create separate localizeDate function
// 7. Use fmt.Sprintf where appropriate
// 8. Non-empty error strings
// 9. Clearer variable names
// 10. Remove unneeded temporary array ss
// 11. Simplify localization and formating of price

import (
	"errors"
	"fmt"
	"sort"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type locale struct {
	header              string
	dateFormat          string
	fraction            rune
	thousand            rune
	positivePriceFormat string
	negativePriceFormat string
}

var locales = map[string]locale{
	"en-US": {
		header:              fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change"),
		dateFormat:          "%[2]s/%[1]s/%[3]s",
		fraction:            '.',
		thousand:            ',',
		positivePriceFormat: "%s%s ",
		negativePriceFormat: "(%s%s)",
	},
	"nl-NL": {
		header:              fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering"),
		dateFormat:          "%s-%s-%s",
		fraction:            ',',
		thousand:            '.',
		positivePriceFormat: "%s %s ",
		negativePriceFormat: "%s %s-",
	},
}

func entryLessThan(entry1, entry2 Entry) bool {
	switch {
	case entry1.Date < entry2.Date:
		return true
	case entry1.Date > entry2.Date:
		return false
	case entry1.Description < entry2.Description:
		return true
	case entry1.Description > entry2.Description:
		return false
	case entry1.Change < entry2.Change:
		return true
	default:
		return false
	}
}

func sortEntries(entries []Entry) []Entry {
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sort.Slice(entriesCopy, func(i, j int) bool {
		return entryLessThan(entriesCopy[i], entriesCopy[j])
	})
	return entriesCopy
}

func localizeDate(date string, locale locale) (string, error) {
	if len(date) != 10 {
		return "", errors.New("Invalid date")
	}
	year, dash1, month, dash2, day := date[0:4], date[4], date[5:7], date[7], date[8:10]
	if dash1 != '-' || dash2 != '-' {
		return "", errors.New("Invalid date")
	}
	return fmt.Sprintf(locale.dateFormat, day, month, year), nil
}

func formatDescription(description string) string {
	if len(description) > 25 {
		return fmt.Sprintf("%-22s...", description[:22])
	}
	return fmt.Sprintf("%-25s", description)
}

func currencySymbol(currency string) (string, error) {
	switch currency {
	case "USD":
		return "$", nil
	case "EUR":
		return "€", nil
	default:
		return "", errors.New("Unknown currency")
	}
}

func localizePrice(cents int, currency string, locale locale) (string, error) {
	negative := false
	if cents < 0 {
		cents = cents * -1
		negative = true
	}
	priceStr := fmt.Sprintf("%c%02d", locale.fraction, cents%100)
	cents /= 100
	for cents > 1000 {
		priceStr = fmt.Sprintf("%c%03d", locale.thousand, cents%1000) + priceStr
		cents /= 1000
	}
	priceStr = fmt.Sprintf("%d", cents) + priceStr
	currencyStr, err := currencySymbol(currency)
	if err != nil {
		return "", err
	}
	if negative {
		priceStr = fmt.Sprintf(locale.negativePriceFormat, currencyStr, priceStr)
	} else {
		priceStr = fmt.Sprintf(locale.positivePriceFormat, currencyStr, priceStr)
	}
	return priceStr, nil
}

func entryToString(i int, entry Entry, currency string, locale locale) (string, error) {
	dateStr, err := localizeDate(entry.Date, locale)
	if err != nil {
		return "", err
	}
	descriptionStr := formatDescription(entry.Description)
	if err != nil {
		return "", err
	}
	priceStr, err := localizePrice(entry.Change, currency, locale)
	if err != nil {
		return "", err
	}
	entryStr := fmt.Sprintf("%10s | %25s | %13s\n", dateStr, descriptionStr, priceStr)
	return entryStr, nil
}

func FormatLedger(currency string, localeStr string, entries []Entry) (string, error) {
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}
	locale, found := locales[localeStr]
	if !found {
		return "", errors.New("Unknown locale")
	}
	entriesCopy := sortEntries(entries)
	output := locale.header
	for i, entry := range entriesCopy {
		entryStr, err := entryToString(i, entry, currency, locale)
		if err != nil {
			return "", err
		}
		output += entryStr
	}
	return output, nil
}
