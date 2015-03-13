package main
import "fmt"

func MakeSuffixizer(s func(string)string) func(string)string {
 return func(v string) string {
   return s(v)
 }
}

func main() {
  ss2 := MakeSuffixizer(func(str string) string { return str + "!!" })
   fmt.Println(ss2("Hello, world"))
}