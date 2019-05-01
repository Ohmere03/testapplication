package app

import (
	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/x/auth"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tendermint/libs/db"

)

const (
	appName ="billOfLading"
)

type bolApp struct {
	*bam.BaseApp
}

func NewBolApp (logger log.Logger, db dbm.DB) *bolApp{
	// First define the top level codec that will be shared by the different modules. Note: Codec will be explained later
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &bolApp{
		BaseApp: bApp,
		cdc:     cdc,
	}

	return app

}

