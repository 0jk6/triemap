# TrieMap

TrieMap is a key-value store that supports searching keys by prefix or suffix.

### How It Works
TrieMap stores key-value pairs in a hash table, and it uses a Trie data structure to enable efficient prefix and suffix searches. 
This allows for fast retrieval of keys based on their starting or ending characters.


### Usage
Build the binary using `go build main.go` and run `./main`, the server will listen on port `8080`, make sure that you have go version >= 1.22

To store data in the TrieMap, use the following command:

```bash
curl -X POST http://localhost:8080/store \
-H "Content-Type: application/json" \
-d '{"key": "hello", "value": "there"}'
```

To retrieve the value associated with a specific key:

```bash
curl -X GET http://localhost:8080/get/hello
```

To retrieve all key-value pairs where the key starts with a specific prefix:

```bash
curl -X GET http://localhost:8080/prefix/he
```

To retrieve all key-value pairs where the key ends with a specific suffix:

```bash
curl -X GET http://localhost:8080/suffix/lo
```
