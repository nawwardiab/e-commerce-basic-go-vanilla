package middleware

import (
	"fmt"
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
  fmt.Println(absStaticDir)
    
  // Serve /static/* from that dir
  fs := http.FileServer(http.Dir(absStaticDir))
  staticDir := http.StripPrefix("/static/", fs)
  fmt.Println("staticDir: ", staticDir) 
  fmt.Println(fs)
  return staticDir
} 