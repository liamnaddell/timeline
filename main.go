//© Liam Naddell https://github.com/liamnaddell/timeline
//licensed under GPL-2.0
//more info in the LICENSE file at the root of the program
package main

import "fmt"

const timeline = `
//examples:↳↱↑↓

  1863: Death of William Thackeray
  ↑
<-------------------->
 ↓
 1811: Birth of William Thackeray
`

type item struct {
	//when this happened
	Date int
	//the description
	Desc string
	//whether the item is raised. Ie on the second or first line above/below the timeline
	Raised bool
}

func main() {
	fmt.Println(timeline)
	var testItem = item{2007, "Random date", true}
	fmt.Println(testItem.Date, testItem.Desc, testItem.Raised)
}
