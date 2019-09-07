Demonstrate forking of multiple goroutines and joining them on the main goroutine. In Java, this is normally accompplished with use of a thread pool followed by a wait/shutdown combination. Another way is through the use of the countdown latch.
Notice the simplicity here:
1. No callbacks/runnables/callables/futures/threads.
2. The code also reads in a "linear" fashion and is easy to follow.

Warning:
1. This code is NOT meant to demonstrate how to share data across goroutines, just how to synchronize.
2. There is no error handling. There are other examples which deal with error handling/timeouts etc.
