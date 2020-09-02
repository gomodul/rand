package rand

//AN func generate random alphanumeric
func AN(length ...int) string {
	n := 6
	if len(length) > 0 {
		n = length[0]
	}
	return Generate(n, []rune("abcdefghijklmnopqrstuvwxyz0123456789"))
}
