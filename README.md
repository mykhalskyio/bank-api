# Evo Fintech

## Live Demo
### swagger
https://bank.mykhalskyio.xyz/swagger/index.html

### api 
https://bank.mykhalskyio.xyz/api/

---

## Requirements
* Git
* Docker
* golang-migrate


## How to run with MakeFile

Clone repository

    $ git clone https://github.com/mykhalskyio/image-api.git

Build and Run

    $ make build

Run migrate up

    $ make migration-up



## Endpoints
### client
* POST   - /api/client/create
* GET    - /api/client/list
* GET    - /api/client/

### account 
* POST   - /api/account/create
* GET    - /api/account/list
* GET    - /api/account/

### transaction
* POST   - /api/transaction/create
* GET    - /api/transaction/list
* GET    - /api/transaction/

### swagger
* GET - /swagger/index.html
