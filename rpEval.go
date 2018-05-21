package main

import (
    "math/rand"
    g"github.com/peterdouglas/giota"
    "fmt"
    "log"
    "flag"
    "github.com/pkg/profile"
)

var (
    seed            g.Trytes
)

var  trs = []g.Transfer{
    {
        Address: "BXHANKTHPJUPUVZOLJPZPQLDZPWVSBPGLMLSOYFZM9RSHVZRRBZJZJDZYTNRHXBVMQKFT9DVKVNDPCGC9ZXXTZCTMB",
        Value:   1500000,
        Tag:     "RPROOF",
    },
}

func main() {
    // This code was token from the godocs at https://godoc.org/github.com/pkg/profile
    // use the flags package to selectively enable profiling.
    mode := flag.String("profile.mode", "", "enable profiling mode, one of [cpu, mem, mutex, block]")
    flag.Parse()
    switch *mode {
    case "cpu":
        defer profile.Start(profile.CPUProfile).Stop()
    case "mem":
        defer profile.Start(profile.MemProfile).Stop()
    case "mutex":
        defer profile.Start(profile.MutexProfile).Stop()
    case "block":
        defer profile.Start(profile.BlockProfile).Stop()
    default:
        // do nothing
    }

    ts := "CLBHL9DOQXUHBWORNBHNPUB9JQUHYLLXXCJQRJVRJXYHAAISJPTDA9ZFVLPPNAHLDNMDDMGYXEDVROMQV"
    s, err := g.ToTrytes(ts)
    if err != nil {
        log.Fatal(err)
    } else {
        seed = s
    }

    var exBundle g.Bundle

    for i := 0; i < 100; i++ {
        api := g.NewAPI("http://localhost:14265", nil)
        randVal := rand.Int63n(1279530283277761)
        trs[0].Value = randVal
        bdl, err := g.PrepareTransfers(api, seed, trs, nil, "")

        if err != nil {
            log.Fatal(err)
        }

        if i == 0 {
            exBundle = bdl
        }
        if err = bdl.IsValid(); err != nil {
            log.Fatal(err)
        }
    }

    tryteLen := 0
    for _, tran := range exBundle {
        tryteLen += len(tran.Trytes())
    }

    fmt.Printf("The length of a bundle in Trytes is %v\n", tryteLen)
}