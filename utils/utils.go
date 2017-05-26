// Package utils provides general purpose functions of a telegram bot
package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Luxurioust/excelize"
)

// error check function
func Check(e error) {
	if e != nil {
		fmt.Println("Problem found")
	}
}

// Reads plain text documents and returns an string of them
func Doc_reader(fname string) string {
	file, err := os.Open(fname)
	Check(err)

	content, err := ioutil.ReadAll(file)
	Check(err)

	return string(content)
}

// working with .xlsx files
func Excel2str(fname string, fdir string) string {

	raw_string := ""
	full_dir, err := filepath.Abs(fdir)
	full_name := full_dir + "/" + fname
	Check(err)
	fmt.Printf("full file address = %+v\n", full_name)

	xlsx, err := excelize.OpenFile(full_name)
	if err != nil {
		fmt.Println(err)
		raw_string = `Sorry, File was not there :(
			add a file or try again later!`
		return raw_string
	}

	// Get sheet index.
	index := xlsx.GetSheetIndex("Sheet1")

	// Get all the rows in a sheet.
	rows := xlsx.GetRows("sheet" + strconv.Itoa(index))

	for _, row := range rows {
		for _, colCell := range row {
			// fmt.Print(colCell, "\t")
			raw_string += colCell + "\t"
		}
		raw_string += "\n"
	}
	return raw_string
}

// Url2File url, fname and returns int64
func Url2File(url, fname string) int64 {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	outfile, _ := os.Create(fname)
	defer outfile.Close()

	n, _ := io.Copy(outfile, resp.Body)
	return n
}
