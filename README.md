# go-embed-version

This is a quick tutorial on embedding build-time information, such as version, into a Go binary. I'm using `make` in this case but it's not required. I'm also not very adept with `make` so take that into consideration.

# -ldflags -X

`go build` will accept a flag `-ldflags` to pass specific arguments to the linker. The `link` tool will accept a `-X` flag to set the value of an exported variable. It will only set variables, not constants. The full import path to the variable must be specified.

# A Version variable

In my projects I tend to create a `<package>/cmd/version.go` file. Each binary can import this package. The location for the variable can be in any package you like.

`cmd/version.go`:

```
package cmd

import (
	"flag"
	"fmt"
	"os"

       )

var (
	Version = "" // set at compile time with -ldflags "-X versserv/cmd.Version=x.y.yz"

	FVersion = flag.Bool("version", false, "show version and exit")

    )

func ShowVersion() {
	fmt.Println("version:", Version)
	os.Exit(0)

}
```

I also provide a flag and utility function to show the version and exit. If the version isn't set ont he command line during `go build` the variable is unchanged:

```
cmd$ go build -o /tmp/versserv server/server.go 
cmd$ /tmp/versserv -version
version: 
```

You can set it to an arbitrary value by hand:

```
cmd$ go build -o /tmp/versserv -ldflags "-X github.com/AgentZombie/go-embed-version/cmd.Version=foo" server/server.go 
cmd$ /tmp/versserv -version
version: foo
```

# Setting the variable automatically

I'm using `make` and I have it create the version string based on `git` tag or a timestamp, depending on if it's a dev or release binary.

`Makefile`:

```
VERSION := $(shell git tag | grep ^v | sort -V | tail -n 1)
LDFLAGS = -ldflags "-X github.com/AgentZombie/go-embed-version/cmd.Version=${VERSION}"
TIMESTAMP := $(shell date +%Y%m%d-%H%M)
DEVLDFLAGS = -ldflags "-X github.com/AgentZombie/go-embed-version/cmd.Version=dev-${TIMESTAMP}"

SERVBIN=/tmp/versserv

dev-serv: server/*
        go build ${DEVLDFLAGS} -o ${SERVBIN} server/*.go

serv: server/*
        go build ${LDFLAGS} -o ${SERVBIN} server/*.go
```

The `VERSION` pipeline in the `Makefile` lists all tags visible from this commit, selects those that being with `v` assuming those are semver tags. It sorts those using `sort`'s version sorting, then keeps only the latest one.'

```
cmd$ make dev-serv
go build -ldflags "-X github.com/AgentZombie/go-embed-version/cmd.Version=dev-20190608-1945" -o /tmp/versserv server/*.go
cmd$ /tmp/versserv -version
version: dev-20190608-1945
cmd$ rm /tmp/versserv 
cmd$ make serv
go build -ldflags "-X github.com/AgentZombie/go-embed-version/cmd.Version=v0.2.0" -o /tmp/versserv server/*.go
cmd$ /tmp/versserv -version
version: v0.2.0
```

# Tagging a version

In `git` tags are arbitrary labels but it's a common convention to use tags with [Semantic Versioning](https://semver.org/). Tags apply to _commits_ so any changes you want a tag to apply to must be committed. A simple tag operation looks like this:

```
git tag -a v0.3.0
```

You can also apply a message with the tag and it's good to use this to communicate what's captured in this version:

```
git tag -a v0.3.0 -m 'Add tutorial in README.md'
```

By default, tags aren't pushed upstream with other commits. To push with tags add `--follow-tags`:

```
git push --follow-tags
```

You can also configure `git` to always push tags with commits but it's not the best idea.

# Additional use for the version

I also build `docker` images using `make` and include the version tag in the build:

```
docker-image:
	docker build -t versserv:${VERSION} dockerdir
```

# Special Thanks

Special thanks to [subcon](https://github.com/subcon42) for gently nudging me to publish this.
