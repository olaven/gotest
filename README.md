# gotest  
I have written this while familiarizing myself with [go](https://golang.org).
It is not a serious project and should not be treated as such. 
With that out of the way.. 

## Usage 
```go
import (
    "testing", 
    "github.com/olaven/gotest"
)

func TestGotest(t *testing.T) {

    gotest.Test(t).
        equals(2, 2).
        equals("a", "a").
        notEquals(1, 2).
        notEquals("a", "b")
}
```