package main 

import "testing"

func AsserEqual[T comparable](t *testing.T, got, want T) {
  t.Helper()

  if got != want {
    t.Errorf("got %v, want %v", got, want)
  }
}


func AssertNotEqual[T comparable](t *testing.T, got want){
  t.Helper()

  if got == want {
    t.Errorf("Didn't want %v", got)
  }
}

func AssertTrue(t * testing.T, got bool){
  t.Helper()

  if !got {
    t.Errorf("got %v, want true",got)
  }
}

func AssertFalse(t *testing.T, got bool) {
  t.Helper()

  if got {
    t.Errorf("got %v, want false", got)
  }
}


