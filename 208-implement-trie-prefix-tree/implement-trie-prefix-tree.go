package leetcode

/*
https://leetcode.com/problems/implement-trie-prefix-tree/

208. Implement Trie (Prefix Tree)
Medium

1526

31

Favorite

Share
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.
 */

type Trie struct {
	content []string
}
//#TODO:refactor it
/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.content = append(this.content, word)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range this.content {
		if v == word {
			return true
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range this.content {
		if len(v) >= len(prefix) && v[0:len(prefix)] == prefix {
			return true
		}
	}
	return false
}

/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
*/
