package options

func replaceOnEmpty(val, repl string) string {
	if val == "" {
		return repl
	}

	return val
}

func getOrEmpty(m map[string]any, key string) string {
	val, ok := m[key].(string)
	if !ok {
		return ""
	}

	return val
}
