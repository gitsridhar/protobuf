package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	//out, err := protojson.Marshal(pb)
	out, err := option.Marshal(pb)

	if err != nil {
		log.Fatalln("CAn't conver to JSON", err)
		return ""
	}

	return string(out)
}

func fromJSON(in string, pb proto.Message) {

	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	//if err := protojson.Unmarshal([]byte(in), pb); err != nil {
	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Couldn't unkarshal from JSON")
	}
}
