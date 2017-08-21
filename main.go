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
var Version string
var path string
var home = os.Getenv("HOME")

func main() {
	app := cli.NewApp()
	app.Name = "timeline"
	app.Usage = "Manage a timeline"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:  "print",
			Usage: "print the timeline",
			Action: func(c *cli.Context) error {
				var f = c.Int("t")
				fmt.Printf("using timeline %d\n", f)
				setpath(f)
				timeline = Load(path)
				timeline.PrintTimeline()
				return nil
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "t",
					Usage: "specify which timeline is being used here, n can be zero, one, two, thee, ect. Timelines are stored in ~/.timeline/",
				},
			},
		},
		{
			Name:  "add",
			Usage: "add to the timeline",
			Action: func(c *cli.Context) error {
				var f = c.Int("t")
				fmt.Printf("using timeline %d\n", f)
				setpath(f)
				timeline = Load(path)
				timeline.AddDate(c.Args()[0], totime(c.Args()[1]))
				timeline.Encode(path)
				return nil
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "t",
					Usage: "specify which timeline is being used here, n can be zero, one, two, thee, ect. Timelines are stored in ~/.timeline/",
				},
			},
		},
		{
			Name:  "remove",
			Usage: "remove from the timeline",
			Action: func(c *cli.Context) error {
				var f = c.Int("t")
				fmt.Printf("using timeline %d\n", f)
				setpath(f)
				timeline = Load(path)
				timeline.DeleteDate(toint(c.Args().First()))
				timeline.Encode(path)
				return nil
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "t",
					Usage: "specify which timeline is being used here, n can be zero, one, two, thee, ect. Timelines are stored in ~/.timeline/",
				},
			},
		},
		{
			Name:  "create",
			Usage: "create a new timeline",
			Action: func(c *cli.Context) error {
				var f = c.Int("t")
				fmt.Printf("using timeline %d\n", f)
				setpath(f)
				NewTimeline(path)
				timeline.Encode(path)
				return nil
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "t",
					Usage: "specify which timeline is being used here, n can be zero, one, two, thee, ect. Timelines are stored in ~/.timeline/",
				},
			},
		},
	}
	app.Run(os.Args)
}

func setpath(timeline int) {
	for i := 0; i != timeline+1; i++ {
		if i == timeline {
			path = fmt.Sprintf("%s/.timeline/timeline%d", home, timeline)
		}
	}
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
