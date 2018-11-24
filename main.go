package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func dotheThing(csvf, jsonf string) {
	f, err := os.Open(csvf)
	if err != nil {
		fmt.Println("Couldn't Open file")
		fmt.Println(err.Error())

	}
	defer f.Close()
	r := csv.NewReader(f)

	var ar = make([]map[string]interface{}, 0)
	names, _ := r.Read()
	for {
		l, e := r.Read()
		if e == io.EOF {
			break
		}
		m := make(map[string]interface{})
		for i, v := range names {
			m[v] = l[i]

		}
		ar = append(ar, m)
	}
	b, _ := json.Marshal(ar)
	if len(jsonf) == 0 {

		fmt.Print(string(b))

	} else {

		if err = ioutil.WriteFile(jsonf, b, 0644); err != nil {
			fmt.Println("Couldn't write output to file")
			fmt.Println(err.Error())
		}
	}

}

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		fmt.Println("Please specify a input file")
	case 1:
		dotheThing(args[0], "")
	case 2:
		csvf, jsonf := args[0], args[1]
		dotheThing(csvf, jsonf)
	default:
		csvf, jsonf := args[0], args[1]
		dotheThing(csvf, jsonf)

	}
}
