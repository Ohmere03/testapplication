package testapplication


import (

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"


	/*codec - the codec provides tools to work with the Cosmos encoding format, Amino
	.
	bank
	- the bank module controls accounts and coin transfers.
	types
	- types contains commonly used types throughout the SDK.

	 */

)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey  sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

/*
bank.Keeper - This is a reference to the Keeper from the bank module. Including it allows code in this module to call functions from the bank module. The SDK uses an object capabilities
approach to accessing sections of the application state. This is to allow developers to employ a least authority approach, limiting the capabilities of a faulty or malicious module from affecting parts of state it doesn't need access to.
*codec.Codec
- This is a pointer to the codec that is used by Amino to encode and decode binary structs.
sdk.StoreKey
- This is a store key which gates access to a sdk.KVStore which persists the state of your application: the Bol struct that the name points to (i.e. map[hash]Bol).
 */


// Sets the entire Bol metadata struct for a hash
func (k Keeper) SetBol(ctx sdk.Context, hash string, bol Bol) {
	if bol.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(hash), k.cdc.MustMarshalBinaryBare(bol))
}


// get the entire Bol metada struct for a hash

func (k Keeper) GetBol(ctx sdk.Context, hash string) Bol{
	store := ctx.KVStore(k.storeKey)
	//if !store.Has([]byte(hash)){
		//return
	//}
	bz := store.Get([]byte(hash))
	var bol Bol
	k.cdc.MustUnmarshalBinaryBare(bz, &bol)
	return bol
}



// Below are function to get specific parameters from the store based on the hash We use the getter and setter function

//Return the value of the hash to be used for hash check and comparison for example
func (k Keeper) GetHash(ctx sdk.Context, hash string) string{
	return k.GetBol(ctx, hash).Value
}

// SetOwner function to change the owner if the bol is transmitted
func (k Keeper)  SetOwner(ctx sdk.Context, hash string, owner sdk.AccAddress){
	bol := k.GetBol(ctx, hash)
	bol.Owner = owner
	k.SetBol(ctx, hash, bol)

}

// Return the owner of the bill of lading

func(k Keeper) GetOwner(ctx sdk.Context, hash string) sdk.AccAddress{
	bol := k.GetBol(ctx, hash)
	address := bol.Owner
	return address
}
//Get where it is stored

func (k Keeper) GetRetrieved(ctx sdk.Context,hash string) string{
	bol := k.GetBol(ctx, hash)
	retrieve := bol.Retrieve
	return retrieve
}
// Get an iterator over all hashes in which the keys are the names and the values are the whois
func (k Keeper) GetHashesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// NewKeeper creates new instances of the testapplication Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}
