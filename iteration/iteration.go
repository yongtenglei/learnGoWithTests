package iteration

func Repeat(s string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += s
	}
	return
}
