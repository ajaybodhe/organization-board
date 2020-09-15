# Personio : organization-board
REST APIs FOR ORGANIZATION's TEAM STRUCTURE/HIERARCHY MANAGMENT

## Getting Started

### Requirements :
```
* GO 1.14
```
Install GO from [here](https://golang.org/dl/) <br/>
All other required dependancies are already part of the project.

### Run Test Cases
```
./tests.sh
```
### Run The project
Script builds GO binary and runs the binary.
```
./run.sh
```
### CURL Requests For Testing The APP
* Login and Get JWT token
```
curl -v -d '{"email": "personia@personio.com", "password": "personia"}' -H 'Content-Type: application/json' http://localhost:9090/api/v1/login | json_pp
```
* POST valid Employee to Manager Mapping
```
curl -d '{"Petre": "Nick","Barbara": "Nick","Nick": "Sophie","Sophie":"Jonas"}' -H 'Content-Type: application/json' -H 'Authorization: BEARER <strong><em>TOKEN_From_First_Curl</strong></em>' http://localhost:9090/api/v1/emplymgrmap | json_pp
```
* POST Employee to Manager Mapping having loop
```
curl -d '{"Peter": "Nick","Barbara": "Nick","Nick": "Sophie","Sophie": "Peter"}' -H 'Content-Type: application/json' -H 'Authorization: BEARER <strong><em>TOKEN_From_First_Curl</strong></em>' http://localhost:9090/api/v1/emplymgrmap? | json_pp
```
* POST Employee to Manager Mapping having Multiple Root Employees
```
curl -d '{"Peter": "Nick","Barbara": "Nick","Nick": "Sophie", "John": "Johnie"}' -H 'Content-Type: application/json' -H 'Authorization: BEARER <strong><em>TOKEN_From_First_Curl</strong></em>' http://localhost:9090/api/v1/emplymgrmap | json_pp
```
* GET complete Employee to Manager mapping
```
curl -H 'Content-Type: application/json' -H 'Authorization: BEARER <strong><em>TOKEN_From_First_Curl</strong></em>' http://localhost:9090/api/v1/emplymgrmap? | json_pp
```
* GET Supervisor Info of an Employee
```
curl -H 'Content-Type: application/json' -H 'Authorization: BEARER <strong><em>TOKEN_From_First_Curl</strong></em>' http://localhost:9090/api/v1/emplymgrmap/Nick?supervisor=true
```

## Application Design

### High Level Design :
Structuring of code is inspired from [MindInventory](https://github.com/Mindinventory/Golang-Project-Structure)
<br/>JWT Tokens are used for authentication.<br/>Write-through cache added on top of DB to expedite GET APIs.
#### Resource :
Resources such as config files, sqlite db files etc
#### Config :
Config key-value definitions for the Project, derived from config files stores in Resource section.
#### Constants :
Project constants such as SQL Query definitions, Authentication constants etc
#### Models :
Domain data model for the project, all POJO go here.
#### Db :
Exposes DB interface, actual concrete implementation is hidden by this interface.
#### Repository :
CRUD layer on Models using DB interace.
#### cache :
In Process caching implementation on DB layer to expedite GET queries.
#### handlers
Combines classic functionality of controllers and service layers into single layer for simplicity. First point of entry for each API request, encaspsulates service logic.
#### Apihelpers
Helper functions to generate the HTTP responses from domain models.



## Assumptions
* For task No 1, We only support POST semantics for hierarchies. So on each new POST request, we do overwrite the hierarchies in sqlite.
* Response for the task No 3, response for retrieving supervisor and super-supervisor is :
```
{
  "supervisor" : "Nick",
  "supervisor_of_supervisor" : "Sophie"
}
```
* For task No 4, we have created a dummy user in DB with email and password credentials.<br/> To use the APIs, first get the JWT token using Login Curl command. All other API calls are authenticated using this JWT token.

## Improvement Ideas
* Caching impllementation is very basic as of now, can be extended to use Bigcache/Redis. Also cache includes entire Employee to Manger Map, can be extended to cache Hierarchical response and Supervisor info of employees.
* Add Context as first param of each function which enables common parameter passing and deadlines.
* BDD frameworks [Gingko](https://onsi.github.io/ginkgo/), [Gomega](https://onsi.github.io/gomega/) can be used for more expressive test cases.
* Scale/Perf run the app with [pprof](https://blog.golang.org/pprof) to find out any cpu, memory, performance bottlenecks.
* Improve metric, tracing and logging of app. Use [zap](https://github.com/uber-go/zap)
* All the errors can be numbered to build full fledged documentation around it.
* Have a postman collection for all the supported API calls.