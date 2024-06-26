<p align="center">
  <a href="#description">Description</a> •  
  <a href="#how-to-run">How To Run</a> •  
  <a href="#how-to-test">How To Test</a> •  
  <a href="#features">Features</a> •  
</p>

## Description

- REST API developed in Go
- Implements Clean Archicture which allows to distribute this project in different logic layers:
  - Infraestructure: dependencies configurations
  - Adapters: dependencies implementations
  - Use cases: business logic
  - Entities: business types

## How To Run

Install dependencies:

```
go get .
```

Run:

```
go run main.go
```

## How To Test

This project seeks to program unit tests for use cases that implement business logic and for utils that support the application with abstract logic.

```
go test ./... -count=1
```

## Features 

✅: Implemented and tested <br/>
⚠️: Implemented and not tested <br/>
⚙️: In progress <br/>
📅: To implement <br/>
💡: Idea for the future <br/>


|Module|Feature|Status|
|:----|:----|:----|
|**Login**|Google OAuth2| ✅|
|**Middlewares**|Authentication with Json Web Token|✅|
||Authorization system|✅|
|**Permissions**|Create|✅|
||Add resource|✅|
||Authorize|✅|
||Remove resource|✅|
|**Users**|Update|✅|
||Get by ID|✅|
||Create|✅|
|**Addresses**|Create|✅|
||Get by ID|✅|
||Update|✅|
||Delete|✅|
||Get by user ID|✅|
<br/>

- Products: 💡
- Ecommerce: 💡
- Logistics: 💡





    
    


