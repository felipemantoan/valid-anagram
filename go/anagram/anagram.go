package anagram

import (
	"bytes"
	"sort"
)

func IsValid(s string, t string) bool {
	bfa := []byte(s)
	bfb := []byte(t)
	bytesSort(&bfa)
	bytesSort(&bfb)
	return bytes.Equal(bfa, bfb)
}

func bytesSort(bytes *[]byte) {
	sort.Slice(*bytes, func (a, b int) bool {
		return (*bytes)[a] < (*bytes)[b]
	})
}