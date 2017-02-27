# no buffer channel
By default sends and receives block until both the sender and receiver are ready.
This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
```golang
package main

import "fmt"

func main() {

    // Create a new channel with `make(chan val-type)`.
    // Channels are typed by the values they convey.
    messages := make(chan string)

    // _Send_ a value into a channel using the `channel <-`
    // syntax. Here we send `"ping"`  to the `messages`
    // channel we made above, from a new goroutine.
    go func() { messages <- "ping" }()

    // The `<-channel` syntax _receives_ a value from the
    // channel. Here we'll receive the `"ping"` message
    // we sent above and print it out.
    msg := <-messages
    fmt.Println(msg)
}
```
### usage
##### a signal something happens
a signal that a go routine is done
```
signalCh := make(chan bool)

go func() {
    //do something
    signalCh <- true
}
//do something
<- signalCh//wait go routine is done

```
but it isn't advised to do it like this, as go routine is blocked if doing something in main thread is time consuming, the routine is fast, when it reaches the singalCh, while the main
thread doesn't go to the waiting of signalCh, the routine is blocked there and the resources that the routine holds can't be released.
# advance usage
### pipeline
```
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
func main() {
    // Set up the pipeline.
    c := gen(2, 3)
    out := sq(c)

    // Consume the output.
    fmt.Println(<-out) // 4
    fmt.Println(<-out) // 9
}
```
it setups a pipeline that firstly gen random numbers, second square numbers one by one.
the following is more like a pipeline.
```
func main() {
    // Set up the pipeline and consume the output.
    for n := range sq(sq(gen(2, 3))) {
        fmt.Println(n) // 16 then 81
    }
}
```

### Fan-out, fan-in
Fan-out **means** multiple functions can read from the same channel until that channel is closed, one producer and multiple consumers.
```
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func main() {
    in := gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    c1 := sq(in)
    c2 := sq(in)

    // Consume the merged output from c1 and c2.
    for n := range merge(c1, c2) {
        fmt.Println(n) // the output result depends on go routines, it may 4,4,9,9 or 4,9,4,9
    }
}
```
There is a pattern to our pipeline functions:
* stages close their outbound channels when all the send operations are done.
    make sure closing the channel in the same function, but it isn't limited to the function lifetime, I means only the function scope is ok.
* stages keep receiving values from inbound channels until those channels are closed.
    resource leak will happen if the inbound data stream is break when some error happens, receivers only get part of the data as the goroutines still holds memory and runtime resources.

so we need to find a way to fix this issue, we need a signal to tell all senders that the receivers doesn't receive any data any more and senders should stop, it is like a broadcast sigal.
a close chan is the right choice, put the signal chan as a parameter for each functions.
```
for n := range in {
            select {
            case out <- n * n:
            case <-done:
                return
            }
        }
```
if you want to control how many go routines are running and wait for all goroutines return, use sync.WaitGroup, do just like the following:
```
    // Start a fixed number of goroutines to read and digest files.
    c := make(chan result)
    var wg sync.WaitGroup
    const numDigesters = 20
    wg.Add(numDigesters)
    for i := 0; i < numDigesters; i++ {
        go func() {
            digester(done, paths, c)
            wg.Done()
        }()
    }
    go func() {
        wg.Wait()
        close(c)
    }()
```
