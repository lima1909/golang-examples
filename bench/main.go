package main

var (
	s = []byte("Hello Heap")
)

func main() {
	mycopy(s)
}

func mycopy(b []byte) []byte {
	r := make([]byte, len(b))
	for i, v := range b {
		r[i] = v
	}
	return r
}
