# Rest API using GO-Lang (clinics-apis)
- We are using GIN http server.
- Essentially it load info about some clinics from remote url(set in the config toml), and publish them in an online searchable APIs.


## How to config

create a toml file in `./conf.d/app.toml`

```toml
# listen addr of http api server
addr = ":5000"

# urls of remote clinics 
urls = ["http://127.0.0.1:5000/data/dental-clinics.json",
"http://127.0.0.1:5000/data/dental-clinics.json"]

[states]
AL= "Alabama"
AK= "Alaska"
AZ= "Arizona"
AR= "Arkansas"
CA= "California"
CO= "Colorado"
CT= "Connecticut"
DE= "Delaware"
DC= "District of Columbia"
FL= "Florida"
GA= "Georgia"
HI= "Hawaii"
ID= "Idaho"
IL= "Illinois"
IN= "Indiana"
IA= "Iowa"
KS= "Kansas"
KY= "Kentucky"
LA= "Louisiana"
ME= "Maine"
MD= "Maryland"
MA= "Massachusetts"
MI= "Michigan"
MN= "Minnesota"
MS= "Mississippi"
MO= "Missouri"
MT= "Montana"
NE= "Nebraska"
NV= "Nevada"
NH= "New Hampshire"
NJ= "New Jersey"
NM= "New Mexico"
NY= "New York"
NC= "North Carolina"
ND= "North Dakota"
OH= "Ohio"
OK= "Oklahoma"
OR= "Oregon"
PA= "Pennsylvania"
RI= "Rhode Island"
SC= "South Carolina"
SD= "South Dakota"
TN= "Tennessee"
TX= "Texas"
UT= "Utah"
VT= "Vermont"
VA= "Virginia"
WA= "Washington"
WV= "West Virginia"
WI= "Wisconsin"
WY= "Wyoming"

```

## How to run

```
go build -o clinics-apis
./clinics-apis
```
## How to run test

```
go clean -testcache  
go test ./...
```

## How to find clinics 

- Use the POSTMAN collection JSON config and test there or you can use CURL.
- Name, State and Availability criteria are supported. multiple criteria is supported too.

### search it with criteria: name
```
curl http://127.0.0.1:5000/search?name=good
```

### search it with criteria:state code/name

```
curl http://127.0.0.1:5000/search?state=ca
```

```
curl http://127.0.0.1:5000/search?state=California
```


### search it with criteria:availability

```
curl http://127.0.0.1:5000/search?opening=12:00
```

### search it with multiple criterias: name and state 

```
curl http://127.0.0.1:5000/search?opening=12:00&name=center&state=ca
```

## How to build docker images

```
docker-compose up
```