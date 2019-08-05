package sliceof

func SliceOfStringConcat(xss [][]string) []string  {
	out := make([]string, 0)
	for _, xs := range xss {
		for _, x := range xs {
			out = append(out, x)
		}
	}
	return out
}