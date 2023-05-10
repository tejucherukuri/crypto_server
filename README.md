# crypto_server
Microservice API that uses the crypto API endpoint: https://api.hitbtc.com/#development-guide and returns the real-time crypto prices.


**API Routes :**

1. **Route API : **http://localhost:8080/currency/all
Route Descripton: API Route to fetch all the currencies.

**Response:** [
  {
    "id": "DBIX",
    "fullName": "DBIX",
    "crypto": true,
    "payinEnabled": false,
    "payoutEnabled": false
  },
  {
    "id": "UNO",
    "fullName": "Unobtanium",
    "crypto": true,
    "payinEnabled": false,
    "payoutEnabled": true
  },
  {
    "id": "AEON",
    "fullName": "Aeon",
    "crypto": true,
    "payinEnabled": true,
    "payoutEnabled": true
  },
  {
    "id": "PRG",
    "fullName": "PRG",
    "crypto": true,
    "payinEnabled": false,
    "payoutEnabled": true
  },
  {
    "id": "UPT",
    "fullName": "UPT",
    "crypto": true,
    "payinEnabled": false,
    "payoutEnabled": true
  },
  {
    "id": "TV",
    "fullName": "Tokenville",
    "crypto": true,
    "payinEnabled": true,
    "payoutEnabled": true
  },
  {
    "id": "VOW",
    "fullName": "Vow token",
    "crypto": true,
    "payinEnabled": true,
    "payoutEnabled": true
  },
  {
    "id": "BIZZ",
    "fullName": "BIZZCOIN",
    "crypto": true,
    "payinEnabled": true,
    "payoutEnabled": true
  },
  {
    "id": "LTC",
    "fullName": "Litecoin",
    "crypto": true,
    "payinEnabled": true,
    "payoutEnabled": true
  },
  {
    "id": "FLUX",
    "fullName": "Flux",
    "crypto": true,
    "payinEnabled": false,
    "payoutEnabled": false
  },
  {
    "id": "OGN",
    "fullName": "Origin Protocol",
    "crypto": true,
    "payinEnabled": true,
    "payoutEnabled": true
  },
  {
    "id": "KRL",
    "fullName": "Kryll",
    "crypto": true,
    "payinEnabled": true,
    "payoutEnabled": true
  },
  {
    "id": "PMA",
    "fullName": "PumaPay",
    "crypto": true,
    "payinEnabled": true,
    "payoutEnabled": true
  },
  {
    "id": "EMGO",
    "fullName": "EMGO",
    "crypto": true,
    "payinEnabled": false,
    "payoutEnabled": true
  },
  {
    "id": "MNX",
    "fullName": "MNX",
    "crypto": true,
    "payinEnabled": false,
    "payoutEnabled": false
  }
  ;
  ;
  ;
  ]

2. **Route API :** http://localhost:8080/currency/{currency}**
**Route Description:** Fetches the currency details of a single currency.
Example API:** http://localhost:8080/currency/LEVL

**Response:** {
  "id": "LEVL",
  "fullName": "Levolution",
  "crypto": true,
  "payinEnabled": true,
  "payoutEnabled": true
}


Commands :
go build
go run crypto_server.go

go build ./parking_services
