package grpc

import (
	"time"

	"github.com/qwiltech/go-config/source"
	proto "github.com/qwiltech/go-config/source/grpc/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
