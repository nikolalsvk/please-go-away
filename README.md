# please-go-away

Make art out of your image by running this Golang file

## Usage

### Install dependencies

```
$ go get -u github.com/disintegration/gift
```

### Set GOPATH variable

First set your $GOPATH in `~/.bashrc` or `~/.zshrc`:

```
export GOPATH="$HOME/where/ever/you/want"
```
### Clone this file in your $GOPATH

Then clone this repo there:

```
$ cd $GOPATH
$ git clone git@github.com:nikolalsvk/please-go-away.git
```

### Run it

```
$ go build
$ ./please-go-away path-to-your-pic.jpeg new-pic.jpeg
```

That's all folks!
