# Swagger 2.0 [![Build Status](https://circleci.com/gh/go-swagger/go-swagger.svg?style=shield)](https://circleci.com/gh/go-swagger/go-swagger) [![Build status](https://ci.appveyor.com/api/projects/status/x377t5o9ennm847o/branch/master?svg=true)](https://ci.appveyor.com/project/casualjim/go-swagger/branch/master) [![codecov](https://codecov.io/gh/go-swagger/go-swagger/branch/master/graph/badge.svg)](https://codecov.io/gh/go-swagger/go-swagger) [![GitHub version](https://badge.fury.io/gh/go-swagger%2Fgo-swagger.svg)](https://badge.fury.io/gh/go-swagger%2Fgo-swagger)
[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/swagger-api/swagger-spec/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/ssfilatov/go-swagger?status.svg)](http://godoc.org/github.com/ssfilatov/go-swagger)
[![Docker Repository on Quay](https://quay.io/repository/goswagger/swagger/status "Docker Repository on Quay")](https://quay.io/repository/goswagger/swagger)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgo-swagger%2Fgo-swagger.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgo-swagger%2Fgo-swagger?ref=badge_shield)
[![GolangCI](https://golangci.com/badges/github.com/ssfilatov/go-swagger.svg)](https://golangci.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/ssfilatov/go-swagger)](https://goreportcard.com/report/github.com/ssfilatov/go-swagger)

This package contains a golang implementation of Swagger 2.0 (aka [OpenAPI 2.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md)):
it knows how to serialize and deserialize swagger specifications.

[Swagger](https://swagger.io/) is a simple yet powerful representation of your RESTful API.<br>

> ![swagger](https://raw.githubusercontent.com/go-swagger/go-swagger/master/docs/favicon-16x16.png) **Swagger in a nutshell**
>
> With the largest ecosystem of API tooling on the planet, thousands of developers are supporting Swagger in almost every modern programming language and deployment environment.
>
> With a Swagger-enabled API, you get interactive documentation, client SDK generation and discoverability. We created Swagger to help fulfill the promise of APIs.
>
> Swagger helps companies like Apigee, Getty Images, Intuit, LivingSocial, McKesson, Microsoft, Morningstar, and PayPal build the best possible services with RESTful APIs. Now in version 2.0, Swagger is more enabling than ever. And it's 100% open source software.

##  Features
`go-swagger` brings to the go community a complete suite of fully-featured, high-performance, API components to  work with a Swagger API: server, client and data model.

* Generates a server from a swagger specification
* Generates a client from a swagger specification
* Generates a CLI (command line tool) from a swagger specification (alpha stage)
* Supports most features offered by jsonschema and swagger, including polymorphism
* Generates a swagger specification from annotated go code
* Additional tools to work with a swagger spec
* Great customization features, with vendor extensions and customizable templates

Our focus with code generation is to produce idiomatic, fast go code, which plays nice with golint, go vet etc.

##  Project status
`go-swagger` is now feature complete and has stabilized its API.

Most features and building blocks are now in a stable state, with a rich set of CI tests.

The go-openapi community actively continues bringing fixes and enhancements to this code base.

There is still much room for improvement: contributors and PR's are welcome. You may also get in touch with maintainers on [our slack channel](https://slackin.goswagger.io).

## Documentation
<https://goswagger.io>

## FAQ
Q&A contributed by the community:

<https://goswagger.io/faq/>

## How is this different from go generator in swagger-codegen?
**tl;dr** The main difference at this moment is that this one actually works...

The swagger-codegen project only generates a _workable_ go client and even there it will only support flat models.
Further, the go server generated by swagger-codegen is mostly a stub.

> **Motivation**
> Why is this not done as a part of the swagger-codegen project? Because:
>
> * I don't really know java very well and so I'd be learning both java and the object model of the codegen which was in heavy flux as opposed to doing go and I really wanted to go experience of designing a large codebase with it.
> * Go's super limited type system makes it so that it doesn't fit well in the model of swagger-codegen
> * Go's idea of polymorphism doesn't reconcile very well with a solution designed for languages that actually have inheritance and so forth.
> * For supporting types like `[][][]map[string][][]int64` I don't think it's possible with mustache
>
> I gravely underestimated the amount of work that would be involved in making something useful out of it.
> My personal mission: I want the jvm to go away, it was great way back when now it's just silly (vm in container on vm in vm in container)

## What's inside?

Here is an outline of available features (see the full list [here](https://goswagger.io/features.html)):

- An object model that serializes swagger-compliant yaml or json
- A tool to work with swagger
  - Serve swagger UI for any swagger spec file
  - Flexible code generation, with customizable templates
    - Generate go API server based on swagger spec
    - Generate go API client from a swagger spec
  -  Validate a swagger spec document, with extra rules outlined [here](https://github.com/apigee-127/sway/blob/master/docs/README.md#semantic-validation)
  -  Generate a spec document based on annotated code
- A runtime to work with Rest API and middlewares
  - Serve spec
  - Routing
  - Validation
  - Authorization
  - Swagger docs UI
  - A Diff tool which will cause a build to fail if a change in the spec breaks backwards compatibility

There is more to that...

- A [typed JSON Schema implementation](https://goswagger.io/use/model.html), supporting most Draft 4 features
- Extended string and numeric formats: [strfmt](https://github.com/go-openapi/strfmt)
- Utilities to work with JSON, convert data types and pointers: [swag](https://github.com/go-openapi/swag)
- A jsonschema (Draft 4) validator, with full $ref support: [validate](https://github.com/go-openapi/validate)
- Custom validation interface

## Installing
`go-swagger` is available as binary or docker releases as well as from source: [more details](https://goswagger.io/install.html).

## Use-cases
The main package of the toolkit, go-swagger/go-swagger, provides command line tools to help working with swagger.

The toolkit is highly customizable and allows endless possibilities to work with OpenAPI2.0 specifications.

Beside the go-swagger CLI tool and generator, the [go-openapi packages](https://github.com/go-openapi) provide modular functionality to build custom solutions on top of OpenAPI.

The CLI supports shell autocompletion utilities: see [here](https://goswagger.io/cli_helpers.html).

### Serve specification UI
Most basic use-case: serve a UI for your spec:

```
swagger serve https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json
```

### Validate a specification
To [validate](https://goswagger.io/usage/validate.html) a Swagger specification:

```
swagger validate https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json
```

### Generate an API server
To generate a [server for a swagger spec](https://goswagger.io/generate/server.html) document:

```
swagger generate server [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```

### Generate an API client
To generate a [client for a swagger spec](https://goswagger.io/generate/client.html) document:

```
swagger generate client [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```
### Generate an CLI (Command line tool)
To generate a [CLI for a swagger spec](https://github.com/ssfilatov/go-swagger/tree/master/examples/cli) document:
```
swagger generate cli [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```
### Generate a spec from source
To generate a [swagger spec document for a go application](https://goswagger.io/generate/spec.html):

```
swagger generate spec -o ./swagger.json
```

### Generate a data model
To generate model structures and validators exposed by the API:

```
swagger generate model --spec={spec}
```

### Transform specs

There are [several commands](https://goswagger.io/use/transform.html) allowing you to transform your spec.

Resolve and expand $ref's in your spec as inline definitions:
```
swagger expand {spec}
```

Flatten your spec: all external $ref's are imported into the main document and inline schemas reorganized as definitions.
```
swagger flatten {spec}
```

Merge specifications (composition):
```
swagger mixin {spec1} {spec2}
```

### Compare specs

The  diff command allows you to check backwards compatibility.
Type ```swagger diff --help``` for info.

```
swagger diff {spec1} {spec2}
```

### Generate spec markdown spec

```
swagger generate markdown -f {spec} --output swagger.mode
```

## Try it

Try `go-swagger` in a free online workspace using Gitpod:

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io#https://github.com/ssfilatov/go-swagger)

## Licensing

The toolkit itself is licensed as Apache Software License 2.0. Just like swagger, this does not cover code generated by the toolkit. That code is entirely yours to license however you see fit.


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgo-swagger%2Fgo-swagger.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgo-swagger%2Fgo-swagger?ref=badge_large)

## Who is using this project?

To name but a few... (feel free to sign in there if you are using this project):

> In the list below, we tried to figure out the public repos where you'll find examples on how to use `go-swagger` and `go-openapi`:

[3DSIM](https://github.com/3DSIM)  
[Alibaba PouchAPI](https://github.com/alibaba/pouch)  
[CheckR](https://github.com/checkr/flagr)  
[Cilium](https://github.com/cilium/cilium)  
[CoreOS](https://github.com/coreos/go-quay)  
[NetBox Community](https://github.com/netbox-community/go-netbox)  
[EVE Central](https://github.com/evecentral)  
Iron.io
[JaegerTracing](https://github.com/jaegertracing/jaeger)  
[Kubernetes-Helm](https://github.com/kubernetes-helm/monocular)  
[Kubernetes](https://godoc.org/k8s.io/apiextensions-apiserver/pkg/apiserver)  
[ManifoldCo](https://github.com/manifoldco)  
[Metaparticle.io](https://github.com/metaparticle-io/metaparticle-ast)  
[Netlify](https://github.com/netlify/open-api)  
[Nutanix](https://github.com/nutanix)  
[OAS2](https://github.com/hypnoglow/oas2)  
[OVH API](https://github.com/appscode/go-ovh)  
[RackHD](https://github.com/RackHD/RackHD)  
[ScaleFT](https://github.com/authclub/billforward)  
[StratoScale](https://github.com/Stratoscale/swagger)  
[Terraform Provider OpenAPI](https://github.com/dikhan/terraform-provider-openapi)  
[VMware](https://github.com/vmware/dispatch)  
...

## Note to users migrating from older releases

### Migrating from 0.25 to [master]

Changes in the behavior of the generated client regarding defaults in parameters and response headers:

  * default values for parameters are no more hydrated by default and sent over the wire
    (assuming the server uses defaults).
  * the previous behavior (explicitly sending defaults over the wire) can be obtained
    with the SetDefaults() and WithDefaults() parameter methods.
  * the body parameter is not pre-hydrated with the default from it schema
  * default values for response headers are hydrated when the header is not received
    (previously, headers remained with their zero value)

### Migrating from 0.24 to 0.25

The options for `generate model --all-definitions` and `--skip-struct` are marked for deprecation. 

For now, the CLI continues to accept these options. They will be removed in a future version.

Generating all definitions is now the default behavior when no other option filters the generation scope.
The `--skip-struct` option had no effect.

### Migrating from 0.14 to 0.15

Generated servers no more import the following package (replaced by go1.8 native functionality):
```
github.com/tylerb/graceful
```

Spec flattening now defaults to minimal changes to models and should be workable for 0.12 users.

Users who prefer to stick to 0.13 and 0.14 default flattening mode may now use the `--with-flatten=full` option.

Note that the `--skip-flatten` option has been phased out and replaced by the more explicit `--with-expand` option.

### Migrating from 0.12 to 0.13

Spec flattening and $ref resolution brought breaking changes in model generation, since all complex things generate their own definitions.

### Migrating from 0.5.0 to 0.6.0

You will have to rename some imports:

```
github.com/ssfilatov/go-swagger/httpkit/validate to github.com/go-openapi/validate
github.com/ssfilatov/go-swagger/httpkit to github.com/go-openapi/runtime
github.com/naoina/denco to github.com/go-openapi/runtime/middleware/denco
github.com/ssfilatov/go-swagger to github.com/go-openapi
```

### Using 0.5.0

Because 0.5.0 and master have diverged significantly, you should checkout the tag 0.5.0 for go-swagger when you use the currently released version.
