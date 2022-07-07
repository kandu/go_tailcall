package tailcall6

// TailRec[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] reprents a tailcall return value of which the status is done or waiting to be evaluated
type TailRec[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
    next func() TailRec[T1, T2, T3, T4, T5, T6]
    result1 T1
    result2 T2
    result3 T3
    result4 T4
    result5 T5
    result6 T6
    done bool
}

// return TailRet[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] a of which the status is done with the final value
func TailRet[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](r1 T1, r2 T2, r3 T3, r4 T4, r5 T5, r6 T6) TailRec[T1, T2, T3, T4, T5, T6] {
    return TailRec[T1, T2, T3, T4, T5, T6] {
        result1: r1,
        result2: r2,
        result3: r3,
        result4: r4,
        result5: r5,
        result6: r6,
        done: true,
    }
}

// start to get the final result from a TailRec[T1, T2, T3, T4, T5, T6]
func TailStart[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](tr TailRec[T1, T2, T3, T4, T5, T6]) (T1, T2, T3, T4, T5, T6) {
    for ;; {
        if tr.done {
            return tr.result1, tr.result2,
                tr.result3, tr.result4,
                tr.result5, tr.result6
        } else {
            tr= tr.next()
        }
    }
}

// TailCall helper to call a function whose arity is 0
func TailCall0[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](f func() TailRec[T1, T2, T3, T4, T5, T6]) TailRec[T1, T2, T3, T4, T5, T6] {
    return TailRec[T1, T2, T3, T4, T5, T6] {
        next: f,
        done: false,
    }
}

// the default TailCall function, elegant and still, used for connecting the next step of the computation
func TailCall[O1 any, O2 any, O3 any, O4 any, O5 any, O6 any, I any](f func(i I) TailRec[O1, O2, O3, O4, O5, O6], i I) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 2
func TailCall2[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any](
    f func(i1 I1, i2 I2) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 3
func TailCall3[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any](
    f func(i1 I1, i2 I2, i3 I3) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 4
func TailCall4[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 5
func TailCall5[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 6
func TailCall6[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 7
func TailCall7[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 8
func TailCall8[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 9
func TailCall9[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 10
func TailCall10[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 11
func TailCall11[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 12
func TailCall12[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 13
func TailCall13[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 14
func TailCall14[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 15
func TailCall15[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 16
func TailCall16[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 17
func TailCall17[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 18
func TailCall18[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 19
func TailCall19[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 20
func TailCall20[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 21
func TailCall21[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20, i21)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 22
func TailCall22[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20, i21, i22)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 23
func TailCall23[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20, i21, i22, i23)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 24
func TailCall24[
    O1 any, O2 any, O3 any, O4 any, O5 any, O6 any,
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any, I24 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) TailRec[O1, O2, O3, O4, O5, O6],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) TailRec[O1, O2, O3, O4, O5, O6] {
    var g= func () TailRec[O1, O2, O3, O4, O5, O6] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20, i21, i22, i23, i24)
    }
    return TailRec[O1, O2, O3, O4, O5, O6] {
        next: g,
        done: false,
    }
}

