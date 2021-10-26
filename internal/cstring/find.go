package cstring

func FindLastAfterSubstr(s, substr string) string {
	if substr == "" {
		return ""
	}
	start := len(s) - len(substr)
	end := len(s)
	for start >= 0 {
		if s[start:end] == substr {
			return s[end:]
		}
		start--
		end--
	}

	return ""
}
