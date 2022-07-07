package protobuf

import (
	pb "pk/protobuf/proto/test"

	"google.golang.org/protobuf/proto"
)

func marshal() {
	award := &pb.Award{
		Uid:      628,
		AwardId:  1,
		UserName: "Haruka",
	}
	msg, err := proto.Marshal(award)
}

func unmarshal() {
	award := &pb.Award{}
	if err := proto.Unmarshal(msg.Value, award); err != nil {
		panic(err)
	}
}
