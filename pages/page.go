package pages

import (
	"github.com/grellyd/filelogging/globallogger"
	"io/ioutil"
	"fmt"
	"os"
)

// Page on the disk, in public/ generated by hugo
type Page struct {
	Title string
	Body  []byte
	File  *os.File
	Ending PageEnding
}

// Save a page
func (p *Page) Save() error {
	filename := fmt.Sprintf("pages/%s.%s", p.Title, p.Ending) 
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load a page
func Load(section string, title string, pending PageEnding) (*Page, error) {
	filename := fmt.Sprintf("public/%s/%s.%s", section, title, pending)
	// TODO: Combine the two accesses
	file, err := os.OpenFile(filename, os.O_RDONLY, 444)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to load the page: %s", err)
	}
	return &Page{Title: title, Body: body, File: file}, nil
}

// CheckExistence of a page. Error on failure, nil otherwise
func CheckExistence(section string, title string, pending PageEnding) error {
	filename := ""
	if len(section) == 0 {
		filename = fmt.Sprintf("public/%s.%s", title, pending)
	} else {
		filename = fmt.Sprintf("public/%s/%s.%s", section, title, pending)
	}
	globallogger.Debug(fmt.Sprintf("looking for '%s'\n", filename))
	// Is the file missing?
	if fileNotExists(filename) {
		return fmt.Errorf("'%s.%s' of '%s' does not exist", title, pending, section)
	} 
	return nil
}

func fileNotExists(filename string) (exists bool) {
	exists = false
	f, err := os.Stat(filename)
	globallogger.Debug(fmt.Sprintf("File status: '%v'", f))
	globallogger.Debug(fmt.Sprintf("File IsNotExist: '%v'", os.IsNotExist(err)))
	return os.IsNotExist(err)
}
