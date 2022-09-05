# Go Say
[![Go Report Card](https://goreportcard.com/badge/github.com/samlitowitz/go-say)](https://goreportcard.com/report/github.com/samlitowitz/go-say)

`go-say` is a `cowsay` clone in Go. Based on Tony Monroe's source found [here](https://github.com/moxiegirl/cowsay).

# Installation
`go get -u github.com/samlitowitz/go-say/cmd/go-say`

# Usage
```shell
 cat /tmp/test/expected/input-2.txt | go-say

# ___________
#< 123 1 456 >
# -----------
#    \
#     \
#      \
#                    ##        .
#              ## ## ##       ==
#           ## ## ## ##      ===
#       /""""""""""""""""___/ ===
#  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
#       \______ o          __/
#        \    \        __/
#          \____\______/
```

## Options
```shell
go-say -h

#Usage of go-say:
#  -T string
#        Tongue (default "  ")
#  -W int
#        Wrap width (default 40)
#  -b    Borg
#  -d    Dead
#  -e string
#        Eyes (default "oo")
#  -f string
#        Cow file (default "default.cow")
#  -g    Greedy
#  -if string
#        Input file instead of stdin
#  -l    List
#  -n    Strip tabs and new lines
#  -p    Paranoid
#  -s    Stoned
#  -t    Tired
#  -w    Wired
#  -y    Young
```
