// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/stackworx-go/entext"
)

func main() {
	opts := []entc.Option{
		entc.Extensions(entext.Extension{}),
	}
	err := entc.Generate("./schema", &gen.Config{}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
