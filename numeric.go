package rand

//N func generate random numeric
func N(length ...int) string {
	n := 6
	if len(length) > 0 {
		n = length[0]
	}
	return Generate(n, []rune("0123456789"))
}
