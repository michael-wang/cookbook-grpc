package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	json2pb()
}

func json2pb() {
	jsFile := "./repo.json"
	js, err := ioutil.ReadFile(jsFile)
	if err != nil {
		log.Fatalf("Failed to read %s, err: %s\n", jsFile, err)
	}
	fmt.Printf("json: %s\n", js)
	fmt.Println("***************************************************")

	pb := Repo{}
	err = protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}.Unmarshal(js, &pb)
	if err != nil {
		log.Fatalf("Failed to unmarshal json to protobuf, err: %s\n", err)
	}
	fmt.Printf("protobuf: %s\n", prototext.Format(&pb))
}
