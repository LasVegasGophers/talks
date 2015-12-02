import (
	"regexp"
	"strings"
)

func CharParser(search byte, terminal string) Parser {
	return func(input string) (ParseNode, string) {
		if len(input) > 0 && input[0] == search {
			return terminal, input[1:]
		} else {
			return nil, input
		}
	}
}

func RegexpParser(matcher *regexp.Regexp) Parser {
	return func(input string) (ParseNode, string) {
		matchStr := matcher.FindString(input)
		if matchStr == "" {
			return nil, input
		}
		return matchStr, strings.Replace(input, matchStr, "", 1)
	}
}