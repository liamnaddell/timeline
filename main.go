//© Liam Naddell https://github.com/liamnaddell/timeline
//licensed under GPL-2.0
//more info in the LICENSE file at the root of the program
package main

import "fmt"

const tl = `
↥
|⇒1811: Birth of William Thackeray
|⇒1863: Death of William Thackeray
↧

`

var id int

type timearray []*event

type event struct {
	Desc string
	Time int
}

var timeline timearray

func init() {
	timeline.AddDate("Death of William Thackeray", 1863)
	timeline.AddDate("Birth of William Thackeray", 1811)
	timeline.AddDate("Death of death", 1653)
	timeline.AddDate("turn of the year of 1920", 1921)
	timeline.AddDate("turn of the year of 1821", 1821)
	//timeline.DeleteDate(1)
}

func main() {
	fmt.Println(tl)
	timeline.PrintTimeline()

}
