IOTΛ Spammer Test Program
==================================

A simple spammer to help increase TPS of the IOTA network. 

Operation
---------

Run the program with the `--help` flag to see available options. It will create a 0 vaule transaction, and ask the node for two tips to 
confirm. It then generates the transaction and sends it off. Using a tangle visualizer such as http://tangle.glumb.de/ you
can see your transactions. Enter the tag given to you when the program started in the tag filter box. **You *must* remove the 
leading `999` from the tag before you enter it.** You can also filter by any
part of a tag, no need to enter the complete tag. 

You can also just enter `GOPOW` in there to see all transactions created by this program. The default tag is `999GOPOW9<PoW Method>9<random>`.
