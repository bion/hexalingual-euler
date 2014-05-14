# ~~Trilingual~~ PENTALINGUAL Project Euler

Using C, Ruby, Javascript, and SuperCollider*

\* for the lulz

## Benchmarks

All times given in seconds

### Problem Six
* Ruby: 0.025 from command line
* C: 0.001 from command line
* Javascript: 0.046 from command line
* SuperCollider: 0.991 from command line, 0.0004 not including
  interpreter start/stop

### Problem Five
* Ruby: 4.605 from command line
* C: 0.211 from command line
* Javascript: 0.345 from command line
* SuperCollider: 35.182 from command line, 34.189 not including
  interpreter start/stop

### Problem Four
* Ruby: 0.361 from command line
* C: 0.049 from command line
* Javascript: 0.345 from command line
* SuperCollider: 2.193 from command line, 1.32 not including
  interpreter start/stop

### Problem Four
* Ruby: 0.361 from command line
* C: 0.049 from command line
* Javascript: 0.345 from command line
* SuperCollider: 2.193 from command line, 1.32 not including
  interpreter start/stop
* Go: 0.0003 seconds - 0.004538 from command line
    (0.29 compile + run from command line)

### Problem Three
* Ruby: 0.030 from command line
* C: 0.001 from command line
* Javascript: 0.066 from command line
* SuperCollider: 0.90 from command line, 0.0002 not including
  interpreter start/stop
* Go: .00010476

Note: SuperCollider can't represent integers over 32 bits, had to use
a float.

### Problem Two
* Ruby: 0.037 from command line
* C: 0.001 from command line
* Javascript: 0.62 from command line
* SuperCollider: 0.89 from command line, 0.0003 not including
  interpreter start/stop
* Go: .000000205 seconds within Go benchmark suite

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

