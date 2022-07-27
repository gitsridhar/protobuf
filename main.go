package main

import (
	"fmt"

	pb "github.com/gitsridhar/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A Name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 42, Name: "My Name"},
		MultipleDummies: []*pb.Dummy{
			{Id: 43, Name: "My Name 2"},
			{Id: 44, Name: "My Name 3"},
		},
	}
}
func main() {
	fmt.Println(doSimple())
	fmt.Println(doComplex())
}
