package parser

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
)

func ReadSpec(path string) {
	spec, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	document, err := libopenapi.NewDocument(spec)
	if err != nil {
		panic(err)
	}

	// We ultimately want to generate this as a v2 or v3 but stuck a bit
	fmt.Printf("Spec:\n\t%s\n", document)
}
