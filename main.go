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

func main() {
	fmt.Println(timeline)
}
