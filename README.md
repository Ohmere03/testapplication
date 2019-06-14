# testapplication

This application is adapted from https://cosmos.network/docs/tutorial/ for a Master Thesis made at the University of Lausanne.

The goal of the application is to be able for a member of the network to create a bill of lading, and transmit the bill of lading to another user of the network. A bill of lading object contains:
-the value of the hash
-who is the owner

The owner can also send coins from one address to another to simulate the functionality of a cryptocurrency.

Finally, users can query the different addresses to know how many coins an account has or query the bill of lading object to get the hash and/or the owner.


Golang 1.12 was used in the implementation. To build and install, check out the cosmos doc at this address:
https://cosmos.network/docs/tutorial/build-run.html


The commands can be issued via terminal. A web application demonstration was build and can be found at this address:

httpsa://github.com/Ohmere03/electron-masterthesis