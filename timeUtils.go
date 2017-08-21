package main

import "fmt"
import "sort"
import "encoding/gob"
import "io/ioutil"
import "bytes"

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

func (t *timearray) Encode(filepath string) {
	var loc bytes.Buffer
	enc := gob.NewEncoder(&loc)
	err := enc.Encode(*t)
	checkerr(err)
	ret := loc.Bytes()
	checkerr(ioutil.WriteFile(filepath, ret, 0755))
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func Load(filepath string) timearray {
	var loc bytes.Buffer
	var dec = gob.NewDecoder(&loc)
	ret, err := ioutil.ReadFile(filepath)
	checkerr(err)
	n, err := loc.Write(ret)
	fmt.Println(n)
	checkerr(err)
	var timeline timearray
	err = dec.Decode(&timeline)
	checkerr(err)
	return timeline
}

func NewTimeline(filepath string) {
	var loc bytes.Buffer
	var t = new(timearray)
	enc := gob.NewEncoder(&loc)
	err := enc.Encode(t)
	checkerr(err)
	ret := loc.Bytes()
	checkerr(ioutil.WriteFile(filepath, ret, 0755))
}
