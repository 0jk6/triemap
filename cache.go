package cache

import (
	"sync"

	trie "github.com/0jk6/triemap/internal/trie"
	utils "github.com/0jk6/triemap/internal/utils"
)

type Cache struct {
	mp         sync.Map
	prefixTree *trie.Trie
	suffixTree *trie.Trie
}

func NewCache() *Cache {
	return &Cache{
		prefixTree: trie.NewTrie(),
		suffixTree: trie.NewTrie(),
	}
}

func (cache *Cache) Put(key, value string) {
	//store it in the cache
	cache.mp.Store(key, value)

	//store it in the prefix tree
	cache.prefixTree.Insert(key)
	cache.suffixTree.Insert(utils.Reverse(key))

}

func (cache *Cache) Get(key string) string {
	if value, ok := cache.mp.Load(key); ok {
		return value.(string)
	}
	return ""
}

func (cache *Cache) PrefixSearch(prefix string) map[string]string {
	foundKeys := cache.prefixTree.FindWordsWithPrefix(prefix)

	results := make(map[string]string)

	for _, key := range foundKeys {
		results[key] = cache.Get(key)
	}

	return results
}

func (cache *Cache) SuffixSearch(suffix string) map[string]string {
	foundKeys := cache.suffixTree.FindWordsWithSuffix(suffix)

	results := make(map[string]string)

	for _, key := range foundKeys {
		results[key] = cache.Get(key)
	}

	return results
}
