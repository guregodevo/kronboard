gomiranal
=========

To start REST server, run :  
go run src/main.go

To run backdoor command : 
go run -v src/backdoor.go -db h

where h is the short code for Help command. 

Authentication API:
=================== 

Create User 
============

POST http://localhost:3000/authenticate?username="christophe"&password="flynow001"&company="ACME"
{
	token: ""
}
Returns :
200 The user has been created
{
token: "43318c0ba635d070013fcda4f4270856430d7851747b257e44edab958a3c9ca9"
}
406 Conflict due to an existing user.

500 A technical error occurred


Authenticate User 
===================

GET http://localhost:3000/authenticate?username="christophe"&password="flynow001"
{
	token: ""
}

Returns :
200 The user has been authenticated
409 Unauthorized The Credential is wrong or the user is not authorized.
500 A technical error occurred

Dashboard API:
=================== 
GET http://localhost:3000/dashboards?id=123
[
{"col":0,"row":0,"sizeX":3,"sizeY":2,"type":"line"},
{"col":3,"row":1,"sizeX":2,"sizeY":1,"type":"bar"}
]
Returns :
200 The dashboard has been found
400 The dashboard has not been found
500 A technical error occurred

Chart API:
=================== 
GET http://localhost:3000/charts?id=123
{
Id: 1
Interval: 2
Type: "line"
Created: "2014-04-05T16:49:48.215071Z"
}
Returns :
200 The chart has been found
400 The chart has not been found
500 A technical error occurred
