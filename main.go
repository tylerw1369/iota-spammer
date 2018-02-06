/*
MIT License

Copyright (c) 2018 iota-tangle.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/CWarner818/giota"
	flag "github.com/ogier/pflag"
	"github.com/paulbellamy/ratecounter"
)

var (
	randomTag     string
	randomAddress string
	alphabet      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// Generate random tag suffix to distinguish instances
	for i := 0; i < 3; i++ {
		randomTag += string(alphabet[rand.Intn(len(alphabet))])
	}

	// Send to a random address for easy confirmation checking on thetangle
	alphabet := alphabet + "9"
	for i := 0; i < 81; i++ {
		randomAddress += string(alphabet[rand.Intn(len(alphabet))])
	}
}

func main() {
	var mwm *int64 = flag.Int64("mwm", 14, "minimum weight magnitude")
	var depth *int64 = flag.Int64("depth", giota.Depth, "depth for tip finding")
	var destAddress *string = flag.String("address", "<random>", "address to send to")
	var tag *string = flag.String("tag", "999GOPOW9<pow>9<random>", "transaction tag")
	var server *string = flag.String("node", "http://localhost:14265", "remote node to connect to")
	var remotePoW *bool = flag.Bool("remote-pow", false, "do PoW on remote node using attachToTangle API")
	flag.Parse()

	seed := giota.NewSeed()

	// Set a random address if none was specified
	if *destAddress == "<random>" {
		*destAddress = randomAddress
	}
	recipientT, err := giota.ToAddress(*destAddress)
	if err != nil {
		panic(err)
	}

	log.Println("Using IRI server:", *server)

	api := giota.NewAPI(*server, nil)

	// Get the most performant PoW method for this system
	name, pow := giota.GetBestPoW()
	if *remotePoW {
		pow = nil
		name = "attachToTangle"
	}

	// Set a random tag if the user didnt specify one
	if *tag == "999GOPOW9<pow>9<random>" {
		*tag = "999GOPOW9" + strings.ToUpper(name) + "9" + randomTag
	}

	ttag, err := giota.ToTrytes(*tag)
	if err != nil {
		panic(err)
	}

	trs := []giota.Transfer{
		giota.Transfer{
			Address: recipientT,
			Value:   0,
			Tag:     ttag,
		},
	}

	log.Println("Using tag: http://thetangle.org/tag/" + *tag)
	log.Println("Using address: http://thetangle.org/address/" + *destAddress)
	log.Println("Using PoW:", name)

	// Setup 1/5/15 minue average TPS counters
	r1 := ratecounter.NewRateCounter(1 * time.Minute)
	r5 := ratecounter.NewRateCounter(5 * time.Minute)
	r15 := ratecounter.NewRateCounter(15 * time.Minute)
	for {
		trytes, err := giota.PrepareTransfers(api, seed, trs, nil, "", 1)
		if err != nil {
			log.Println("Error preparing transfer:", err)
			continue
		}

		// This is where the PoW is done right before sending the txn
		err = giota.SendTrytes(api, *depth, []giota.Transaction(trytes), *mwm, pow)
		if err != nil {
			log.Println("Error sending trytes:", err)
			continue
		}

		// Increment counters
		r1.Incr(1)
		r5.Incr(1)
		r15.Incr(1)

		log.Println("SENT: http://thetangle.org/transaction/" + trytes[0].Hash())
		// 1/5/15 min TPS averages
		log.Printf("TPS: %.3f %.3f %.3f\n", float64(r1.Rate())/float64(60), float64(r5.Rate())/float64(60*5), float64(r15.Rate())/float64(60*15))
	}
}
