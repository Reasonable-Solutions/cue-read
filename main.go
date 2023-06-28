package main

import (
	"fmt"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func main() {
	c := cuecontext.New()
	entrypoints := []string{"hello.cue"}

	bis := load.Instances(entrypoints, nil)

	for _, bi := range bis {
		if bi.Err != nil {
			fmt.Println("an error during load: ", bi.Err)
			continue
		}

		value := c.BuildInstance(bi)

		if value.Err() != nil {
			fmt.Println("error during build?!", value.Err())
		}

		err := value.Validate()
		if err != nil {
			fmt.Println("validation error ", err)

		}

		fmt.Println("value: ", value)
	}
}
