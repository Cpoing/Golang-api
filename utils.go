package main

import (
  "net/http"
  "encoding/json"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  json.NewEncoder(w).Encode(v)
}
