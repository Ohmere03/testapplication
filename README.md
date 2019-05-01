# testapplication

This application is a test adapted from https://cosmos.network/docs/tutorial/


The goal of the application is to be able for a member of the network to create a bill of lading, the bol will contain:
-the value of the hash
-who is the owner
-where to retrieve it

We will use one store to map the hash to the bol containing the owner, the hash again and where to retrieve it.


We will need to messages:
-MsgCreateBol: This message will allow someone to create the bol and initialize the hash, the owner and where to retrieve it.
-MsgTransmit: This message will allow to change the owner of the bol.

When transactions reaches Tendermint, they are decoded to get the message and routed to the module using the Handler. If the state needs to be updated, the Keeper is used for this effect. (Nameservice tutorial)


the app.go will bring all the modules together and link them to make the application run. It defines what is done when a transaction is received.

the x/test application folder contains:

types.go:

Where the Bol struct is defined and a function is created that create a Bol on call with the owner, the 
value of the hash and where to find it

keeper.go: the keeper is the main core of the module and handle interaction with the store, reference to other keeper if we implement any cross module interactions and contain most of the functionalities.

In the keeper file the setter and getter are defined to store and get access to the Bol struct.

For the functions to modify the Bol, following are implemented:

-GetHash
-SetOwner
-GetOwner
-GetRetrieve

There is also an iterator "GetHashesIterator" and a constructor NewKeeper

The Msgs will induce state transitions. They are wrapped in Txs that clients submit to the network. Msgs must satisfy a specific interface.

_____________________________MESSAGE INTERFACE___________________________________________________________


// Transactions messages must fulfill the Msg
type Msg interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	Type() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Route() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() Error

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []AccAddress
}


____________________________________________________________________________________________________________

The handlers define the action that needs to be taken(action on the stores and what are the conditions) when Msg arrives.

We will have Two messages: CreateBol and TransmitBol.

CreateBol will be used to create the initial Bol and store it on the network and TransmitBol will be used when the bol changes the owner.

In the handler file we have a router that route the message to the proper handler that than execute specfiic actions



The querrier.go file define the queries the application state users will be able to make. We need two queries:
1)resolve: takes a hash and return the value of the hash of the document
2) bol: takes a hash and return the entire metadata of the bol stored in the state : owener, value and retrieve



