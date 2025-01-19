# kurl
api client with easy of maintenance and documentation 
why use `kurl`
- we can use git version control to collaborate and maintain the code
- api documentation can we added along with the codebase

## Installation
if golang is installed on our system then we could use the following methods
```sh
go install github.com/kolukattai/kurl
```

## Create Project
to create new project run `kurl init` followed by project name like mentioned bellow
```sh
kurl init myapp
```

to initialize project in current directory use the bellow command
```sh
kurl init .
```
this will create a project the `config.yaml` file in project root directory with `README.md` and `example-api-call.md` in `api/` which goes like in bellow

```sh
myapp
  |- api/
      |- example-api-call.md
      |- README.md
  |- config.yaml
```


## Create new endpoint
to create a new endpoint use `kurl add` command followed by the api call name with `-` separated to create a endpoint template like shown bellow
```sh
kurl add my-get-api
```
this will create endpoint inside `api/`(default endpoint folder) folder, `api/my-get-api.md` the endpoint file is divided into to segments, top part(front matter) to configure API Call bottom part for documentation

## Endpoint File Front Matter
| key | description |
| --- | --- |
| refID | this is predefined and should not be changed, it is used to link it's saved response with this request |
| method | http method for you API Call, currently allowed methods `GET`,`PUT`,`POST`,`DELETE`,`PATCH`
| url | url for the endpoint |
| queryParams | if we like to make our endpoint neat we could just use the field to organize url params eg, `http://localhost:3000/api?page=1` for this endpoint we can use write it like a object `queryParams: {page: 1}` and skip it in url




## Config File
