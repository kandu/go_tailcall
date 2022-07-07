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
    }
    fmt.Printf("a(): %d\n", r)
}

