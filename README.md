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
