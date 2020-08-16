# gopipes
Emulates executing multiple commands on multiple devices. The sequence of connections to the devices is random. Commands are executed one after the other in goroutines. The duration of connecting and executing each command is 1 second for testing.

For comparison, the "syncpipes" folder contains the synchronous version.
