package main 

import (
  "sync"
)

//Constraints (type sets) not the normal interface
type StringOrInt interface {
  string | int
}
type safeCounter[T StringOrInt] struct {
  counts map[T] int  //as maps are not thread-safe
  mu sync.Mutex
}

func (sc *safeCounter[T]) addToMap(key T, n int) {
  sc.mu.Lock()
  defer sc.mu.Unlock()
  sc.counts[key] = n
}

//Also called Type Parameters
func main() {
  sc := safeCounter[string ]{
  counts: make(map[string] int),
  }

  sc2:= safeCounter[int] {
    counts: make(map[int] int),
  }
  sc.addToMap("Adi", 10)
  sc2.addToMap(12, 32)
}
