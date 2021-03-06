# goutil

## concurrency
Purpose:
1. Learn concurrency Go style for Java developers:
	1. Helping Java developers get trained in the Go programming language from a point of view of Java concurrency.
	2. Help Java developers migrate from Java's Thread/Future based concurrency to Go's Goroutine/Channel based concurrency.
	3. Reproduce common concurrency patterns used in Java through Go's model of concurrency.

Layout of samples (source code contains documentation to introduce ideas/concepts)

1. ch1 - Demonstrates forking and joining of goroutines. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch1/forkjoin/main.go)
2. ch2 - Futures demo (while there are no futures in Go, helps to bridge from Java to Go)
   1. future - Simple channel demo to communicate data across 2 goroutines. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch2/future/main.go)
   2. future2 - A "streaming" channel. This is when a channel communicates more than one value.(https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch2/future2/main.go) 
   3. future3 - A pipeline - chaining of channels. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch2/future3/main.go)
3. ch3 - Demonstrates generator/streaming pattern.
   1. basic - A simple generator. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch3/generator/basic/main.go)
   2. advanced1 - The consumer chooses when to stop the generator with a "done" channel. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch3/generator/advanced1/main.go)
   3. advanced2 - The producer chooses when to stop the generator based on some condition. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch3/generator/advanced2/main.go)
4. ch4 - fanin - Demonstrates the fanin pattern - multiple tasks are generated by multiple sources and "fannedin" to one goroutine for processing.
   1. basic - A single task generator produces multiple tasks. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch4/fanin/basic/main.go)
   2. advanced - Multiple task generators produde multiple tasks. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch4/fanin/advanced/main.go)
5. ch5 - fanout - Demonstrates the fanout process - Tasks are generated by a single goroutine and "fanned" out to multiple goroutines. Also demonstrates uses of sync.Mutex and sync.Waitgroup. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch5/fanout/main.go)
6. ch6 - servicebus - Demonstrates the fanin and fanout process - Tasks are generated by a single goroutine and "fannedin" to one goroutine. This goroutine in turns "fans" out to multiple goroutines. One use of this is as a serivce bus. Also demonstrates uses of sync.Mutex and sync.Waitgroup. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch6/servicebus/main.go)
7. ch7 - Reproduction of the CompletableFuture pattern.
   1. compose - A simple generator. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch7/compose/main.go)
   2. advancedcompose - The consumer chooses when to stop the generator with a "done" channel. Handles errors as well as cancellation. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch7/advancedcompose/main.go)
8. ch8 - first response - Demonstrates returning first response from a set of requests. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch8/main.go)
9. ch9 - fanin - Demonstrates throttling - ensuring that a serivce is not over-burdened with requests.
   1. basic - Basic throttling without timeout. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch9/throttle/basic/main.go)
   2. advanced - Throttling with timeout. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch9/throttle/advanced/main.go)
10. Circuit breaker - Reproduce the functionality provided by Netflix circuit breaker. (https://github.com/Netflix/Hystrix/wiki/How-it-Works) Flowchart: (https://raw.githubusercontent.com/wiki/Netflix/Hystrix/images/hystrix-command-flow-chart.png)
    Requirements of the circuit breaker:
    1. Fast fail - Client never blocks - either executes request or returns error with explanation.
    2. Auto repair - Circuit will repair itself after it is shutdown due to excess load.
    3. Default behavior - Breaker will execute default behavior in service failure scenarios. This is followed by call to cancel behavior.
    4. Cancel behavior - Breaker will execute cancel behavior in case of service failure. Cancel will be called after default.
    5. Timeout - Clients can configure timeouts, after which breaker will execute default and cancel behavior in order.
    Source code : (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch10/breaker/v1/breaker.go)
    Test code - contains examples that can be copied and pasted for working with the circuit. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch10/breaker/v1/breaker_test.go)
   


## http server and client
Purpose:
1. Demonstrate simple and advanced uses of the http standard library.
2. Demonstrate setup of TLS and mutual TLS
3. Demonstrate timeouts for client and server.
4. Demonstrate building proxy servers.
5. Demonstrate client side tracing of requests.

