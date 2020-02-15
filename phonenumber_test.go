package phonenumber

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"testing"
)

type Test struct {
	Input string `json:"input"`
	Output []string `json:"output"`
}

func TestLetterCombinations (testUtil *testing.T) {
	f, err := os.Open("./tests.json")

	if err != nil {
		testUtil.Error(err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)

	var tests map[string]Test

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				testUtil.Run(name, func (tstUtil *testing.T) {
					output := LetterCombinations(test.Input)
					sort.Strings(output)
					sort.Strings(test.Output)
					str := fmt.Sprintf("%v", output)
					if str != fmt.Sprintf("%v",test.Output) {
						tstUtil.Errorf(str)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			testUtil.Error(err)
		}
	}
}

