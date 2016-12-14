# GoogleAppEngineDeploy-Sandbox

Google App Engine Deploy Sandbox

### Download the App Engine SDK for Go
https://cloud.google.com/appengine/docs/go/download?hl=ja


### update app.yaml

```
application: <your api key>
version: 1
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
```

### serve command

```
goapp serve ./
```

### serve url
```
http://localhost:8080
```

### deploy command

```
goapp deploy  ./
```
