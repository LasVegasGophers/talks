import (
	"strconv"

	"github.com/rayman22201/goParserCombinate/parserCombinator"
)

func urlParser(input string) (parserCombinator.ParseNode, string) {
	populateURLStruct := func(nodes ...parserCombinator.ParseNode) parserCombinator.ParseNode {
		url := URLStruct{}
		url.Protocol = nodes[0].(string)

		switch port := nodes[2].(type) {
		case int:
			url.Port = strconv.Itoa(port)
		}

		switch base := nodes[1].(type) {
		case []string:
			url.Base = base
		case string:
			if len(base) > 0 {
				url.Base = []string{base}
			}
		}
		url.Query = nodes[5].(string)
		url.Anchor = nodes[4].(string)

		switch pathSegments := nodes[3].(type) {
		case []string:
			url.Path = pathSegments[0]
			if len(pathSegments) > 1 {
				url.Filename = pathSegments[1]
			}
			if len(pathSegments) > 2 {
				url.FileExt = pathSegments[2]
			}
		}

		if len(url.Base) == 0 && url.Filename == "" && (url.Path == "" || url.Path == "/") {
			return nil
		}

		return url
	}

	//url := [ protocol ] [ ipV6 | ipV4 | domain ] [ port ] [ path ] [ queryString | anchorTag ]
	absURL := parserCombinator.Or(passFirstNode, ipv6, ipv4, domain)
	isUrl := parserCombinator.And(
		populateURLStruct,
		maybe(protocol),
		maybe(absURL),
		maybe(port),
		maybe(path),
		maybe(anchorString),
		maybe(queryString))

	return isUrl(input)
}
