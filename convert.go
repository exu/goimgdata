package goimgdata

import (
	 "image"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
	"os"
	"log"
	"encoding/base64"
	"io/ioutil"
)

func Convert(filename string) string {
	file,err := os.Open(filename)
	defer file.Close()

	if err == nil {
		_, t, _ := image.Decode(file)
		byteData, _ := ioutil.ReadFile(filename)
		base64encoded := base64.StdEncoding.EncodeToString(byteData)
		return "data:image/" + t + ";base64," + base64encoded
	} else {
		log.Fatal(err)
	}
	return ""
}
