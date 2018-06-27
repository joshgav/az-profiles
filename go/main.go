package main

import (
	hybrid "github.com/joshgav/az-profiles/go/hybrid/cmd"
	latest "github.com/joshgav/az-profiles/go/latest/cmd"
)

func main() {
	hybrid.Execute()
	latest.Execute()
}
