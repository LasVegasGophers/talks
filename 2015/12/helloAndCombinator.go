func And(processResults Nodify, parsers ...Parser) Parser {
	return func(input string) (ParseNode, string) {
		var resultNodes []ParseNode
		var output string

		output = input
		for _, curParser := range parsers {
			var result ParseNode
			result, output = curParser(output)
			// If we didn't match even one input, we failed.
			if result == nil {
				return nil, input
			} else {
				resultNodes = append(resultNodes, result)
				input = output
			}
		}
		if len(resultNodes) != len(parsers) {
			return nil, input
		}
		return processResults(resultNodes...), output
	}
}