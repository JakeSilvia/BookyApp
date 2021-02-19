# Booky -  http://booky-env-1.us-east-1.elasticbeanstalk.com/

#### To Run Locally
```
npm i
npm run dev
go run main.go
```

#### Compiles and minifies for production
```
npm run prod
```


<br>

#### Notes: This project is laid out to make unit testing reach 100% coverage with relative ease. Can be containerized easily with a Dockerfile to deploy to EKS/ECS.


## Infra Layout:
Deployed on Elastic Beanstalk inside of a VPC with a security group.
A Classic tcp load-balancer to handle websocket connections.
MongoDB for ease of setup and deployment but a better solution may be redis or dynamodb depending on use case.


## Project Layout:

- `server`
    - `data_gateways` - Data layer implementation. Each folder is a separate datastore implementation based on the requirements from `BookInterface`
    - `handlers` - Endpoint layer logic. Each file is a separate endpoint that only holds logic for parsing request and initializing/running the interactor. Every Handler contains the data-gateways that are required to run the endpoint. This allows swapping datastores from main without affecting the endpoint.
        - `booky_context` - holds possible middleware and helper functions for all endpoints. 
    - `pkg` - Business layer logic
        - `book_entities` - Holds the definitions for book objects and Query/Mutation objects as well as object specific functions (ie. ToList())
        - `book_interactors` - Implementation of business logic for each action. Each interactor is responsible for handling validation of both input and consistency, as well as applying Data logic without knowledge of datastore implementation via BookInterface.
        - `book_interfaces` - Description of requirements for a datastore implementation candidate. By using an interface to describe all datastore method requirements for the `book_interactors`, we can replace the datastore and be confident the interactors will not change. This also allows for a <b><i>Mock</i></b> datastore for unit testing.  
    - `settings` - settings for the application (ie. port to listen, db connection string) 

- `src` - Frontend
