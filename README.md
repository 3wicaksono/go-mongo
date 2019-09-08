# Go Mongo
Simple API using Go & MongoDB, this is prototype was built for pre-assessment task only
  
```  
Author: Tri Wicaksono <3wickasono@gmail.com>  
```  

## Installation with Docker  
#### Install MongoDB & API
Just run simple command :
```  
$ ./run.sh  
```  

wait until all stuff get done

## Manual Installation
### System Requirement  
```  
1. Golang v1.10.3  
2. Glide Package Manager  
3. MongoDB v3.6
```  
  
## Installation guide  
#### 1. Install MongoDB
```
# please follow installation instruction from
# https://docs.mongodb.com/manual/installation
```

#### 2. install go version 1.10.3 
```
# please read this link installation guide of go  
# https://golang.org/doc/install  
```  
  
#### 3. Create directory workspace 
```
# run command below: $ mkdir $HOME/go  
$ mkdir $HOME/go/src  
$ mkdir $HOME/go/pkg  
$ mkdir $HOME/go/bin 
$ chmod -R 775 $HOME/go  
``` 

edit bash profile in $HOME/.bash_profile 

add below to new line in file .bash_profile           
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
run command :  
```
source $HOME/.bash_profile  
```  
  
#### 4. Install glide package manager 
```
#read installation guide https://github.com/Masterminds/glide  
```  
  
#### 5. Build the application 
```
# run command :  
$ cd $HOME/go/src  
$ git clone https://github.com/3wicaksono/go-mongo.git  
$ cd $HOME/go/src/go-mongo  
$ glide install  

# edit configurations/App.yaml with server environtment (db connection) 
$ nano configurations/App.yaml  
  
# build app  
$ go build  
  
# run application after build or create on supervisord 
$ ./go-mongo serve  
``` 
 
  
## API Documentation
For more fancy look visit 
[Go Mongo Postman API Documentation](https://documenter.getpostman.com/view/849676/SVmpWMTD?version=latest)

Default API host is `http://localhost:8787`  

### Comments
#### Comments - Publish  
Endpoint  `{url}:{port}/orgs/{org-name}/comments`

Method:`POST`

Header: `Content-Type: application/json`

Body:
```
{  
    "comment": "Looking to hire SE Asia's top dev talent!"  
}
```
Response
```  
{
    "message":  "success"
}
```  
#### Comments - GetAll  
Endpoint  `{url}:{port}/orgs/{org-name}/comments`

Method:`GET`

Header: `Content-Type: application/json`

Body: -

Response
```  
[
  {
    "_id": "5d714fa89d912b3b602e867f",
    "org_name": "mycompany",
    "comment": "Looking to hire SE Asia's top dev talent!"
  }
]
``` 
#### Comments - Delete  

Endpoint  `{url}:{port}/orgs/{org-name}/comments`

Method:`DELETE`

Header: `Content-Type: application/json`

Body: -

Response
```  
{
    "message":  "success"
}
```  

### Members
#### Members - Publish  
Endpoint  `{url}:{port}/orgs/{org-name}/members`
Method:`POST`
Header: `Content-Type: application/json`
Body:
```
{ 
    "username": "Lily",
    "avatar_url": "https://i.pinimg.com/originals/47/1a/1a/471a1ad342659289433e05a611d206f8.png",
    "total_follower": 800,
    "total_following": 20
}
```
Response
```  
{
    "message": "success"
}
```  
#### Members - GetAll  
Endpoint  `{url}:{port}/orgs/{org-name}/members`

Method:`GET`

Header: `Content-Type: application/json`

Body: -

Response
```  
[
  {
    "_id": "5d7150ed9d912b3b602e8680",
    "org_name": "mycompany",
    "username": "Lily",
    "avatar_url": "https://i.pinimg.com/originals/47/1a/1a/471a1ad342659289433e05a611d206f8.png",
    "total_follower": 800,
    "total_following": 20
  }
]
``` 
#### Members - Delete
Endpoint  `{url}:{port}/orgs/{org-name}/members`

Method:`DELETE`

Header: `Content-Type: application/json`

Body: -

Response
```  
{
    "message": "success"
}
```  
