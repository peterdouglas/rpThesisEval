package main

import (
    "testing"
    g "github.com/peterdouglas/giota"
    "math/rand"
    "log"
)

const ENDPOINT  = "http://localhost:14265"

func init()  {
    ts := "CLBHL9DOQXUHBWORNBHNPUB9JQUHYLLXXCJQRJVRJXYHAAISJPTDA9ZFVLPPNAHLDNMDDMGYXEDVROMQV"
    s, err := g.ToTrytes(ts)
    if err != nil {
        log.Fatal(err)
    } else {
        seed = s
    }
}

func benchmarkRPBundle(i int, api *g.API) (g.Bundle) {

    randVal := rand.Int63n(1279530283277761)
    trs[0].Value = randVal
    bdl, err := g.PrepareTransfers(api, seed, trs, nil, "")

    if err != nil {
        log.Fatal(err)
    }

    if i == 0 {
        return bdl
    } else {
        return nil
    }
}


func BenchmarkBundle(b *testing.B) {
    api := g.NewAPI(ENDPOINT, nil)
    var err error
    b.ReportAllocs()
    if err != nil {
        b.Errorf("There was an error initialising the test: %s", err)
    }

    var exBundle g.Bundle

    for i := 0; i < b.N; i++ {
       tempBdl := benchmarkRPBundle(i, api)
       if tempBdl != nil {
           exBundle = tempBdl
       }
    }

    tryteLen := 0
    for _, tran := range exBundle {
        tryteLen += len(tran.Trytes())
    }
    b.Logf("The length of a bundle in Trytes is %v\n", tryteLen)
}

func BenchmarkBundleValid(b *testing.B) {
    b.ReportAllocs()
    var err error

    if err != nil {
        b.Errorf("There was an error initialising the test: %s", err)
    }

    api := g.NewAPI(ENDPOINT, nil)
    for i := 0; i < b.N; i++ {
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

func BenchmarkAddressing(b *testing.B) {
    seed := g.NewSeed()
    b.ReportAllocs()

    for i := 0; i < b.N; i++  {
        _, err := g.NewAddress(seed, i)
        if err != nil {
            b.Error(err)
        }
    }
}

func BenchmarkGetBalances(b *testing.B) {
    b.ReportAllocs()
    api := g.NewAPI(ENDPOINT, nil)
    ts := "CLBHL9DOQXUHBWORNBHNPUB9JQUHYLLXXCJQRJVRJXYHAAISJPTDA9ZFVLPPNAHLDNMDDMGYXEDVROMQV"
    s, err := g.ToTrytes(ts)
    if err != nil {
        b.Error(err)
    } else {
        seed = s
    }

    for i := 0; i < b.N; i++  {
        // If inputs with enough balance
        _, err := g.GetInputs(api, seed, 0, 100, 100)
        if err != nil {
            b.Error(err)
        }
    }

}
