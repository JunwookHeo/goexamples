package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const dataShards = 2
const parShards = 2

var outFile = flag.String("out", "", "Alternative output path/file")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  simple-decoder [-flags] basefile.ext\nDo not add the number to the filename.\n")
		fmt.Fprintf(os.Stderr, "Valid flags:\n")
		flag.PrintDefaults()
	}
}

func main() {
	// Parse flags
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Error: No filenames given\n")
		flag.Usage()
		os.Exit(1)
	}
	fname := args[0]

	log.Println("filename : ", fname)

	// // Create matrix
	// enc, err := reedsolomon.New(dataShards, parShards)
	// checkErr(err)

	// // Create shards and load the data.
	// shards := make([][]byte, dataShards+parShards)
	// for i := range shards {
	// 	infn := fmt.Sprintf("%s.%d", fname, i)
	// 	log.Println("Opening", infn)
	// 	shards[i], err = ioutil.ReadFile(infn)
	// 	if err != nil {
	// 		log.Println("Error reading file", err)
	// 		shards[i] = nil
	// 	}
	// }

	// // Verify the shards
	// ok, err := enc.Verify(shards)
	// if ok {
	// 	log.Println("No reconstruction needed")
	// } else {
	// 	log.Printf("Verification failed. Reconstructing data : %v", err)
	// 	log.Print("Reconstruct before", shards)
	// 	err = enc.Reconstruct(shards)
	// 	log.Print("Reconstruct after", shards)
	// 	if err != nil {
	// 		log.Println("Reconstruct failed -", err)
	// 		os.Exit(1)
	// 	}
	// 	ok, err = enc.Verify(shards)
	// 	if !ok {
	// 		log.Println("Verification failed after reconstruction, data likely corrupted.")
	// 		os.Exit(1)
	// 	}
	// 	checkErr(err)
	// }

	// // Join the shards and write them
	// outfn := *outFile
	// if outfn == "" {
	// 	outfn = fname + ".out"
	// }

	// log.Println("Writing data to", outfn)
	// f, err := os.Create(outfn)
	// checkErr(err)

	// // We don't know the exact filesize.
	// err = enc.Join(f, shards, len(shards[0])*dataShards)
	// checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(2)
	}
}
