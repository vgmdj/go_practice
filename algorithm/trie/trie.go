package trie

// DefaultCap max cap all ascii code
const DefaultCap = 26

// Trie the data structure
type Trie struct {
	next   map[rune]*Trie
	isWord bool
}

// NewTrie return the construction of trie
func NewTrie() *Trie {
	return &Trie{
		next:   make(map[rune]*Trie, DefaultCap),
		isWord: false,
	}
}

// Insert insert a word into trie
func (t *Trie) Insert(word string) {
	root := t
	for _, v := range word {
		_, ok := root.next[v]
		if !ok {
			root.next[v] = NewTrie()
		}
		root = root.next[v]
	}

	root.isWord = true
}

// Search return is the word in the trie
func (t *Trie) Search(word string) bool {
	root := t
	for _, v := range word {
		if root.next[v] == nil {
			return false
		}
		root = root.next[v]
	}
	return root.isWord
}
