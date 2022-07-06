package tailcall

type TailRec[T any] struct {
    next func() TailRec[T]
    result T
    done bool
}

func TailRet[T any](r T) TailRec[T] {
    return TailRec[T] {
        result: r,
        done: true,
    }
}

func TailStart[T any](tr TailRec[T]) T {
    for ;; {
        if tr.done {
            return tr.result
        } else {
            tr= tr.next()
        }
    }
}

func TailCall0[T any](f func() TailRec[T]) TailRec[T] {
    return TailRec[T] {
        next: f,
        done: false,
    }
}

func TailCall[O any, I any](f func(i I) TailRec[O], i I) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall2[
    O any, I1 any, I2 any](
    f func(i1 I1, i2 I2) TailRec[O],
    i1 I1, i2 I2) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall3[
    O any, I1 any, I2 any, I3 any](
    f func(i1 I1, i2 I2, i3 I3) TailRec[O],
    i1 I1, i2 I2, i3 I3) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall4[
    O any, I1 any, I2 any, I3 any, I4 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall5[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall6[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall7[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall8[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall9[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall10[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall11[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall12[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall13[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall14[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall15[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall16[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall17[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall18[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall19[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
            i9, i10, i11, i12, i13, i14, i15, i16,
            i17, i18, i19)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall20[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
        i9, i10, i11, i12, i13, i14, i15, i16,
        i17, i18, i19, i20)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall21[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
        i9, i10, i11, i12, i13, i14, i15, i16,
        i17, i18, i19, i20, i21)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall22[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
        i9, i10, i11, i12, i13, i14, i15, i16,
        i17, i18, i19, i20, i21, i22)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall23[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
        i9, i10, i11, i12, i13, i14, i15, i16,
        i17, i18, i19, i20, i21, i22, i23)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

func TailCall24[
    O any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any,
    I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any,
    I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any, I24 any] (
    f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) TailRec[O],
    i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8,
    i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16,
    i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) TailRec[O] {
    var g= func () TailRec[O] {
        return f(i1, i2, i3, i4, i5, i6, i7, i8,
        i9, i10, i11, i12, i13, i14, i15, i16,
        i17, i18, i19, i20, i21, i22, i23, i24)
    }
    return TailRec[O] {
        next: g,
        done: false,
    }
}

