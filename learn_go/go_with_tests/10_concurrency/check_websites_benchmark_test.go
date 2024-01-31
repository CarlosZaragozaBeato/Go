package concurrency 

import (
  "testing"
  "time"
)

func slowStubWebsiteChecker (_ string) bool {
  time.Sleep(20 * time.Millisecond)
  return true
}

func BenchmarkCheckWebsites(b *testing.B) {
  urls := make 
}
