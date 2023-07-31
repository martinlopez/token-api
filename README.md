
  # Tokens API 🚀  
  Import your existing Readme using the import button at the bottom, 
  or create a new Readme from scratch by typing in the editor.  
  
  ## Scaffolding 🚀  
  - cmd
  - db
  - internal
  - 
  

  ## Prebuilt Components/Templates 🔥  
  You can checkout prebuilt components and templates by clicking on the 'Add Section' button or menu icon
  on the top left corner of the navbar.
      
  ## Save Readme ✨  
  Once you're done, click on the save button to download and save your ReadMe!
  
# Tokens API  

# Scaffolding  

- cmd (executable files)
- db (contains migrations)
- internal (business logic)
  - [service_name] (each service will have a domain and usecase)
    - domain
    - usecase
- pkg  (code that can be shared)

Each service use case should implement a CRUD interface. This interface is known by a generic handler keeping simple creating new services. If a service has to do an specific action the usecase could implement other interface, but will be necessary create new handler.

# Components

Currently this is using mysql DB connection which credentials we get from AWS secrets manager.

To run migrations:

~~~bash  
 export DBPASSWORD=$(aws secretsmanager get-secret-value --secret-id dev/core --region us-east-1 --output json | jq -r -S '.SecretString| fromjson| .password')
~~~
~~~bash  
make execute-migration
~~~



### Token Job
Based on csv file, it get token information from https://blockpartyplatform.mypinata.cloud/ipfs/:cid and store it in database.

Run locally using docker:
~~~bash  
  make run-token-job
~~~
or without docker:
~~~bash  
  go run cmd/function/sync_tokens/main.go
~~~

### Token API

Only read mode. It returns information about tokens.

Run locally using docker:
~~~bash  
  make run-token-api
~~~
or without docker:
~~~bash  
  go run cmd/http/main.go        
~~~

#### Endpoints

~~~bash  
  curl "http://localhost:8080/tokens"
~~~

~~~bash  
  curl "http://localhost:8080/tokens/:cid
~~~
 

## Tech Stack  

 Golang
