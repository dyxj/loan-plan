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
	"startDate": "2018-01-01T00:00:01Z"
}
```
Example Output:
```
[
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-01-01T00:00:01Z",
        "initialOutstandingPrincipal": 5000,
        "interest": 20.83,
        "principal": 198.53,
        "remainingOutstandingPrincipal": 4801.47
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-02-01T00:00:01Z",
        "initialOutstandingPrincipal": 4801.47,
        "interest": 20.01,
        "principal": 199.35,
        "remainingOutstandingPrincipal": 4602.12
    },
    ...
    {
        "borrowerPaymentAmount": 219.28,
        "date": "2019-12-01T00:00:01Z",
        "initialOutstandingPrincipal": 218.37,
        "interest": 0.91,
        "principal": 218.37,
        "remainingOutstandingPrincipal": 0
    }
]
```

## Documentation

### Test
To test the entire package run in the local-plan folder: `go test ./...`