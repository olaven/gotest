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
		Equals(2, 2).
		Equals("a", "a").
		NotEquals(1, 2).
		NotEquals("a", "b")
}
```