package main

import "fmt"
import "sort"

func (t *timearray) PrintTimeline() {
	timeline.Sort()
	fmt.Println("↥")
	var i = 1
	for _, value := range *t {
		fmt.Printf("|⇒%d:: %d: %s\n", i, value.Time, value.Desc)
		i++
	}
	fmt.Println("↧")
}

func (t *timearray) Sort() {
	sort.Slice(*t, func(i, j int) bool {
		return (*t)[i].Time < (*t)[j].Time
	})
}

func (t *timearray) AddDate(desc string, time int) {
	var NewEvent = event{desc, time}
	*t = append(*t, &NewEvent)
	id++
}

func (t *timearray) DeleteDate(i int) {
	t.Sort()
	i--
	*t = append((*t)[:i], (*t)[i+1:]...)
}
