package concurrency 

// WebSiteChecker checks a url, returning a bool 
type WebSiteChecker func(string) bool 
type result struct {
  string 
  bool
}

// CheckWebsites takes a WebSiteChecker and a slice 
// of urls and returns a map of urls to the result of 
// checking each url with WebsiteChecker function 
func CheckWebsites(wb WebSiteChecker, urls []string) map[string]bool {
  results := make(map[string]bool

  )
  resultChannel := make(chan result)

  for _, url := range urls {
    go func (u string){
      resultChannel <- result (u, wc(u))
    }(url)
  }

  for i := 0; i < len(urls); i++ {
    r := <- resultChannel
    results[r.string] = r.bool 
  }

  return results
}

