# timeline

A program for managing a timeline.

## Installing

```
go get -d github.com/liamnaddell/timeline
cd $GOPATH/src/github.com/liamnaddell/timeline
make
sudo make install
(installs it to /usr/local/bin)
```

## How to use

```
NAME:
   timeline - Manage a timeline

USAGE:
   timeline [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     print    print the timeline
     add      add to the timeline
     remove   remove from the timeline
     create   create a new timeline
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

The timeline is stored in `~/.timeline`, and right now, you can only have one timeline.

## Contributing

There will probably not be many communications on this repo, so any way of contacting issues/improvements is greatly appreciated

### Code of conduct

* Please don't be mean :)
* you are not talking to brick walls, we all have feelings 
