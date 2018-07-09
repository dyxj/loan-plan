# loan-plan
Golang Server to generate loan repayment plan.

## Requirements
Go version 1.8.3

## Run
Change directory to loan-plan, run main with `go run main.go`.  
Backend server deployed to port 8080.  

### RESTful service URLs
`POST /generate-plan`  
with json body :-
```
{
	"loanAmount": "5000",
	"nominalRate": "5.0",
	"duration": 24,
	"startDate": "2018-01-01"
}
```
Example Output:
```
[
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-01-01T00:00:00Z",
        "initialOutstandingPrincipal": 5000,
        "interest": 20.83,
        "principal": 198.53,
        "remainingOutstandingPrincipal": 4801.47
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-02-01T00:00:00Z",
        "initialOutstandingPrincipal": 4801.47,
        "interest": 20.01,
        "principal": 199.35,
        "remainingOutstandingPrincipal": 4602.12
    },
    ...
    {
        "borrowerPaymentAmount": 219.28,
        "date": "2019-12-01T00:00:00Z",
        "initialOutstandingPrincipal": 218.37,
        "interest": 0.91,
        "principal": 218.37,
        "remainingOutstandingPrincipal": 0
    }
]
```  

`POST /calc-annuity`  
with json body :-
```
{
	"loanAmount": "5000",
	"nominalRate": "5.0",
	"duration": 24
}
```
Example Output:
```
{
    "annuity": "219.36"
}
```

## Documentation

### Test
To test the entire package run in the local-plan folder: `go test ./...`

### Docker and Docker Compose  
An example is provided if you would like to use docker and docker-compose.  
In this case docker and docker-compose needs to be installed.  
To start the server run `docker-compose up` in the loan-plan directory.  
Be sure to run `docker-compose down` when you're done, to perform clean up.  