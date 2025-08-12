package utility

// var name, link string
// unpack(strings.Split("foo:bar", ":"), &name, &link)
func Unpack(s []string, vars ...*string) {
	n := len(s)
	if len(vars) < n {
		n = len(vars)
	}
	for i := 0; i < n; i++ {
		if vars[i] == nil {
			continue
		}
		*vars[i] = s[i]
	}
}
