package main

import "fmt"
import "sort"
import "encoding/gob"
import "io/ioutil"
import "bytes"
import "time"

func (t *timearray) PrintTimeline() {
	timeline.Sort()
	fmt.Println("↥")
	var i = 1
	for _, value := range *t {
		y, m, d := value.Time.Date()
		fmt.Printf("|⇒%d:: %d %s %d: %s\n", i, y, m, d, value.Desc)
		i++
	}
	fmt.Println("↧")
}

func (t *timearray) Sort() {
	sort.Slice(*t, func(i, j int) bool {
		return (*t)[i].Time.Before((*t)[j].Time)
	})
}

func (t *timearray) AddDate(desc string, time time.Time) {
	var NewEvent = event{desc, time}
	*t = append(*t, &NewEvent)
	id++
}

func (t *timearray) DeleteDate(i int) {
	t.Sort()
	i--
	*t = append((*t)[:i], (*t)[i+1:]...)
}

func (t *timearray) Encode(filepath string) error {
	var loc bytes.Buffer
	enc := gob.NewEncoder(&loc)
	err := enc.Encode(*t)
	if err != nil {
		return err
	}
	ret := loc.Bytes()
	err = ioutil.WriteFile(filepath, ret, 0755)
	return err
}

func Load(filepath string) (timearray, error) {
	var loc bytes.Buffer
	var dec = gob.NewDecoder(&loc)
	ret, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	_, err = loc.Write(ret)
	if err != nil {
		return nil, err
	}
	var timeline timearray
	err = dec.Decode(&timeline)
	if err != nil {
		return nil, err
	}
	return timeline, nil
}

func NewTimeline(filepath string) error {
	var loc bytes.Buffer
	var t = new(timearray)
	enc := gob.NewEncoder(&loc)
	err := enc.Encode(t)
	if err != nil {
		return err
	}
	ret := loc.Bytes()
	err = ioutil.WriteFile(filepath, ret, 0755)
	if err != nil {
		return err
	}
	fmt.Printf("created new timeline at \"%s\"\n", filepath)
	return nil
}
