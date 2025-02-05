package main

import (
    "log"
    "log/slog"
    "net/http"
    "flag"
    "strings"

    receipt "receipt/receipt"
)

func main() {
    // Create service instance.
    service := &receiptsService{
        //    receipts: map[string]receipt.receipt{},
        receipts: map[string]receiptWithPoints{},
    }
    // pass -debug={debug,info,warn,error}
    debugLevel := flag.String("debug","info","Specify debugging level. 'debug','info','warn','error'")

    flag.Parse()

    // do we want to log?
    if debugLevel != nil {
        //logger := slog.New(slog.NewTextHandler(os.Stdout,nil))
        //slog.SetDefault()
        log.SetFlags( log.Ltime|log.Lshortfile )
        switch strings.ToLower(*debugLevel) {
        case "debug":
            slog.SetLogLoggerLevel(slog.LevelDebug)
        case "info":
            slog.SetLogLoggerLevel(slog.LevelInfo)
        case "warn":
            slog.SetLogLoggerLevel(slog.LevelWarn)
        case "error":
            slog.SetLogLoggerLevel(slog.LevelError)
        default:
            panic( "Can't parse debug level: " + (*debugLevel) )
        }
    }
    // Create generated server.
    srv, err := receipt.NewServer(service)
    if err != nil {
        log.Fatal(err)
    }
    if err := http.ListenAndServe(":8080", srv); err != nil {
        log.Fatal(err)
    }
}
