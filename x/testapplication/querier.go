package testapplication

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

)


//query endpoints supported by the nameservice Querier


const (
	QueryResolve = "resolve" // used to get the value of the hash
	QueryBol = "bol"  // used for to query the entire bill of lading data
	QueryHashes = "hashes" // used for the interator to get all the hashes
)

//NewQuerier is the module level router for state queries

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryBol:
			return queryBol(ctx, path[1:], req, keeper)
		case QueryHashes:
			return queryHashes(ctx, req, keeper)

		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

// QueryResolve config
//nolint: unparam for queryResolve

func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	hash := path[0]

	value := keeper.GetHash(ctx, hash)

	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve hash")
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResResolve{value})
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// Query Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}

// QueryBol config
// nolint: unparam
func queryBol(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	hash := path[0]

	bol := keeper.GetBol(ctx, hash)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, bol)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// implement fmt.Stringer
func (b Bol) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Retrieve at: %s`, b.Owner, b.Value, b.Retrieve))
}


//iterator config
func queryHashes(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var hashesList QueryResHashes

	iterator := keeper.GetHashesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		hash := string(iterator.Key())
		hashesList = append(hashesList, hash)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, hashesList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// Query Result Payload for a names query
type QueryResHashes []string

// implement fmt.Stringer
func (n QueryResHashes) String() string {
	return strings.Join(n[:], "\n")
}