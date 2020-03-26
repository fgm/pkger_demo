//go:generate pkger

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/markbates/pkger"
)

/* PageData is an example template data structure */
type PageData struct {
	Path string
	Year int
}

func compileTemplates(dir string) (*template.Template, error) {
	const fun = "compileTemplates"
	tpl := template.New("")
	// Since Walk receives a dynamic value, pkger won't be able to find the
	// actual directory to package from the next line, which is why we used
	// pkger.Include() in main().
	err := pkger.Walk(dir, func(path string, info os.FileInfo, _ error) error {
		// Skip non-templates.
		if info.IsDir() || !strings.HasSuffix(path, ".gohtml") {
			return nil
		}
		// Load file from pkpger virtual file, or real file if pkged.go has not
		// yet been generated, during development.
		f, _ := pkger.Open(path)
		// Now read it.
		sl, _ := ioutil.ReadAll(f)
		// It can now be parsed as a string.
		tpl.Parse(string(sl))
		return nil
	})
	return tpl, err
}

func main() {
	const addr = ":8080"

	// Tell pkger that it has to package that directory.
	dir := pkger.Include("/templates")
	// Only compile templates on startup.
	tpl, _ := compileTemplates(dir)

	// Now serve pages
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := PageData{Path: r.URL.Path, Year: time.Now().Local().Year()}
		tpl.ExecuteTemplate(w, "page", ctx)
	})
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
