Go is a statically typed, each variable has a type, it is known and fixed at compile time.  For example int, float, []byte, [5]byte, and so on. And even if the underlying type is the same, they are seen as distinct types.  

```golang
type MyInt int

var i int
var j MyInt
```
i and j are distinct types, and i can't be assigned to j.  

---
### respresation of interface
one important category of type is interface, it is considered as a collection of function signatures, actually the interface variable can be seen to have two members, one is the actual value, the other is the actual value type description. 

```golang
var r io.Reader
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
if err != nil {
    return nil, err
}
r = tty
```
r contains tty value and its type *os.File which implements the io.Reader interface.

```
var w io.Writer
w = r.(io.Writer)
```
r has the its original type description, so it can check if it implements io.Writer, if it does, it can be assigned to w.

---
### reflection
#####Reflection goes from interface value to reflection object.
At the basic level, reflection is just a mechanism to examine the type and value pair stored inside an interface variable. 
