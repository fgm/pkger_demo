# fgm/pkger-demo

A short demo using https://github.com/markbates/pkger to load templates

## Features

This demo shows how to load multiple templates from pkger using the pkger
functions instead of traditional `io`|`ioutil` file access.

## Requirements

- Go 1.14 or later
- markbates/pkger 0.15.0 or later

## Building
### For development

During development, the `pkged.go` file is absent. 

- If you generated it previously, just remove it.
- Compile and run normally, e.g. with `go build -o dev .` The first build will download
  dependency modules before compiling.
- Run the program, e.g. `./dev`, and visit http://localhost:8080/some-path  
- In that configuration, the templates are loaded from the module directory: you
  may change them and restart the program without recompiling, and the changes
  will be applied.  
- You can use `pgker list` to list when `pkger` will bundle during build. It should
  display results like:

```
github.com/fgm/pkger_demo
> github.com/fgm/pkger_demo:/templates
> github.com/fgm/pkger_demo:/templates/layout
> github.com/fgm/pkger_demo:/templates/layout/footer.gohtml
> github.com/fgm/pkger_demo:/templates/page.gohtml
```  
  
### For deployment

For deployment, the `pkged.go` file needs to be generated.

Since the `main.go` contains a `go generate` build directive, the steps are:

- `go generate`: this will create the `pkged.go` file
- `go build -o run`: compile the file, bundling the templates in the binary.
- `ls -l dev run`: show the size difference in the binaries, caused by the bundled
  templates.
- You can now modify or delete the templates without changing the program output. 

*A word of caution*: running `pkger list` removes the generated `pkged.go`
file, so the next deployment build will need a `go generate` again.

## Going further

- [Pkger documentation on go.dev](https://pkg.go.dev/github.com/markbates/pkger?tab=overview)

## License

MIT.


    
