package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("hel")

	if trie.Search("hell") {
		t.Error("excepted false but got true")
		return
	}

	if !trie.Search("hel") {
		t.Error("excepted true but got false")
		return
	}

	trie.Insert("河北")
	trie.Insert("河南")
	trie.Insert("湖南")

	actual := []bool{
		trie.Search("河西"),
		trie.Search("河南"),
		trie.Search("湖北"),
		trie.Search("湖南"),
	}

	want := []bool{
		false,
		true,
		false,
		true,
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != want[i] {
			t.Errorf("want %v but got %v, the index is %d\n", want[i], actual[i], i)
			return
		}

	}

}
