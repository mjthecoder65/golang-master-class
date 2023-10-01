package dsa

// Trie is a tree-like data structure that stores a dynamic set of strings
// Is also called digital tree or prefix tree.
// It is a k-nary tree, where k is the size of the alphabet.
// and searching for strings in a trie is faster than searching for strings in a hash table or binary search tree.
// It is also know as a prefix tree because it can efficient find all strings with a common prefix.
// In a trie, each node represents a single character.
// The root node is typically empty, or represent empty string.

// Trie interface
type ITrie interface {
	Insert(word string)
	SearchPrefix(word string) bool
	Search(word string) bool
}

// TrieNode struct
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie struct
type Trie struct {
	root *TrieNode
}

// Insert a word into the trie.
func (t *Trie) Insert(word string) {
	node := t.root

	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// SearchPrefix returns true if the word is in the trie.
func (t *Trie) SearchPrefix(word string) bool {
	node := t.root

	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return true
}

// Search returns true if there is any word in the trie that starts with the given word.
func (t *Trie) Search(word string) bool {
	node := t.root

	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}
