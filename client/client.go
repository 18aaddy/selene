package client

import (
	"github.com/BlocSoc-iitr/selene/config"
)

type ClientBuilder struct {
	Network *config.Network
	ConsensusRpc *string
	ExecutionRpc *string
	Checkpoint *[32]byte
	RpcBindIp *string
	RpcPort *uint16
	DataDir *string
	Config *config.Config
	Fallback *string
	LoadExternalFallback bool
	StrictCheckpointAge bool
}