package tailcall0

// TailRec reprents a tailcall return value of which the status is done or waiting to be evaluated
type TailRec struct {
    next func() TailRec
    done bool
}

// return a TailRec of which the status is done
func TailRet() TailRec {
    return TailRec {
        done: true,
    }
}

// start to get the final result from a TailRec
func TailStart(tr TailRec) {
    for ;; {
        if !tr.done {
            tr= tr.next()
        }
    }
}

// TailCall helper to call a function whose arity is 0
func TailCall0(f func() TailRec) TailRec {
    return TailRec {
        next: f,
        done: false,
    }
}

// the default TailCall function, elegant and still, used for connecting the next step of the computation
func TailCall[I any](f func(i I) TailRec, i I) TailRec {
    var g= func () TailRec {
        return f(i)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 2
func TailCall2[
    I1 any, I2 any](
    f func(i1 I1, i2 I2) TailRec,
    i1 I1, i2 I2) TailRec {
    var g= func () TailRec {
        return f(i1, i2)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 3
func TailCall3[
    I1 any, I2 any, I3 any](
    f func(i1 I1, i2 I2, i3 I3) TailRec,
    i1 I1, i2 I2, i3 I3) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 4
func TailCall4[
    I1 any, I2 any, I3 any, I4 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 5
func TailCall5[
    I1 any, I2 any, I3 any, I4 any, I5 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 6
func TailCall6[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 7
func TailCall7[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 8
func TailCall8[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 9
func TailCall9[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 10
func TailCall10[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 11
func TailCall11[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 12
func TailCall12[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 13
func TailCall13[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 14
func TailCall14[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 15
func TailCall15[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 16
func TailCall16[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 17
func TailCall17[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 18
func TailCall18[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 19
func TailCall19[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 20
func TailCall20[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 21
func TailCall21[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20, i21)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 22
func TailCall22[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20, i21, i22)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 23
func TailCall23[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20, i21, i22, i23)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

// TailCall helper to call a function whose arity is 24
func TailCall24[
    I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any, I24 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) TailRec,
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) TailRec {
    var g= func () TailRec {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19, i20, i21, i22, i23, i24)
    }
    return TailRec {
        next: g,
        done: false,
    }
}

