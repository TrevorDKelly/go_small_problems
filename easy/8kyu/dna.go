/* Deoxyribonucleic acid, DNA is the primary information storage molecule in biological systems. It is composed of four nucleic acid bases Guanine ('G'), Cytosine ('C'), Adenine ('A'), and Thymine ('T').

Ribonucleic acid, RNA, is the primary messenger molecule in cells. RNA differs slightly from DNA its chemical structure and contains no Thymine. In RNA Thymine is replaced by another nucleic acid Uracil ('U').

Create a function which translates a given DNA string into RNA.

For example:

"GCAT"  =>  "GCAU"
The input string can be of arbitrary length - in particular, it may be empty. All input is guaranteed to be valid, i.e. each input string will only ever consist of 'G', 'C', 'A' and/or 'T'. */

/*

  get string
  return string

  look thru letters of the given string. any that are 'T', change them to 'U'

  iterate thru string and add each letter to a new string. if its a t, add a u instead

  init new string
  iterate thru letters - range
    if rune is 'T'
      rune = 'u'
    add string rune to string
  return string
*/

package main

import "fmt"

func DNAtoRNA(dna string) (rna string) {
  for _, char := range dna {
    if char == 'T' {
      char = 'U'
    }
    rna += string(char)
  }
  return rna
}

func test(word, expected string) {
  result := DNAtoRNA(word)
  fmt.Println(result == expected)
}

func main() {
  test("GCAT", "GCAU")
  test("", "")
  test("G", "G")
  test("T", "U")
}
