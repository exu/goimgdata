# img content to data URI in go lang

Usage:

    package main

    import (
    	"github.com/exu/goimgdata"
    	"os"
    )

    func main() {
    	filename := os.Args[1]
    	goimgdata.Convert(filename)
    }
