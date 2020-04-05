package cg

import (
	"errors"
	"encoding/json"

	"ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}