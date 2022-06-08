package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	elliot := &Person{
		Name: "Mahmud",
		Age:  19,
	}
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshalling err: ", err)
	}

	fmt.Println(data)

}
