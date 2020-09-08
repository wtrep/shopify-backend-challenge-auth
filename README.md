# Shopify Backend Challenge - Auth Microservice
This is a simple GO microservice that allows the creation and authentication of users through a REST API. The following actions are supported :
 - Create a new user
 - Get a temporary authentication token that authenticates the user with other microservices

For a detailed documentation on how to query the API, please visit the [SwaggerHup API Page](https://app.swaggerhub.com/apis-docs/wtrep/shopify-images-repo/1.0.0).

## Details about the microservice
### File structure
The main logic of the microservice can be found in the auth package. The common package is shared between this microservice and the [image microservice](https://github.com/wtrep/shopify-backend-challenge-image).

### Authentification
The authentification is handled through username and password. The password is kept hashed in the database using the Bcrypt cryptographic algorithm.

### Database
The microservice needs to have access to a MySQL database located at `localhost:3306`. You can find the Terraform code for a GCP Cloud SQL instance in the [main repository](https://github.com/wtrep/shopify-backend-challenge/tree/master/terraform/cloud_sql). For local testing you can use the [mysql docker image](https://hub.docker.com/_/mysql). To allow access to Cloud SQL in GKE, you need to use the Cloud SQL sidecar proxy as shown in the [main
repository](https://github.com/wtrep/shopify-backend-challenge/blob/master/kubernetes/image-microservice-deployment.yml).


### Docker Image and Kubernetes
The microservice is packaged into a Docker image to allow deployment into a Kubernetes Cluster. You can also download the built image directly from [Docker Hub](https://hub.docker.com/r/wtrep/shopify-backend-challenge-image)

## Environment variables
The following environment variables need to be set for the microservice to work :
| Environment variable           | Description                                                                                                                            |
| -------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------:|
| DB_USERNAME                    | Username to access the MySQL DB                                                                                                        |
| DB_PASSWORD                    | Password to access the MySQL DB                                                                                                        |
| DB_NAME                        | Name of the MySQL database                                                                                                             |
| JWT_KEY                        | Private key to verify JWT Tokens. Must be the same as the [auth microservice](https://github.com/wtrep/shopify-backend-challenge-auth) |

## Build and run
To build the microservice : 
```
go build main.go
```

To run the microservice, you must configure the required environment variables first. Then run :
```
go run main.go
```

## Build the Docker image
The provided Dockerfile allow the microservice to be packaged into a container. To build the Docker image:
```
docker build .
```

## List of possible improvements 
 * Automated tests
 * Use context to handle timeout on each SQL request
 * Implement OAuth2
 * Allow the deletion of users
 * Bake the authentication of user inside an API Gateway instead of another microservice
 * Implement 2FA
 * CI/CD Pipeline that builds and upload to Docker Hub a new Docker image at each merge to the master branch