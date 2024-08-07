package main

import "fmt"

/*
*
前缀树 数组大法
*/
type Trie struct {
	arr    [26]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	curr := t
	for _, ch := range word {
		if curr.arr[ch-'a'] == nil {
			curr.arr[ch-'a'] = &Trie{}
		}
		curr = curr.arr[ch-'a']
	}
	curr.isWord = true
}

func (t *Trie) Search(word string) bool {
	curr := t
	for _, ch := range word {
		if curr.arr[ch-'a'] == nil {
			return false
		}
		curr = curr.arr[ch-'a']
	}
	return curr.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t
	for _, ch := range prefix {
		if curr.arr[ch-'a'] == nil {
			return false
		}
		curr = curr.arr[ch-'a']
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	fmt.Println(trie.StartsWith("jpp"))

}
