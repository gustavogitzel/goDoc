package cpfcnpj

import (
	"godoc/internal/utils"
	"regexp"
	"strconv"
)

// Regexp pattern for CPF and CNPJ.
var (
	CPFRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

// IsCPF verifies whether the provided string represents a valid CPF document.
func IsCPF(doc string) bool {
	const (
		size     = 9
		position = 10
	)

	return isCPFOrCNPJ(doc, CPFRegexp, size, position)
}

// IsCNPJ verifies whether the provided string represents a valid CNPJ document.
func IsCNPJ(doc string) bool {
	const (
		size     = 12
		position = 5
	)

	return isCPFOrCNPJ(doc, CNPJRegexp, size, position)
}

func isCPFOrCNPJ(doc string, pattern *regexp.Regexp, size int, position int) bool {
	if !pattern.MatchString(doc) {
		return false
	}

	utils.CleanNonDigits(&doc)

	if isAllEquals(doc) {
		return false
	}

	aux := doc[:size]
	digit := calculateDigit(aux, position)

	aux += digit
	digit = calculateDigit(aux, position+1)

	return doc == aux+digit
}

func isAllEquals(doc string) bool {
	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

func calculateDigit(doc string, position int) string {
	var sum int
	for _, character := range doc {
		sum += utils.ToInt(character) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11

	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}
