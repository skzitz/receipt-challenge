package main

import (
    "bytes"
    "context"
    "encoding/json"
    "flag"
    "os"
    "fmt"
    "errors"

    "github.com/fatih/color"

    receipt "receipt/receipt"
)

// examples provided
// {
//     "retailer": "Target",
//     "purchaseDate": "2022-01-02",
//     "purchaseTime": "13:13",
//     "total": "1.25",
//     "items": [
//         {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
//     ]
// }// simple-receipt.json
// morning-reeipt.json
// {
//     "retailer": "Walgreens",
//     "purchaseDate": "2022-01-02",
//     "purchaseTime": "08:13",
//     "total": "2.65",
//     "items": [
//         {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
//         {"shortDescription": "Dasani", "price": "1.40"}
//     ]
// }


func run(ctx context.Context) error {
    var arg struct {
        BaseURL string
        ID      int64
    }
    flag.StringVar(&arg.BaseURL, "url", "http://localhost:8080", "target server url")
    flag.Int64Var(&arg.ID, "id", 1, "pet id to request")
    flag.Parse()

    client, err := receipt.NewClient(arg.BaseURL)
    if err != nil {
        return fmt.Errorf("create client: %w", err)
    }

    res, err := client.GetPetById(ctx, petstore.GetPetByIdParams{
        PetId: arg.ID,
    })
    if err != nil {
        return fmt.Errorf("get pet %d: %w", arg.ID, err)
    }

    switch p := res.(type) {
    case *petstore.Pet:
        data, err := p.MarshalJSON()
        if err != nil {
            return err
        }
        var out bytes.Buffer
        if err := json.Indent(&out, data, "", "  "); err != nil {
            return err
        }
        color.New(color.FgGreen).Println(out.String())
    case *petstore.GetPetByIdNotFound:
        return errors.New("not found")
    }

    return nil
}

func main() {
    if err := run(context.Background()); err != nil {
        color.New(color.FgRed).Printf("%+v\n", err)
        os.Exit(2)
    }
}
