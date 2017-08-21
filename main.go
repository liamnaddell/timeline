//© Liam Naddell https://github.com/liamnaddell/timeline
//licensed under GPL-2.0
//more info in the LICENSE file at the root of the program
package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"time"
)

var id int

type timearray []*event

type event struct {
	Desc string
	Time time.Time
}

var timeline timearray

func main() {
	//var actions = []string{"Print the timeline", "exit", "quit", "leave", "add", "delete", "remove"}
	app := cli.NewApp()
	app.Name = "timeline"
	app.Usage = "Manage a timeline"
	var home = os.Getenv("HOME")
	app.Commands = []cli.Command{
		{
			Name:  "print",
			Usage: "print the timeline",
			Action: func(c *cli.Context) error {
				timeline = Load(home + "/.timeline")
				timeline.PrintTimeline()
				return nil
			},
		},
		{
			Name:  "add",
			Usage: "add to the timeline",
			Action: func(c *cli.Context) error {
				timeline = Load(home + "/.timeline")
				timeline.AddDate(c.Args()[0], totime(c.Args()[1]))
				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "remove from the timeline",
			Action: func(c *cli.Context) error {
				timeline = Load(home + "/.timeline")
				timeline.DeleteDate(toint(c.Args().First()))
				return nil
			},
		},
		{
			Name:  "create",
			Usage: "create a new timeline",
			Action: func(c *cli.Context) error {
				NewTimeline(home + "/.timeline")
				return nil
			},
		},
	}
	app.Run(os.Args)
	timeline.Encode(home + "/.timeline")
}

func toint(s string) int {
	i, err := strconv.Atoi(s)
	checkerr(err)
	return i
}

func totime(s string) time.Time {
	t, err := time.Parse("20060102", s)
	checkerr(err)
	return t
}
