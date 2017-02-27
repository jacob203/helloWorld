when you go routine, there is no handle which can be returned, it is designed to be like that. because the designed thinks independent routine is created and you shouldn't intervene it. if you want to talk with that go routine, you can use channel which is designed to be like sync-call, so as to a no buffer channel, both sides are blocked expect if both osides are ready.

# function that returns a channel
channels are first -class values, just like strings or integers.
channels are used to communicate between routines, so if there is a function which returns a channel, it means that the function should create a go routine

# multiplexing
use a channel to fanin multiple channels.

# functions are also a first-class values, just like channels

use channels and go routines, you don't need locks, condition variables and no callbacks.

# Rob Pike "Concurrency is not parallelism"
for example, here is a task, move a pile of bricks to another place.
 1. parallelism means to employ more workers to move a brick to another place
 2. concurrency means to break down the move action, for example, one for loading a brick on a cart, one for moving, one for unloading a brick from a cart, and employ workers to do these three actions, and use channels to move bricks.

 ```
type request struct {
    action func() int
    chan int
}
 ```
 send a request to the server and when the server finishes the action, use the channel to return the result, a channel is a first-class value

 # readable means reliable