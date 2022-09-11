package main

import (
	"flag"
	"fmt"
	bloomclock "github.com/tc1io/bloom-clock-go"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s: Print the probability of false positives for a given bloom filter configuration.\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")

		flag.PrintDefaults()
	}

	m := flag.Int("m", 256, "Size of the bloom clock")
	k := flag.Int("k", 2, "Number of hash functions used")
	n := flag.Int("n", 10, "Number of elements stored")
	flag.Parse()

	p := bloomclock.Probability(float64(*m), float64(*k), float64(*n))

	fmt.Printf("P(false positive) for m=%5d, k=%d, n=%3d: %0.9f\n", *m, *k, *n, p)
}
