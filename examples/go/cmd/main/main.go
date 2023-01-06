package main

import (
	"log"
	"os"

	_ "example.com/translate-example/pkg/translations"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	lng, err := language.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	p := message.NewPrinter(lng)

	p.Printf("Hello\n")
	p.Printf("World\n")

	p.Printf("User has been registered successfully\n")

	country := "Latvia"
	p.Printf("Congrats! You are in %s\n", country)

	location := struct{ city, country string }{
		city:    "Riga",
		country: "Latvia",
	}

	p.Printf("%[1]s is a city of %[2]s!\n",
		location.city,
		location.country)

	numberOfFiles := 5
	p.Printf("%d files remaining!\n", numberOfFiles)
}
