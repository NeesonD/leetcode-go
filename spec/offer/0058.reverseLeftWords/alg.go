package _058_reverseLeftWords

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
