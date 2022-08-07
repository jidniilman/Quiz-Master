package utils

// NormalizeQuotes taking inspiration and insight from: https://stackoverflow.com/a/36361482 (John Weldon)
func NormalizeQuotes(r rune) rune {
	switch r {
	case '“', '‹', '”', '›':
		return '"'
	case '‘', '’':
		return '\''
	}
	return r
}
