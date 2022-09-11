# bloomclock

Exploring Bloom Clocks as described in this paper:

[The Bloom Clock](https://www.researchgate.net/publication/333505687_The_Bloom_Clock)


# Tools

- [bloomprob](./examples/bloomprob): Prints the probability for false positives for a
  given bloom filter configuration.

      go run examples/bloomprob/main.go -m 512 -k 2 -n 10
      P(false positive) for m=  512, k=2, n= 10: 0.001470426



