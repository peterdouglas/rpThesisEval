package main

import (
    "testing"
    g "github.com/peterdouglas/giota"
    "math/rand"
    "log"
)

var (
    seed             g.Trytes
)

func init()  {
    ts := "CLBHL9DOQXUHBWORNBHNPUB9JQUHYLLXXCJQRJVRJXYHAAISJPTDA9ZFVLPPNAHLDNMDDMGYXEDVROMQV"
    s, err := g.ToTrytes(ts)
    if err != nil {
        log.Fatal(err)
    } else {
        seed = s
    }
}

var  trs = []g.Transfer{
   {
        Address: "BXHANKTHPJUPUVZOLJPZPQLDZPWVSBPGLMLSOYFZM9RSHVZRRBZJZJDZYTNRHXBVMQKFT9DVKVNDPCGC9ZXXTZCTMB",
        Value:   1500000,
        Tag:     "RPROOF",
    },
}


func BenchmarkRPBundle(b *testing.B) {
    var err error
    b.ReportAllocs()
    if err != nil {
        b.Errorf("There was an error initialising the test: %s", err)
    }

    var exBundle g.Bundle

    for i := 0; i < 10; i++ {
        api := g.NewAPI("http://localhost:14265", nil)
        randVal := rand.Int63n(1279530283277761)
        trs[0].Value = randVal
        bdl, err := g.PrepareTransfers(api, seed, trs, nil, "")

        if err != nil {
            b.Error(err)
        }

        if i == 0 {
            exBundle = bdl
        }
    }

    tryteLen := 0
    for _, tran := range exBundle {
        tryteLen += len(tran.Trytes())
    }

    b.Logf("The length of a bundle in Trytes is %v\n", tryteLen)
}

func BenchmarkRPBundleValid(b *testing.B) {

    var err error

    if err != nil {
        b.Errorf("There was an error initialising the test: %s", err)
    }

    for i := 0; i < 10; i++ {
        api := g.NewAPI("http://localhost:14265", nil)
        randVal := rand.Int63n(1279530283277761)
        trs[0].Value = randVal
        bdl, err := g.PrepareTransfers(api, seed, trs, nil, "")

        if err != nil {
            b.Error(err)
        }

        if err = bdl.IsValid(); err != nil {
            b.Error(err)
        }
    }
}

