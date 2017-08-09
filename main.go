//© Liam Naddell https://github.com/liamnaddell/timeline
//licensed under GPL-2.0
//more info in the LICENSE file at the root of the program
package main

import "fmt"

const tl = `
//examples:↳↱↑↓

  1863: Death of William Thackeray
  ↑
<-------------------->
 ↓
 1811: Birth of William Thackeray
`

type date struct {
	//when this happened
	Date int `json:"date"`
	//the description
	Desc string `json:"desc"`
}

type timeline struct {
	Dates []date `json:"dates"`
}

var ttl timeline

func init() {
	var testItem = date{2007, "Random date"}
	var testItem2 = date{1234, "party time lol"}
	ttl.Dates = append(ttl.Dates, testItem)
	ttl.Dates = append(ttl.Dates, testItem2)
}

func main() {
	fmt.Println(ttl.Dates)
}
