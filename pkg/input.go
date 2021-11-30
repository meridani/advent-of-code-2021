package pkg

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cavaliercoder/grab"
)

// Input holds the puzzle inputs as a string
type Input string

// FromIntSlice creates an Input from a slice of ints
func (i *Input) FromIntSlice(numbers []int) {
	// fmt.Println(i)
	*i = Input(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), "\n"), "[]"))
	// fmt.Println(i)
}

// AsIntSlice returns the Input as a slice of ints
func (i *Input) AsIntSlice() []int {
	if len(*i) == 0 {
		return []int{}
	}
	numbers := i.AsStringSlice()
	ints := make([]int, len(numbers))
	for i, current := range numbers {
		num, err := MustAtoi(current)
		if err == nil {
			ints[i] = num
		}
	}
	return ints
}

// AsStringSlice returns the Input as a slice of strings
func (i *Input) AsStringSlice() []string {
	if len(*i) == 0 {
		return []string{}
	}
	return strings.Split(strings.Trim(string(*i), "\n"), "\n")
}

// ReadInput reads the input file into a string
func ReadInput(fileName string) (puzzle Input, err error) {
	content, err := os.ReadFile(fileName)
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
		session, err := os.ReadFile("../../assets/.aocsession")
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
