# ~~Trilingual~~ Quadrilingual Project Euler

Using C, Ruby, Javascript, and SuperCollider*

\* for the lulz

## Benchmarks

All times given in seconds

### Problem Four
* Ruby: 0.361 from command line
* C: 0.049 from command line
* Javascript: 0.345 from command line
* SuperCollider: 2.193 from command line, 1.32 not including
  interpreter start/stop

### Problem Three
* Ruby: 0.030 from command line
* C: 0.001 from command line
* Javascript: 0.066 from command line
* SuperCollider: 0.90 from command line, 0.0002 not including
  interpreter start/stop
* Go: 0.0005320 

Note: SuperCollider can't represent integers over 32 bits, had to use
a float.

### Problem Two
* Ruby: 0.037 from command line
* C: 0.001 from command line
* Javascript: 0.62 from command line
* SuperCollider: 0.89 from command line, 0.0003 not including
  interpreter start/stop
* Go: .000205 seconds within Go benchmark suite

### Problem One
* Ruby: 0.049 from command line
* C: 0.001 from command line
* Javascript: 0.056 from command line
* SuperCollider: 1.01 from command line, 0.0004 not including
  interpreter start/stop
* Go:
BenchmarkThreeFiveOneK	 1000000	      1580 ns/op
BenchmarkThreeFiveTenK	  200000	      9798 ns/op
BenchmarkThreeFiveMill	    2000	    902160 ns/op

