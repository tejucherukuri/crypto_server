# crypto_server
Microservice API that uses the crypto API endpoint: https://api.hitbtc.com/#development-guide and returns the real-time crypto prices.


**Improvisations: ** 

I feel like i could have do more with the code if i had more time to work on it. Below, are the one's possibly i could think can be made to make the code look more organized and reusable.

1. Can split the endpoints into a different folders - routers 
2. Having the Crypto ( or Currency struct ) defined into a different folder - models
3. Having a separate folder to have the handlers code . 
4. I would have a separate reusable for passing the endpoint and getting the response instead of duplicating it in each handler.
5. Log and handle the errors for easy debugging.
6. Write a unit test for both Handlers , that sends a GET request to the handler function and checks that it returns a status code of 200.

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

go build ./crypto_server




