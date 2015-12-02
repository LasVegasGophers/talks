func Or(processResults Nodify, parsers ...Parser) Parser {
	return func(input string) (ParseNode, string) {
		var output string

		output = input
		for _, curParser := range parsers {
			// If we encountered the end of the input before matching all options, we failed.
			if len(output) == 0 {
				return nil, input
			}

			var result ParseNode
			result, output = curParser(output)
			// If we match even one, then we succeeded
			if result != nil {
				return processResults(result), output
			}
		}

		// If we didn't match any of the results, we failed.
		return nil, input
	}
}