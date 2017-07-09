# Snap Case Study

Do you have existing APIs or do you love to jump into the code first and then build the docs along with your code? No problem! Go Swagger has you covered. This case study uses [Snap](https://github.com/intelsdi-x/snap), the open telemetry framework as an example to walk you through the basic building blocks of go-swagger code. All code examples are from three Snap repositories.

> [Snap REST APIs](https://github.com/intelsdi-x/snap/tree/master/mgmt/rest/v2)  
> [Snap API Client](https://github.com/intelsdi-x/snap-client-go)  
> [Snap CLIs](https://github.com/intelsdi-x/snap-cli)  

## Index

1. [Setup go-swagger](#setup-go-swagger)
2. [Annotate APIs](#annotate-apis)
3. [Generate Specification](#generate-specification)
3. [Create Docs](#create-docs)  
4. [Generate Client](#generate-client)  
5. [Use Client](#use-client) 
6. [Last Not Least](#last-not-least) 

## Setup go-swagger

If your project uses Go, using the `go get` command in your project $GOPATH is simplest. Otherwise, refering to [go-swagger](https://github.com/go-swagger/go-swagger) for other installation methods.

```sh
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

Now type the following command to see if go-swagger is correctly installed.

```sh
▶ swagger -h
Usage:
  swagger [OPTIONS] <command>

Swagger tries to support you as best as possible when building APIs.

It aims to represent the contract of your API with a language agnostic description of your application in json or yaml.


Help Options:
  -h, --help  Show this help message

Available commands:
  expand    expand $ref fields in a swagger spec
  flatten   flattens a swagger document
  generate  genererate go code
  init      initialize a spec document
  mixin     merge swagger documents
  serve     serve spec and docs
  validate  validate the swagger document
  version   print the version
```

## Annotate APIs
The following elements describe a REST API.

| API Element       | Description                                                   |
|:------------------|:--------------------------------------------------------------|
| Endpoint          | a uri that accepts web requests. e.g./tasks or /tasks/{id}.   |
| Method            | one of the HTTP verbs. GET, POST, PUT, PATCH, DELETE.         |
| Parameters        | the path, query, and body parameters.                         |
| Responses         | the API request reponses including error responses            |

If you are annotating those element for an API, you're done. go-swagger can generate API specification, API docs, and API client SDKs.  

### Global `doc.go`
There is no reason to repeat the global, common or default API information in each API. You put them in a `doc.go` file inside your project `main` package. You can view Snap's [doc.go](https://github.com/intelsdi-x/snap/blob/master/doc.go) here. The mime type `application/json` is the default for every API attribute `consumes` and `produces`.


### API Routing

An API routing normally comprises an endpoint and a method. The syntax is
> swagger:route [method] [endpoint/path pattern] [?tag1 tag2 tag3] [operationid1 operationid2]

_**Example Route**_
```go
// swagger:route GET /plugins/{ptype}/{pname}/{pversion} plugins getPlugin
//
// Get
//
// An error will be returned if the plugin does not exist.
//
// Produces:
// application/json
//
// Schemes: http, https
//
// Responses:
// 200: PluginResponse
// 400: ErrorResponse
// 404: ErrorResponse
// 500: ErrorResponse
// 401: UnauthResponse
api.Route{Method: "GET", Path: prefix + "/plugins/:type/:name/:version", Handle: s.getPlugin}
```
In the above swagger:route `plugins` is a tag and `getPlugin` is an operation. Any attribute defined locally in each API overwrites the same attribute defined globally in the global `doc.go` file.


### API Parameters

An API path paramter is in the URI path. E.g. `publisher/file/1`. The query parameter is in the URL after the `?` mark. E.g. `type=publisher&name=file&ver=1`. The body paramter could be a JSON, XML, or a binary file.

| Parameter    | Example                                                         |
|:-------------|:-----------------------------------------------------------     |
| Path         | http://localhost:8181/v2/plugins/publisher/file/1               |
| Query        | http://localhost:8181/v2/plugins?type=publish&name=file&ver=1   |
| Body         | POST {"user":"jean"} to http://localhost:8181/v2/configs        |

The parameter type could be different. The annotation syntax is the same.
> swagger:parameters [operationid1 operationid2]

_**Example of Path Parameter**_

```go
// swagger:parameters getPlugin unloadPlugin getPluginConfigItem setPluginConfigItem
type PluginParams struct {
	// required: true
	// in: path
	PName string `json:"pname"`
	// required: true
	// in: path
	PVersion int `json:"pversion"`
	// required: true
	// in: path
	// enum: collector, processor, publisher
	PType string `json:"ptype"`
}
```

_**Example of Query Parameter**_

```go
// swagger:parameters getPlugins
type PluginsParams struct {
	// in: query
	Name string `json:"name"`
	// in: query
	// enum: collector, processor, publisher
	Type string `json:"type"`
	// in: query
	Running bool `json:"running"`
}
```

_**Example of Body Parameter**_

```go
// swagger:parameters addTask
type TaskPostParams struct {
	// Create a task.
	//
	// in: body
	//
	// required: true
	Task Task `json:"task"yaml:"task"`
}
```

### API Responses
API responses could be a request or an error response. The annotaion syntax is the same.
> swagger:response [?response name]

```go
// ErrorResponse represents the Snap error response type.
//
// It includes an error message and a map of fields.
//
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in:body
	SnapError Error `json: "snap_error"`
}
```

```go
// TaskResponse returns a task.
//
// swagger:response TaskResponse
type TaskResp struct {
	// in: body
	Task Task `json:"task"`
}
```

We have just covered all types of annotations used in Snap. Those are all we have to do when we're coding. Let's see how go-swagger can help us next.


## Generate Specification
If you annotate your main package with
> //go:generate swagger generate spec -o swagger.json

You can generate the specification with `go generate` in the main package. Snap puts `make swagger` as the default for the [make](https://github.com/intelsdi-x/snap/blob/master/Makefile#L63) command. Both `swagger generate spec` and `swagger validate` commands are inside a [script](https://github.com/intelsdi-x/snap/blob/master/scripts/swagger.sh#L28) file. it generates a specification first, then validates the spec afterwards.


## Create Docs
There is a [doc](https://github.com/go-swagger/go-swagger/blob/master/docs/usage/serve_ui.md) describing how to create the documentation for your APIs. In Snap, we use the following command to view API docs locally:

```sh
$swagger serve swagger.json --host=127.0.0.1
2017/07/06 15:21:11 serving docs at http://127.0.0.1:51154/docs
```

![alt text](https://user-images.githubusercontent.com/13841563/27982457-752271e0-6356-11e7-8636-5a89a2521571.png)


> If you prefer the representation of PetStore, you can paste Snap [swagger.json](https://raw.githubusercontent.com/intelsdi-x/snap/master/swagger.json) link into [PetStore](http://petstore.swagger.io) to explore. If you like to try out the APIs from the PetStore interactively, please enable CORS in Snap.

![alt text](https://user-images.githubusercontent.com/13841563/27982525-3a720af4-6358-11e7-8685-85446bf4fc42.png)


## Generate Client

Generating API client SDKs is really simple. Snap uses this command:

```sh
$swagger generate client -f swagger.json -A snap
```
The syntax is
> swagger generate client -f [http-url|filepath] -A [application-name]

Currently Snap obtains the API specification from the `vendor` directory. The specification may be published in the future. If you look into the generated client SDKs, two packages `client` and `models` are generated. Operations are grouped inside the annotated tag names, plugins and tasks. The `snap_client.go` has functions/methods for creating an API client to interact with server APIs.

```sh
├── client
│   ├── plugins
│   ├── snap_client.go
│   └── tasks
├── models
```

If you peek into the generated `tasks` package, the file `task_client.go` has all operations that can communicate with the `task` server APIs. The `*_parameters.go` provides methods/functions for setting input parameters. The `*_responses.go` provides methods/functions for getting appropriate responses.

```sh
└── tasks
│       ├── add_task_parameters.go
│       ├── add_task_responses.go
│       ├── get_task_parameters.go
│       ├── get_task_responses.go
│       ├── get_tasks_parameters.go
│       ├── get_tasks_responses.go
│       ├── remove_task_parameters.go
│       ├── remove_task_responses.go
│       ├── tasks_client.go
│       ├── update_task_state_parameters.go
│       ├── update_task_state_responses.go
│       ├── watch_task_parameters.go
│       └── watch_task_responses.go
```
You can type the following command to check the available client options.

```sh
▶ swagger generate client -h
Usage:
  swagger [OPTIONS] generate client [client-OPTIONS]

generate all the files for a client library

Help Options:
  -h, --help                  Show this help message

[client command options]
      -f, --spec=             the spec file to use (default swagger.{json,yml,yaml})
      -a, --api-package=      the package to save the operations (default: operations)
      -m, --model-package=    the package to save the models (default: models)
      -s, --server-package=   the package to save the server specific code (default: restapi)
      -c, --client-package=   the package to save the client specific code (default: client)
      -t, --target=           the base directory for generating the files (default: ./)
      -T, --template-dir=     alternative template override directory
      -C, --config-file=      configuration file to use for overriding template options
          --existing-models=  use pre-generated models e.g. github.com/foobar/model
      -A, --name=             the name of the application, defaults to a mangled value of info.title
      -O, --operation=        specify an operation to include, repeat for multiple
          --tags=             the tags to include, if not specified defaults to all
      -P, --principal=        the model to use for the security principal
      -M, --model=            specify a model to include, repeat for multiple
          --default-scheme=   the default scheme for this client (default: http)
          --default-produces= the default mime type that API operations produce (default: application/json)
          --skip-models       no models will be generated when this flag is specified
          --skip-operations   no operations will be generated when this flag is specified
          --dump-data         when present dumps the json for the template generator instead of generating files
          --skip-validation   skips validation of spec prior to generation
```

## Use Client

[snap-cli]() provides interaction of REST APIs through CLIs. This repo fully utilizes the go-swagger generated API client SDKs. 

### CMD: snaptel task list
List all scheduled tasks

_**Equivalent CURL Request**_
```
curl -L http://localhost:8181/v2/tasks
```

_**Example Client Get**_
```go
params := tasks.NewGetTasksParams()
getTasks, err := c.Tasks.GetTasks(params)
```

_**Example CLI Response**_
```sh
▶ snaptel task list
ID 					 NAME 						 STATE 		 HIT 	 MISS 	 FAIL 	 CREATED 			 LAST FAILURE
f1975f92-a2c3-4f3d-b56e-29c74fd082d2 	 test                                      	 Running 	 3 	 0 	 0 	 Fri, 07 Jul 2017 22:40:26 PDT
46236676-6110-4797-897d-cab46e15e411 	 test                                      	 Running 	 1 	 0 	 0 	 Fri, 07 Jul 2017 22:40:55 PDT
27b1a783-561a-4102-bbed-795f963e7a2d 	 Task-27b1a783-561a-4102-bbed-795f963e7a2d 	 Running 	 1 	 0 	 0 	 Fri, 07 Jul 2017 22:41:04 PDT
```

### CMD: snaptel task create -t {task manifest}
Create a task with the JSON/YAML input using, for example, mock-file.json with the following content:
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "max-failures": 10,
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/mock/foo": {},
                "/intel/mock/bar": {},
                "/intel/mock/*/baz": {}
            },
            "config": {
                "/intel/mock": {
                    "name": "root",
                    "password": "secret"
                }
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "mock-file",
                            "config": {
                                "file": "/tmp/published"
                            }
                        }
                    ]
                }
            ]
        }
    }
}
```

_**Equivalent CURL Request**_
```
curl -vXPOST http://localhost:8181/v2/tasks -d @mock-file.json --header "Content-Type: application/json"
```

_**Example Client Post**_
```go
params := tasks.NewAddTaskParams()
params.SetTask({task manifest} || {workflow manifest})

createTask, err := c.Tasks.AddTask(params, nil)
```

_**Example CLI Response**_
```sh
▶ snaptel task create -t ~/task/task-mock-file.json
Task created
ID: 27b1a783-561a-4102-bbed-795f963e7a2d
Name: Task-27b1a783-561a-4102-bbed-795f963e7a2d
State: Running
```

### CMD snaptel task stop {id}
Stop a running task given a task ID

_**Equivalent CURL Request**_
```
curl -XPUT http://localhost:8181/v2/tasks/27b1a783-561a-4102-bbed-795f963e7a2d/stop
```

_**Example Client Put**_
```go
params := tasks.NewUpdateTaskStateParams()
params.SetID("27b1a783-561a-4102-bbed-795f963e7a2d")
params.SetAction("stop")

stopTask, err := c.Tasks.UpdateTaskState(params, nil)
```
_**Example Response**_
```sh
▶ snaptel task stop 27b1a783-561a-4102-bbed-795f963e7a2d
Task stopped:
ID: 27b1a783-561a-4102-bbed-795f963e7a2d
```

### CMD: snaptel task remove {id}
Remove a task from the scheduled task list given a task ID

_**Equivalent CURL Request**_
```
curl -X DELETE http://localhost:8181/v1/tasks/7cd4b229-e12c-4b09-985a-b60e76daac90  
```

_**Example Client Delete**_
```go
params := tasks.NewRemoveTaskParams()
params.SetID("27b1a783-561a-4102-bbed-795f963e7a2d")

removeTask, err := c.Tasks.RemoveTask(params, nil)
```

_**Example Response**_
```sh
▶ snaptel task remove 27b1a783-561a-4102-bbed-795f963e7a2d
Task removed:
ID: 27b1a783-561a-4102-bbed-795f963e7a2d
```

## Last Not Least

When you define a struct, it's the best that the variable name matches the JSON representation. For example

```go
//swagger:parameters setPluginConfigItem
type PluginConfigParam struct {
	// in: body
	Config map[string]interface{} `json:"config"`
}
```

You can checkout [go-swagger.io](https://goswagger.io/) to gain more in depth knowledge of go-swagger.

You're welcome to browse Snap's repos to see how we adopted go-swagger. Your suggestions and improvements are welcome. Many thanks to go-swagger team for making our work possible!

