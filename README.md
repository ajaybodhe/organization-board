# organization-board
REST API FOR ORGANIZATION TEAM STRUCTURE MANAGMENT

## LOGIN CURL REQUEST
curl -v -d '{"email": "personia@org.com", "password": "personia"}' -H 'Content-Type: application/json' http://localhost:9090/api/v1/login?
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> POST /api/v1/login? HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.64.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 53
> 
* upload completely sent off: 53 out of 53 bytes
< HTTP/1.1 200 OK
< Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoxLCJlbWFpbCI6InBlcnNvbmlhQG9yZy5jb20ifX0._w2Ksdm7sOxiAHJw61ZR1X7aldRLa564wK-9e5O9f-c
< Content-Type: application/json
< Date: Sat, 12 Sep 2020 17:26:42 GMT
< Content-Length: 57
< 
* Connection #0 to host localhost left intact
{"status":200,"data":{"id":1,"email":"personia@org.com"}}* Closing connection 0
