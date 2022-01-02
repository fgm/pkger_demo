# fgm/pkger-demo V2

A short demo using Go `embed` to load templates

## Features

This demo shows how to load multiple templates from `embed` using the `embed.FS`
methods instead of traditional `io`|`ioutil` file access, or the legacy
`markbates/pkger` contributed package.

## Requirements

- Go 1.16 or later

## Building
 
- `go build .`: compile the file, bundling the templates in the binary.
- You can now modify or delete the templates without changing the program output. 

## Going further

- [embed documentation on go.dev](https://pkg.go.dev/embed)

## License

MIT.
