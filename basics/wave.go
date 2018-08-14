package main

/**
 * How to create a mexican wave from letters with goland
 * https://media.giphy.com/media/kvYQSEwUPhtMCYS98k/giphy.gif
 * 
 * For example, if you have string "mary", you should output next lines:
 * Mary
 * mAry
 * maRy
 * marY
 */

import s "strings"
import "fmt"
import "unicode"

func isLetter(c string) bool {
  r := []rune(c)[0]
  b := unicode.IsLetter(r);

  return b;
}

func wave(phrase string) {
  // lowercase phrase and explode by letters
  lower := s.ToLower(phrase)
  letters := s.Split(lower, "")

  for index, letter := range letters {
    // skip non-letter characters
    if !isLetter(letter) {
       continue
    }

    // clone letters array
    tmp := make([]string, len(letters))
    copy(tmp, letters)

    // wave by letter and print result
    tmp[index] = s.ToUpper(letter)
    fmt.Println(s. Join(tmp, ""))
  }
}

func main() {
  wave("Mary had a little lamb, it's fleece was white as snow, yeah")
}

