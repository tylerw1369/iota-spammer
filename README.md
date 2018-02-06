IOTΛ PoW Confirmation Test Program
==================================

I believe there is a problem with iota.lib.go and using local PoW versus calling `attachToTangle` on the remote node.
In my experience, anytime I generate the PoW locally, the transactions will never get confirmed. When using `attachToTangle`,
however, they get confirmed very soon. I created this program for others to test
with. 

Observations
------------

Here you can see nothing but unconfirmed transactions with a specific tag (orange circles).
None of these transactions ever confirmed. 

![local pow screenshot](https://i.imgur.com/VDpB0Wi.jpg "Screenshot of
transactions with PoW done by iota.lib.go")

This is the same code, but calling `attachToTangle` and having the remote node
do the PoW. You can see most of the transactions confirming as the tangle moves along, just like it should.

![local pow screenshot](https://i.imgur.com/yzAILqI.jpg "Screenshot of
transactions with PoW done by remote node")


Operation
---------

Run the program with the `--help` flag to see available options. It will create a 0 vaule transaction, and ask the node for two tips to 
confirm. It then generates the transaction and sends it off. Using a tangle visualizer such as http://tangle.glumb.de/ you
can see your transactions. Enter the tag given to you when the program started in the tag filter box. **You *must* remove the 
leading `999` from the tag before you enter it.** This is due to a bug that mutates the first two tag trytes on the visualizer. 

You can also just enter `GOPOW` in there to see all transactions created by this program. The default tag is `999GOPOW9<PoW Method>9<random>`.
This lets us filter by PoW method to try and see where the problem lies. 
