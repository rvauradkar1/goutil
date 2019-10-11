# goutil

Purpose of this repository:
1. Learn concurrency Go style for Java developers:
	1. Helping Java developers get trained in the Go programming language from a point of view of Java concurrency.
	2. Help Java developers migrate from Java's Thread/Future based concurrency to Go's Goroutine/Channel based concurrency.
	3. Reproduce common concurrency patterns used in Java through Go's model of concurrency.

Layout of samples (source code contains documentation to introduce ideas/concepts)

1. ch1 - Demonstrates forking and joining of goroutines. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch1/forkjoin/main.go)
2. ch2 - Futures demo (while there are not futures in Go, this is an instructional tool)
   1. future - Simple channel demo to communicate data across 2 goroutines. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch2/future/main.go)
   2. future2 - A "streaming" channel. This is when a channel communicates more than on value.(https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch2/future2/main.go) 
   3. future3 - A pipeline - chaining of channels. (https://github.com/rvauradkar1/goutil/blob/master/concurrent/ch2/future2/main.go)

  1.2.1 future2


2. Learn Web development
	1. Basic setup of a Http/Https Webserver.
	2. Folder structures for Web apps.
	3. Demonstrate proxy/reverse-proxy capabilities.

As more capabilties are added, this page will be updated.
