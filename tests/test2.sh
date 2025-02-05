# should be worth 109 points
curl -i -X "POST" -H "Content-Type: application/json" --data \
"{
  \"retailer\": \"M&M Corner Market\",
  \"purchaseDate\": \"2022-03-20\",
  \"purchaseTime\": \"14:33\",
  \"items\": [
    {
      \"shortDescription\": \"Gatorade\",
      \"price\": \"2.25\"
    },{
      \"shortDescription\": \"Gatorade\",
      \"price\": \"2.25\"
    },{
      \"shortDescription\": \"Gatorade\",
      \"price\": \"2.25\"
    },{
      \"shortDescription\": \"Gatorade\",
      \"price\": \"2.25\"
    }
  ],
  \"total\": \"9.00\"
}" http://localhost:8080/receipts/process
