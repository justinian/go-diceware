# go-diceware

This is a little tool for generating [diceware][]-like passwords on the command
line. PLEASE NOTE, this is NOT AS SECURE as generating a real diceware password
with real dice. Please do not use this generator for anything other than
TEMPORARY passwords - see the actual [diceware][] page for generating
real-world passwords.

[diceware]: http://world.std.com/~reinhold/diceware.html

## Installation

Install [go][], then run:

```
go get github.com/justinian/go-diceware
```

This should install the `go-diceware` binary in your `$GOPATH/bin`.

[go]: http://golang.org/

## Usage

```
Usage of ./go-diceware:
  -count=10: Number of phrases to generate
   -words=5: Number of words in a phrase
```
