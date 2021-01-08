# SIMPLE USER PROFILE w@ GOLANG
 
# Note

> Uses GIN as a Framework
> Uses Gorm as ORM
> Uses Mysql as RDBMS
> Uses Go Modules as Package Manager
> Uses Viper as Configuration
> Uses JWT as a auth
> Uses Gomon as auto rebuilder, or you may use GO RUN 

### Usage

  - Unzip file
  - Go to Directory [simple-user-profile] at this case
  - Running script with "Go Mod Vendor" to include package manager
  - Create Database Mysql called "simple_user_profile"
  - Running App with gomon app.go or you may use go run app.go
 
### Postman (this app used postman / not browser)
   
    - # USER With Method GET
    - In your Headers please add 'Key'=>'Value'
    - 'X-Lemo-Token' => '853D6Ca929939037112080a3aBb24155b74B731F'
    - 'App-Key' => 'LM2021XYZ'
    - Load Hostname:Port/api/v1/users {Localhost:8080/v1/api/users}
    - /v1/api/users => Retrive all users, not paginated for now
    - /v1/api/users/:ID => Retrive user with ID .eg* 1
    - 

### Login / Registration
    - # Method POST
    - [Host:Port]/api/v1/registration
    - Payload with raw json type {
	    "username":"lemonilo",
	    "email":"user@lemonilo.su",
	    "password":"lemonilo",
	    "address":"Jakarta"
        }
        
    - /v1/api/login
    -  Payload with raw json type {
	    "username":"lemonilo",
	    "password":"lemonilo",
        }
