package cg

import (
	"encoding/json"
	"errors"
	"sync"

	"ipc"
)

var _ ipc.Server = &CenterServer