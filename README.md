# testapplication

This application is a test adapted from https://cosmos.network/docs/tutorial/


The goal of the application is to be able for a member of the network to create a bill of lading, the bol will contain:
-who is the owner
-the value of the hash
-where to retrieve it

We will use one store to map the hash to the bol containing the owner, the hash again and where to retrieve it.


We will need to messages:
-MsgCreateBol: This message will allow someone to create the bol and initialize the hash, the owner and where to retrieve it.
-MsgTransmit: This message will allow to change the owner of the bol.

When transactions reaches Tendermint, they are decoded to get the message and routed to the module using the Handler. If the state needs to be updated, the Keeper is used for this effect. (Nameservice tutorial)

