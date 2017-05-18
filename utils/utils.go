// Package utils provides general purpose functions of a telegram bot
package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// error check function
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Doc_reader(fname string) string {
	file, err := os.Open(fname)
	Check(err)

	content, err := ioutil.ReadAll(file)
	Check(err)

	return string(content)
}

// url2File url, fname and returns int64
func url2File(url, fname string) int64 {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	outfile, _ := os.Create(fname)
	defer outfile.Close()

	n, _ := io.Copy(outfile, resp.Body)
	return n
}
