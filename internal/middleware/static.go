package middleware

import (
	"log"
	"net/http"
	"path/filepath"
)

//ServeStatic – prepares static directory and serves it to http.Handler
func ServeStatic(path string) http.Handler{
  // Abstract static folder path
	absStaticDir, err := filepath.Abs(path)
  if err != nil {
    log.Fatalf("invalid static directory path %q: %v", path, err)
  }
    
  // Serve /static/* from that dir
  fs := http.FileServer(http.Dir(absStaticDir))
  staticDir := http.StripPrefix("/staticFiles/", fs)
  return staticDir
} 