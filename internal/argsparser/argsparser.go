package argsparser

func Parse(args []string, keys []string) map[string]string {

	result := make(map[string]string)

	for k, v := range keys {
		if len(args) > k {
			result[v] = args[k]
		} else {
			result[v] = ""
		}
	}

	return result
}
