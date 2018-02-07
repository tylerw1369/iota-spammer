IOTΛ Spammer Test Program
==================================

A simple spammer to help increase TPS of the IOTA network. 

Usage
---------
Download the latest spammer build from the [release page](https://github.com/iota-tangle-io/iota-spammer/releases).
If on linux or macOS, make sure to `chmod 770 <file_name>` and `mv <file_name> spammer` the binary.
Run the spammer via `./spammer` (on Linux/macOS) or `spammer.exe` (on Windows). 

**Note**: it is highly suggested to use different nodes when you are running multiple spammer instances.
The specific node to use can be specified via `./spammer --node=<node_url>` (i.e `http://iota-tangle.io:14265`)

The spammer will do PoW locally and ask the specified IRI node for tips and broadcasting of the generated spam TXs.

Using a Hetzner CX51 VPS (8 cores, 32gb ram) instance gives about ~1.1 TPS under optimal conditions with SSE PoW.

The only build on the release page using SSE PoW is the linux/amd64 build, thereby **it's suggested to use a linux system
for maximum throughput**. 

Windows and macOS builds will use PoWGo, you can however build the spammer yourself on those operation
systems to see whether you can enable other PoW methods.

Operation
---------

Run the program with the `--help` flag to see available options. It will create a 0 vaule transaction, and ask the node for two tips to 
confirm. It then generates the transaction and sends it off. Using a tangle visualizer such as http://tangle.glumb.de/ you
can see your transactions. Enter the tag given to you when the program started in the tag filter box. **You *must* remove the 
leading `999` from the tag before you enter it.** You can also filter by any
part of a tag, no need to enter the complete tag. 

You can also just enter `GOPOW` in there to see all transactions created by this program. The default tag is `999GOPOW9<PoW Method>9<random>`.
