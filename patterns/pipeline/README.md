# note

`$ go run main.go`

```
[main] start func
[main] define intStream
[generator] start func
[generator] return channel
[main] define pipeline
[multiply] start func
[multiply] return channel
[add] start func
[add] return channel
[multiply] start func
[multiply] return channel
[main] wait for pipeline...
[generator] 1 start
[generator] 1 send to stream
[generator] end select in goroutine
[generator] 2 start
[generator] 2 send to stream
[generator] end select in goroutine
[generator] 3 start
[generator] 3 send to stream
[generator] end select in goroutine
[generator] 4 start
[generator] 4 send to stream
[generator] end select in goroutine
[generator] end for in goroutine
[multiply] 1 start
[multiply] 1 * 2 send to stream
[multiply] end select in goroutine
[multiply] 2 start
[add] 2 start
[add] 2 + 1 send to stream
[add] end select in goroutine
[add] 4 start
[multiply] 3 start
[multiply] 3 * 2 send to stream
[multiply] end select in goroutine
[multiply] 5 start
[multiply] 2 * 2 send to stream
[multiply] end select in goroutine
[multiply] 3 start
[main] received 6 from pipeline
6
[main] received 10 from pipeline
10
[multiply] 5 * 2 send to stream
[multiply] end select in goroutine
[add] 4 + 1 send to stream
[add] end select in goroutine
[add] 6 start
[add] 6 + 1 send to stream
[add] end select in goroutine
[multiply] 7 start
[multiply] 7 * 2 send to stream
[multiply] end select in goroutine
[main] received 14 from pipeline
14
[multiply] 3 * 2 send to stream
[multiply] end select in goroutine
[multiply] 4 start
[multiply] 4 * 2 send to stream
[multiply] end select in goroutine
[multiply] end for in goroutine
[add] 8 start
[add] 8 + 1 send to stream
[add] end select in goroutine
[add] end for in goroutine
[multiply] 9 start
[multiply] 9 * 2 send to stream
[multiply] end select in goroutine
[multiply] end for in goroutine
[main] received 18 from pipeline
18
[main] end func
```
