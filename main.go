package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/joe-thompson-exa/gcsForwarderDecode"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Read the existing address book.
	in, err := ioutil.ReadFile("test.proto")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", "test.proto")
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	// [START marshal_proto]
	event := &pb.Event{}
	// [START_EXCLUDE]
	if err := proto.Unmarshal(in, event); err != nil {
		log.Fatalln("Failed to parse event:", err)
	}
}
