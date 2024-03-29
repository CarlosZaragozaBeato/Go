package main 


import (
  "testing"
)



func TestSearch(t *testing.T) {
  dictionary := dictionary{"test": "this is just a test"}
  
  t.Run("known word", func(t *testing.T){
    got, _ := dictionary.Search("test")
    want := "this is just a test"
    
    assertString(t, got, want)
  })

  t.Run("unknown word", func(t *testing.T){
    _, got := dictionary.Search("unknown")

    assertError(t, got, ErrNotFound)
  })
}



func TestAdd(t *testing.T){
  t.Run("new word", func(t *testing.T){
    dictionary := Dictionary{}
    word := "test"
    definition := "this is just a test"

    err := dictionary.Add(word, definition)

    assertError(t, err, nil)
    assertDefinition(t, dictionary, word, definition)
  })
}




func TestDelete(t *testing.T){
  word := "test"
  dictionary := Dictionary{word: "test definition"}
  dictionary.Delete(word)

  _, err := dictionary.Search(word)
  assertError(t, err, ErrNotFound)
}




func assertStrings(t testing.FB, got,want string) { 
  t.Helper()

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}




func assertError(t testing.TB, got, want string){
  t.Helper()

  if got != want {
    t.Errorf("got error %q want %q", got, want)
  }
}



fun assertDefinition(t testing.TB, dictionary Dictionary, word, definition string){
  t.Helper()

  got, err = dictionary.Search(word)

  if err != nil {
    t.Fatal("Should find added word, "err)
  }
  if definition != got {
    t.Errorf("got %q want %q", got, definition)
  }
}




































