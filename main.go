package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"math"
	"net/http"
	"strings"

	receipt "receipt/receipt"
)

func main() {
	// Create service instance.
	service := &receiptsService{
		receipts: map[string]receiptWithPoints{},
	}

	// flag -debug={debug,info,warn,error}
	debugLevel := flag.String("debug", "none", "Specify debugging level. 'debug','info','warn','error','none'")
	// flag -port=PORT
	port := flag.Int("port", 8080, "Listen to specified port")

	flag.Parse()

	// what level for logging
	if debugLevel != nil {
		log.SetFlags(log.Ltime | log.Lshortfile)
		switch strings.ToLower(*debugLevel) {
		case "debug":
			slog.SetLogLoggerLevel(slog.LevelDebug)
		case "info":
			slog.SetLogLoggerLevel(slog.LevelInfo)
		case "warn":
			slog.SetLogLoggerLevel(slog.LevelWarn)
		case "error":
			slog.SetLogLoggerLevel(slog.LevelError)
		case "none":
			// specify a level not used to hackishly disable logging.
			slog.SetLogLoggerLevel(math.MaxInt)
		default:
			panic("Can't parse debug level: " + (*debugLevel))
		}
	}

	// Create generated server.
	srv, err := receipt.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Receipt Server starting on :", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), srv); err != nil {
		log.Fatal(err)
	}
}
