package util

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

// ToCP1251 ...
func ToCP1251(s string) string {
	r := strings.NewReader(s)
	tr := transform.NewReader(r, charmap.Windows1251.NewEncoder())
	buf, _ := ioutil.ReadAll(tr)
	return string(buf)
}

// FromCP1251 ...
func FromCP1251(s string) string {
	r := strings.NewReader(s)
	tr := transform.NewReader(r, charmap.Windows1251.NewDecoder())
	buf, _ := ioutil.ReadAll(tr)
	return string(buf)
}
