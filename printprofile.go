package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/pprof/profile"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p, err := profile.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	total := int64(0)
	for _, s := range p.Sample {
		// Print function call stack. Printing in reverse to appear like a stack when printed.
		for i := len(s.Location) - 1; i >= 0; i-- {
			for j := len(s.Location[i].Line) - 1; j >= 0; j-- {
				fmt.Println(s.Location[i].Line[j].Function.Name)
			}
		}

		// Determine and print default value from value index. Pprof specifies
		// that the last samples in the sample index is the default when there
		// is no other information.
		defaultSampleIndex := len(s.Value) - 1
		value := s.Value[defaultSampleIndex]
		fmt.Println(value)

		total += value

		if len(s.Label) > 0 {
			fmt.Println(s.Label)
		}
		fmt.Println("---")
	}

	fmt.Println("Total: ", time.Duration(total))
}
