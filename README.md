# Rest API using GO-Lang (usings `clinics` datasets)
⭐ For API we are using `GIN` http server.
<br>⭐ Instead of database we are using JSON files in the `data` directory.
<br>⭐ Essentially it load info about some clinics from remote url(set in the config toml), and publish them in an online searchable APIs.
<br>⭐ We have demonstrated how to use `security related headers` for the endpoints.
<br>⭐ Set-up `docker-compose` for easier deployments.
<br>⭐ Set-up `GitHub actions` for CI purposes.
<br>⭐ Included `TEST` suites
<br>⭐ Included `POSTMAN` collections for easier testing.

<br>
<br>

## How to config
- API search data are stored in following configs:
    - data/dental-clinics.json
    - data/vet-clinics.json
- HTTP configurations (port, data-states) are in the toml file: `./conf.d/app.toml`

## API Security
We are taking following measures to ensure API security.
- Setting up CORS security headers.
- Setting up the 'attack-vector' headers

```
	//setting-up cors headers
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//setting-up security headers
	r.Use(secure.New(secure.Config{
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		//SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))
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