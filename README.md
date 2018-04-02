# gclone

Clone repos directly to where they belong in your Go workspace from anywhere.

## Installation

```bash
go get https://github.com/ijsnow/gclone.git
```

## Usage

```bash
gclone https://github.com/gorilla/mux.git
# Will be cloned to $GOPATH/src/github.com/gorilla/mux
```

## Caveats

### Only one `$GOPATH`

This tool uses `$GOPATH` to decide where to clone. I'm aware that `$GOPATH` works
like the `$PATH` variable and you're able to specify multiple `$GOPATH`'s, but I
only have one which is to my Go workspace. If you want to support multiple `$GOPATH`s,
feel free to open a PR.
