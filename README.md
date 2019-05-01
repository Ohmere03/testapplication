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
-GetRetrieved

There is also an iterator "GetHashesIterator" and a constructor NewKeeper