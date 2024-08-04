package util

import (
	"errors"
	"math/rand"

	"github.com/bwmarrin/snowflake"
	"github.com/sony/sonyflake"
)

// IDGenerator is the inteface for generatring unique ID
type IDGenerator interface {
	NextID() (uint64, error)
}

func NewSonyFlake() (IDGenerator, error) {
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	if sf == nil {
		return nil, errors.New("sonyflake not created")
	}
	return sf, nil
}

type SnowFlake struct {
	node *snowflake.Node
}

func NewSnowFlake() (IDGenerator, error) {
	nodeId := rand.Int63n(1024)
	node, _ := snowflake.NewNode(nodeId)
	sf := new(SnowFlake)
	sf.node = node
	return sf, nil
}

func (sf *SnowFlake) NextID() (uint64, error) {
	return uint64(sf.node.Generate()), nil
}
