package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

//X ...
type X struct {
	FirstName string
	LastName  string
	Age       int
}

//Y ...
type Y struct {
	FirstName string
	Age       int
}

func main() {

	log.SetFlags(log.Llongfile)

	var b bytes.Buffer

	enco := gob.NewEncoder(&b)
	deco := gob.NewDecoder(&b)

	err := enco.Encode(X{"cesar", "caycho", 16})

	if err != nil {
		log.Fatal(err)
	}

	err = enco.Encode(X{"arturo", "navas", 20})

	if err != nil {
		log.Fatal(err)
	}

	var yAux Y

	err = deco.Decode(&yAux)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(yAux)

	err = deco.Decode(&yAux)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(yAux)
}
