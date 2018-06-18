package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/mail"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("fooda_email")

	if err != nil {
		panic("Couldn't read")
	}

	message, err := mail.ReadMessage(file)

	if err != nil {
		panic("Oh no")
	}

	fmt.Printf(message.Header.Get("From"))
	fmt.Printf(message.Header.Get("Content-Type"))
	mediaType, params, err := mime.ParseMediaType(message.Header.Get("Content-Type"))

	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(message.Body, params["boundary"])
		p, err := mr.NextPart()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		slurp, err := ioutil.ReadAll(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q\n", slurp)
		// fmt.Printf("Part %q: %q\n", p.Header.Get("Content-Type"), slurp)
	}

	// body, err := ioutil.ReadAll(message.Body)

	// if err != nil {
	// 	panic("Oh no 2")
	// }

	fmt.Println("")
	// fmt.Println(string(body))

}
