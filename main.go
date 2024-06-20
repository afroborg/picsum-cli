package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var BASE_URL = "https://picsum.photos"

func main() {
	flags := NewFlags()

	client := NewClient(BASE_URL, flags)

	for i := 0; i < flags.Count; i++ {
		image := client.GetImage()
		defer image.Close()

		var filename string

		if filename = flags.Output; flags.Count > 1 {
			filename = fmt.Sprintf("%s-%d", flags.Output, i+1)
		}

		filename = strings.Split(filename, ".")[0]

		var extension string

		if extension = "jpg"; flags.Webp {
			extension = "webp"
		}

		if flags.Directory != "" {
			filename = flags.Directory + "/" + filename
			if _, err := os.Stat(flags.Directory); os.IsNotExist(err) {
				os.Mkdir(flags.Directory, 0755)
			}
		}

		file, err := os.Create(filename + "." + extension)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		_, err = io.Copy(file, image)

		if err != nil {
			panic(err)
		}
	}

}
