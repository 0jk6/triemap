package trie

import (
	utils "github.com/0jk6/triemap/internal/utils"
)

//define the Node of a trie
type Node struct {
	children [26]*Node
	end      bool
}

func (node *Node) ContainsLetter(ch rune) bool {
	return node.children[ch-'a'] != nil
}

func (node *Node) Put(ch rune, childNode *Node) {
	node.children[ch-'a'] = childNode
}

func (node *Node) Get(ch rune) *Node {
	return node.children[ch-'a']
}

func (node *Node) SetEnd() {
	node.end = true
}

func (node *Node) GetEnd() bool {
	return node.end
}

//define the trie strcut
type Trie struct {
	root *Node
}

//constructor
func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (trie *Trie) Insert(word string) {
	node := trie.root

	for _, ch := range word {
		//insert only if the node doesn't have a child ch
		if node.ContainsLetter(ch) == false {
			node.Put(ch, &Node{})
		}

		//move deeper
		node = node.Get(ch)
	}

	//mark the end
	node.SetEnd()
}

func (trie *Trie) Search(word string) bool {
	node := trie.root

	for _, ch := range word {
		if node.ContainsLetter(ch) == false {
			return false
		}

		node = node.Get(ch)
	}

	return node.GetEnd()
}

//prefix search implementation
func (trie *Trie) FindWordsWithPrefix(prefix string) []string {
	//move to the last letter of prefix in the trie
	node := trie.root

	for _, ch := range prefix {
		if node.ContainsLetter(ch) == false {
			return []string{}
		}

		node = node.Get(ch)
	}

	//now the node will be pointing to the children where prefix's last letter is present
	//from here on, we can do dfs through current node's children, and combine them with prefix
	//and we'll get the required answers
	var result []string
	var temp string
	trie.dfs(node, prefix, temp, &result, false)

	return result
}

func (trie *Trie) FindWordsWithSuffix(suffix string) []string {
	node := trie.root

	//since we are searching in suffix tree, we will reverse the string
	suffix = utils.Reverse(suffix)

	for _, ch := range suffix {
		if node.ContainsLetter(ch) == false {
			return []string{}
		}

		node = node.Get(ch)
	}

	var result []string
	var temp string
	trie.dfs(node, suffix, temp, &result, true) //set reverse to true, to reverse all the found strings during dfs

	return result
}

//dfs implementation
func (trie *Trie) dfs(node *Node, prefix string, temp string, result *[]string, reverseString bool) {
	//base conditions for stopping the recursion
	if node.GetEnd() == true {
		//we've found one of the possible word starting with prefix, append it to the result array
		found := prefix + temp
		if reverseString == true {
			found = utils.Reverse(found)
		}
		*result = append(*result, found)
	}

	//loop through all the children and do dfs
	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			//move in
			ch := rune(i + 'a')
			trie.dfs(node.children[i], prefix, temp+string(ch), result, reverseString)
		}
	}
}
