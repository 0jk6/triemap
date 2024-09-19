package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	cache "github.com/0jk6/triemap/internal/cache"
)

type CacheRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var cacheInstance = cache.NewCache()

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"msg" : "available endpoints: POST /store, GET /get"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonEncoder := json.NewEncoder(w)

	err := jsonEncoder.Encode(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func GetHandler(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.URL.Path, "/")
	if len(splits) != 3 || splits[1] != "get" {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	key := splits[2]

	value := cacheInstance.Get(key)

	resp := map[string]string{"key": key, "value": value}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonEncoder := json.NewEncoder(w)

	err := jsonEncoder.Encode(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	var req CacheRequest
	
	// Decode the request body into req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	cacheInstance.Put(req.Key, req.Value)	
	
	resp := map[string]string{"msg" : "success"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonEncoder := json.NewEncoder(w)
	err = jsonEncoder.Encode(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func PrefixSearchHandler(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.URL.Path, "/")
	if len(splits) != 3 || splits[1] != "prefix" {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	prefix := splits[2]

	value := cacheInstance.PrefixSearch(prefix)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonEncoder := json.NewEncoder(w)

	err := jsonEncoder.Encode(value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func SuffixSearchHandler(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.URL.Path, "/")
	if len(splits) != 3 || splits[1] != "suffix" {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	suffix := splits[2]

	value := cacheInstance.SuffixSearch(suffix)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonEncoder := json.NewEncoder(w)

	err := jsonEncoder.Encode(value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}