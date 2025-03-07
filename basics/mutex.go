package main 

import (
  "sync"
)

type safeCounter struct {
  counts map[string] int  //as maps are not thread-safe
  mu sync.Mutex
}

func (sc *safeCounter) addToMap(s string, n int) {
  sc.mu.Lock()
  defer sc.mu.Unlock()
  sc.counts[s] = n;
}

func main() {

  //func protect() {
  //mux.Lock()
  //defer mux.Unlock()
  //}
 //use in case of maps as they are not shread safe
 s := safeCounter{
   counts : make(map [string]int),
 }
  s.addToMap("Acx", 10) 
}
