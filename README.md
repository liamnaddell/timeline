# timeline

A program for managing a timeline.

## Installing

This program requires `dep` to be installed in order to manage packages correctly
you can install it by running `go get -u github.com/golang/dep/cmd/dep` on linux, or `brew install dep`
```
go get -d github.com/liamnaddell/timeline
cd $GOPATH/src/github.com/liamnaddell/timeline
make
sudo make install
(installs it to /usr/local/bin)
```
## Uninstalling

`sudo make uninstall`

this will remove it from `/usr/local/bin`

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

There will probably not be very many communications on this repo, so any way of contacting issues/improvements is greatly appreciated

### Code of conduct

* Please don't be mean :)
* you are not talking to brick walls, we all have feelings 

## Work in progress

this repo is currently being worked on

### What needs to be done

* Better error handling
* Better docs while on command line(im lookin at u subcommands)
* Developer guide
* In line comments
