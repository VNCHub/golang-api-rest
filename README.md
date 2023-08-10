# golang-api-rest
Repository with rest API in golang, using gorilla mux and gorilla handlers for dealing with handlers. Used docker for database management. Implemented CRUD through the GORM ORM.

## Dependencies

To carry out the project, it is necessary to have previously requested:
1. The Golang language - version 1.20+. If not installed, check [this page](https://go.dev/doc/install).
2. The Docker container manager. If not installed, check [this page](https://www.docker.com/products/docker-desktop/).

## Project Installation

In the terminal, in any directory of your choice, download the project using the command

```
git clone https://github.com/VNCHub/golang-api-rest.git

```

Access the downloaded directory folder using the command

```
cd golang-api-rest/

```

INSIDE the downloaded directory, to download the containers and start running the REST API server, execute the commands

```
docker-compose up -d
go run main.go

```
## Execution
If all the previous steps were successful, the database server should be live on port 54321, and the REST API server should be live on port 8000.

The project by default has some initialized data in the "characters" table

**1. To view all data in JSON format**, perform a GET request to the address <http://localhost:8000/api/characters>

**2. To view data in JSON format**, perform a GET request to the address <http://localhost:8000/api/characters/{id}> where "id" is a number from 1-10.

**3. To create data**, perform a POST request to the address <http://localhost:8000/api/characters> with the body of the request containing a JSON file, informing the "name" and "history" according to the model:

```
{
     "name": "Ron Weasley",
     "history": "A great friend."
}
```
**4. To delete data**, perform a DELETE request to the address <http://localhost:8000/api/characters/{id}>, where "id" is the number of the record you want to delete.

**5. To edit data**, make a PUT request to the address <http://localhost:8000/api/characters/{id}>, where "id" is the number of the record you want to update, containing a JSON file in the request parameters of "Topic 3 - Create Data".

## Notes:
1. All steps were performed on Linux version Ubuntu 22.04.02 LTS, compatibility with other operating systems and versions is not guaranteed.
