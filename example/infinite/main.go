package main

import . "github.com/kandu/go_tailcall"
import "fmt"
import "strconv"


func a(v int64) TailRec[int64] {
    v++
    fmt.Printf("a %d\n", v)

    var s= strconv.FormatInt(v, 10)
    return TailCall(b,s)
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

