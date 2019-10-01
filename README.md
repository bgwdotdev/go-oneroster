# GoOneRoster

This project aims to implement a OneRoster compliant RESTful API
webserver in GO with a MongoDB backing for persistance

## Setup

Start up the api server with either envs or flags

### flags

```
goors \
    -k "mySecretKey" \
    -m "mongodb://myinstance.domain.com:27017" \
```

### envs

```
GOORS_AUTH_KEY='mySecretKey'
GOORS_AUTH_KEY_ALG='HS256'
GOORS_MONGO_URI='mongodb://myinstance.domain.com:27017'
```

### Upload Data

Import data using the 
(sis sync tool)[https://github.com/fffnite/go-oneroster-sis-sync]
or by making POST requests.


## Query examples

```bash
# login
curl "myserver.domain.com/ims/oneroster/v1p1/login" \
    -X POST \
    -d "clientid=$ci&clientsecret=$cs"

# Upsert user id 1
curl "myserver.domain.com/ims/oneroster/v1p1/users/1" \
    -X PUT \
    -H "Authorization: Bearer $t"
    -d '{"sourcedId": "1", "status": "active", "givenName": "bob"}'

# Get all users
curl "myserver.domain.com/ims/oneroster/v1p1/users" \
    -H "Authorization: Bearer $t"
```

```powershell
# login
$args = @{
    uri = "http://myserver.domain.com/ims/oneroster/v1p1/login"
    method = "POST"
    body = @{ "clientid" = $ci; "clientsecret" = $cs }
}
$token = Invoke-RestMethod @args

# Upsert user id 1
$upsert = @{
    uri = "http://myserver.domain.com/ims/oneroster/v1p1/users/1
    method = "PUT"
    headers = @{ "authorization" = "bearer $t"}
    body = "{""sourcedId"": ""1"", ""status"": ""active"", ""givenName"": ""bob""}"
}
Invoke-RestMethod @upsert

$getUsers = @{
    uri = "http://myserver.domain.com/ims/oneroster/v1p1/users"
    heades = @{ "authorization" = "bearer $t" }
    FollowRelLink = $true
}
Invoke-RestMethod @getUsers
```

## To Do
- [x] Connect to SQL database
- [x] Read from Conf.hcl
- [x] Build DB handler
- [x] Connect packages
- [x] Output from DB
- [x] Output JSON
- [x] Output correct json from single endpoint sample
- [x] Implement RESTful parameters & operators
- [x] Implement logging
- [x] Implement error handling
- [x] Implement security
- [x] Build core endpoints
- [x] Sync DB
- [ ] Build extra endpoints
- [ ] Test test test
