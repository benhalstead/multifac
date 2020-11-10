package main

import "github.com/graniticio/granitic/v2"
import "multifac/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
