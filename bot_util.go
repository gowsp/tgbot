package tgbot

import "regexp"

// characters '_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!' must be escaped with the preceding character '\'.
func EscapeMd(input string) func(string) string {
	specialCharacters := `*_[]()~` + "`" + `>#+-=|{}.!`
	re := regexp.MustCompile(`[` + regexp.QuoteMeta(specialCharacters) + `]`)
	return func(input string) string {
		return re.ReplaceAllStringFunc(input, func(match string) string {
			return `\` + match
		})
	}
}
