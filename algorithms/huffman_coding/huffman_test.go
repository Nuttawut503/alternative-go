package huffman_coding_test

import (
	huffman "app/main/algorithms/huffman_coding"
	"reflect"
	"testing"
)

func TestHuffman(t *testing.T) {
	t.Run("Check if huffman tree builds correctly", func(t *testing.T) {
		text := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
		freq := []int{5, 9, 12, 13, 16, 45}
		got := huffman.BuildHuffman(text, freq)
		expect := map[byte]string{
			'a': "1100",
			'b': "1101",
			'c': "100",
			'd': "101",
			'e': "111",
			'f': "0",
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})
}
