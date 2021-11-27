package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/cavaliercoder/grab"
)

var AOCSESSION = ""

type Input string

func (i *Input) FromIntSlice(numbers []int) Input {
	fmt.Println(i)
	*i = Input(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), "\n"), "[]"))
	fmt.Println(i)
	return *i
}

func (i *Input) AsIntSlice() []int {
	numbers := i.AsStringSlice()
	ints := make([]int, len(numbers))
	for i, current := range numbers {
		ints[i] = MustAtoi(current)
	}
	return ints
}

func (i Input) AsStringSlice() []string {
	return strings.Split(strings.Trim(string(i), "\n"), "\n")
}

// ReadInput reads the input file into a string
func ReadInput(fileName string) (puzzle Input, err error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	puzzle = Input(content)

	return puzzle, nil
}

// CheckAndDownloadFile checks if we have the input file for the day
// and downloads it if it is not available
func CheckAndDownloadFile(filename, url string) {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Print(os.Getwd())
		log.Printf("File not found: %v Downloading from %v\n", filename, url)
		client := grab.NewClient()
		req, err := grab.NewRequest(filename, url)

		if err != nil {
			log.Println("Can't read .aocsession file so no download for you!")
			return
		}
		session, err := ioutil.ReadFile("../../assets/.aocsession")
		if err != nil {
			log.Println(err)
		}
		req.HTTPRequest.Header.Set("cookie", fmt.Sprintf("session=%v", string(session)))
		resp := client.Do(req)
		if err := resp.Err(); err != nil {
			log.Println(err)
		}
	}
}
