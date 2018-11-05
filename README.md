# note-concurrency

## Go
### run
  + `$ go run main.go`

```
sample1> go run main.go
# command-line-arguments
.\main.go:8:3: undefined: wgsample
```

```
sample1> go run main.go waitgroup.go
2nd goroutine sleeping...
1st goroutine sleeping...
2nd goroutine wake up
1st goroutine wake up
All goroutines complete.
finish
```

- `mutexsample`

```
start
---------------------
Incrementing: 1
Decrementing: 0
Incrementing: 1
Incrementing: 2
Incrementing: 3
Incrementing: 4
Incrementing: 5
Decrementing: 4
Decrementing: 3
Decrementing: 2
Decrementing: 1
Decrementing: 0
Arithmetic complete.
---------------------
finish
```

- `procon`

```
start
---------------------
Readers  RWMutex    Mutex
1        7.0911ms   7.1031ms
2        6.2342ms   6.8817ms
4        7.1439ms   5.4234ms
8        8.9125ms   7.2748ms
16       8.2982ms   8.0957ms
32       5.992ms    6.2519ms
64       7.0073ms   6.0021ms
128      7.1991ms   7.4257ms
256      7.9944ms   6.8938ms
512      5.5176ms   6.9309ms
1024     7.995ms    7.3431ms
2048     7.1171ms   8.0003ms
4096     6.4871ms   8.1109ms
8192     7.0008ms   6.4987ms
16384    7.9918ms   6.008ms
32768    7.0262ms   6.0076ms
65536    10.9997ms  13.9998ms
131072   22.9991ms  28.001ms
262144   45.999ms   57.0013ms
524288   92.9986ms  114.0011ms
---------------------
finish
```
