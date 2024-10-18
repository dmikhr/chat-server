package data

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimeToTimestamppb - конвертация time.Time в protobuf timestamppb.Timestamp
func TimeToTimestamppb(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}

// TimestamppbToTime - конвертация protobuf timestamppb.Timestamp в time.Time
func TimestamppbToTime(timeproto *timestamppb.Timestamp) time.Time {
	return timeproto.AsTime()
}
