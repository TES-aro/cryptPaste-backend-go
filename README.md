## cryptPaste backend  
simple RESTful API for crypPaste writen in Go.  
It is written to work with mySQL for now.  
***no encryption is done by the backend***  

### using  
environmental variable MSQL_URL is used to connect to the mySQL database.  
environmental variable MSQL_DB_NAME is used to decide the database.  
"pasteCrypt" is the default database name which is used if env var is not provided  
CRYPT_PORT is used to determine which port the server uses. 1337 is default.  

> export CRYPT_PORT="{portNumber}"  
> export MSQL_DB_NAME="{dbName}"  
> export MSQL_URL="username:password@tcp(address)"  

new entries are created with http POST at /api.  
entries are expected to be JSON with language and text fields.  

entries can be fetched with http GET from /api/{id}

