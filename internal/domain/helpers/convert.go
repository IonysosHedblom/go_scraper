package helpers

func StringSliceToByteSlice(in []string) ([]byte) {
	var out []byte
	for _, s := range in {
		out = append(out, []byte(s)...)
	}
	return out
}