package _005_replaceSpace

func replaceSpace(s string) string {
	var result string
	tmp := "%20"
	for _, i := range s {
		if string(i) == " " {
			result += tmp
		} else {
			result += string(i)
		}
	}
	return result
}
