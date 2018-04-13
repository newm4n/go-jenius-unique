package jeniusunique

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
	"math"
)

const (
	RANDCHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type UniqueGen struct {
	Nano  int64
	Count int64
	Iface int64
}

var (
	UniqueGenInstance *UniqueGen
	UniqueGenLock     sync.Mutex
)

func GetUniqueGenInstance() *UniqueGen {
	if UniqueGenInstance == nil {
		var ifacen int64 = 0
		interfaces, err := net.Interfaces()
		if err == nil {
			for _, i := range interfaces {
				if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
					for _, b := range i.HardwareAddr {
						ifacen = ifacen | int64(b)
						ifacen = ifacen << 8
					}
				}
			}
		} else {
			rand.Seed(time.Now().UnixNano())
			ifacen = rand.Int63()
		}
		UniqueGenInstance = &UniqueGen{
			Nano:  time.Now().UnixNano() / 1000,
			Iface: int64(math.Abs(float64(ifacen))),
			Count: 0,
		}
	}
	return UniqueGenInstance
}

func (ug *UniqueGen) NewXReferenceNo(length int) string {
	UniqueGenLock.Lock()
	defer UniqueGenLock.Unlock()
	nnano := time.Now().UnixNano() / 1000
	var strRnd string
	ug.Count++
	if ug.Nano != nnano {
		ug.Count = 0
	}
	ug.Nano = nnano
	strRnd = fmt.Sprintf("%x%x", ug.Nano^ug.Iface,ug.Count)[1:]

	if len(strRnd) > length {
		strRnd = strRnd[len(strRnd)-length:]
	} else if len(strRnd) < length {
		var buffer bytes.Buffer
		buffer.WriteString(strRnd)
		for buffer.Len() < length {
			offset := rand.Intn(len(RANDCHARSET))
			buffer.WriteString(RANDCHARSET[offset : offset+1])
		}
		strRnd = buffer.String()
	}
	return strRnd
}
