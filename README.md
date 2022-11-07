# Go TailCall

Functions in this go module implement tail calls via trampolining. This module depends on Go's new generics language feature, so Go 1.18 or newer is required.

Tail calling functions return their result using TailRet or call the next function using TailCall. Both return a TailRec struct.

The result of evaluating a tailcalling function can be retrieved from using function TailStart.

## Installation


To install go\_tailcall module, you can use the below command

```sh
go get -u github.com/kandu/go_tailcall@0.3
```

## Quick start

``` go
package main

import . "github.com/kandu/go_tailcall/tailcall2"
import "fmt"
import "log"
import "strconv"

func a(v int64) TailRec[int64, error] {
    v++
    fmt.Printf("a %d\n", v)

    var s= strconv.FormatInt(v, 10)
    return TailCall(b, s)
}

func b(s string) TailRec[int64, error] {
    s= s + "0"
    fmt.Printf("b %s\n", s)

    var v, _= strconv.ParseInt(s, 10, 64)
    var limit int64= 1 << 50
    if v >= limit {
        return TailRet[int64, error](v, nil)
    } else {
        return TailCall(a, v)
    }
}

func main() {
    var r, err= TailStart(a(0))
    if err != nil {
        log.Fatal(err)
    } else {
      fmt.Printf("a(): %d\n", r)
    }
}
```

And below are infinite mutually recursive functions. They will run forever and will not suffer from stack overflow.

``` go
package main

import . "github.com/kandu/go_tailcall/tailcall"
import "fmt"
import "strconv"


func a(v int64) TailRec[int64] {
    v++
    fmt.Printf("a %d\n", v)

    var s= strconv.FormatInt(v, 10)
    return TailCall(b, s)
}

func b(s string) TailRec[int64] {
    s= s + "0"
    fmt.Printf("b %s\n", s)

    var v, _= strconv.ParseInt(s, 10, 64)
    var limit int64= 1 << 50
    if v >= limit {
        return TailCall(a, 0)
    } else {
        return TailCall(a, v)
    }
}


func main() {
    fmt.Println("infinite mutually recursive functions")
    TailStart(a(0))
}
```

