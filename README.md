# go-xor

### Usage
```go
var bs []bool
bs = []bool{true, false, false}
fmt.Println(xor.XOR(bs)) // true

bs = []bool{true, false, true}
fmt.Println(xor.XOR(bs)) // false

bs = []bool{false, false, false}
fmt.Println(xor.XOR(bs)) // false
```