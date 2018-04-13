# README

Generate Unique ID.

```
go get github.com/newm4n/go-jenius-unique
```

## HOW ???

You should read the test class  `idgen_test.go` on how to use the library.
TL;DR :
```
import (
	"jeniusunique"
)
...
unique24CharString := jeniusunique.GetUniqueGenInstance().NewXReferenceNo(24)
...
```
`NewXReferenceNo` accept an number parameter which is the length of string to be generated.

## HOW YOU DO IT ?

```
nn = last known nano second time stamp in int64
i = mac-address in int64

func GetUniqueGenInstance(l) string
    n = nano second timestamp in int64
    c = 0
    if n == nn {
        c++
        unique = n xor i xor c
    } else {
        c = 0
        unique = n xor i xor 0
    }
    nn = n
    if len(unique) < l {
        append unique with random char until it has len == l
    }
}
```