<h1 align="center">Book CRUD backend using ECHO</h1>

---
<p align="center"><b>This repository contains the fully organized and structured version of the <a href=#>GO Bootcamp</a> final project. The following is the introduction and a guide on how to navigate and start the backend.</b></p>

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

---

## Clone the project
```
git clone https://github.com/Raisul191491/Book_CRUD_Echo.git
cd Book_CRUD_Echo
```

## Configure database
- Navigate to app.env file.
- Change the values of environment variables according to your mysql profile. For other databases follow 
- Manually create the database that you want to use.
- Make sure the port you entered is free.

## Start backend
```
go mod tidy
```

- #### Using go run
    ```
    go run main.go
    ```
- #### Using nodemon
    
    Install nodemon

        npm install -g nodemon

    For easy launch, create a Makefile.

        run:
            nodemon -x go run main.go --signal SIGKILL -e go --verbose


    To run the application

        make run

## Functionalities and Target

 * *Create*, *Read*, *Update* & *Delete* books.
 * Provide a basic understanding of how **Echo** works. 
 * Application of **GORM**.
 * Provide understanding about backend **project structuring**.
 * Implementation knowledge of **Interface binding**. (*Repositories and Services*)

 ## Where and What?
 * `config` loads the environment variables.
 * `connection` contains the database connection and migration functions.
 * `consts` contains commonly used messages.
 * `containers` contains serve function that initializes the project.
 * `controllers` functions creates responses based on requests to the endpoints.
 * `domain` contains the repository and service interfaces. 
 * `models` holds the database schema.
 * `repositories` contains database methods.
 * `routes` contains the endpoints/paths.
 * `services` contains service methods.
 * `types` holds the different structures used throughout the project and validation.

 