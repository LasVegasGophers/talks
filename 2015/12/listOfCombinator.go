func ListOf(processResults Nodify, parse Parser, delimeter Parser) Parser {
	return func(input string) (ParseNode, string) {
		var results []ParseNode
		var output string
		var delimCount int
		output = input

		for {
			if len(output) == 0 {
				break
			}
			var result, delim ParseNode
			result, output = parse(output)
			if result != nil {
				results = append(results, result)
			} else {
				break
			}
			if len(output) == 0 {
				break
			}
			delim, output = delimeter(output)
			if delim == nil {
				break
			} else {
				delimCount++
			}
		}
		if len(results) > 0 && delimCount == (len(results)-1) {
			return processResults(results...), output
		} else {
			return nil, input
		}
	}
}