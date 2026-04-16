func isPalindrome(s string) bool {
	strings.ReplaceAll(s, " ", "")
	
	l := 0
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	result := re.ReplaceAllString(s, "")
	r := strings.ToLower(result)
	ri := len(r) - 1
	for i := l; i <= ri; i++ {

		if r[l] == r[ri] {
			l++
			ri--
		} else {
			return false
		}
	}
	return true
}
