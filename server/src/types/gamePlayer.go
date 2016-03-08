package types

import (
	"fmt"
)

type Player interface {
	GetName() string
	GetId() uint64
}

type Countryer interface {
	GetId() uint32
}
