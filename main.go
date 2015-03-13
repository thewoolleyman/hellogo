package main

import "fmt"

type SuffixService struct {
  suffixer func(string) string
}

func NewSuffixService(a func(string) string) *SuffixService {
  return &SuffixService{
    suffixer: a,
  }
}

func (this *SuffixService) Suffixize(str string) string {
  return this.suffixer(str)
}

func main() {
  ss := NewSuffixService(func(str string) string {
    return str + "!"
  })
  fmt.Println(ss.Suffixize("Hello, world")) // Hello, world!
}
