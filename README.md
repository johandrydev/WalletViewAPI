# WalletViewApi

## Installation

Clone te repository, go to cmd folder and run the program:

```bash
git clone https://github.com/johandrydev/WalletViewAPI.git
cd WalletViewAPI/cmd
go run wallet_view_api.go
```

To create a binary file or executable, run the following command in the cmd folder:

```bash
go build wallet_view_api.go
```

## API
Actually the project has an port const in the wallet_view_api.go file. The port is 8080. The API has the following endpoints:

### Wallet Balance
Let us to get the exchange rate of an account

```curl
	GET http://localhost:8080/walletBalance
```

#### Query Parameters
| Parameter | Type | Description |
| --- | --- | --- |
| address | string | The address of the wallet |
| currency | string | The currency to get the exchange rate balance |

#### Response
| Parameter | Type | Description |
| --- | --- | --- |
| message | string | The message of the response |
| data | object | The data of the response |

#### **Samples**

##### Success

```bash
$ curl --request POST 'http://localhost:8080/walletBalance?address=0xcfC9586Ce5d015612F69A3927178157b905AaDf0&currency=ETH' \
--header 'Content-Type: application/json'
```

```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 13 Sep 2023 21:55:31 GMT
Content-Length: 71

{
	"message": "success",
	"data": {
		"ETH": "113.15032589",
		"USDT": "0.03125716"
	}
}
```

##### Error

```bash
$ curl --request POST 'http://localhost:8080/walletBalance?currency=ETH' \
--header 'Content-Type: application/json'
```

```json
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Wed, 13 Sep 2023 21:55:31 GMT
Content-Length: 33

{
	"message": "address is required"
}
```