# Testapp
This is my base framework for doing quick webapp development - a
react-based webapp with a bootstrap theme that I like, and a
golang-based backend supporting GraphQL with a mongo database

# Idea and design
When creating this project, the following ideas were put together for a dream list of 
 * go-based backend
   * [gin](https://github.com/gin-gonic/gin)
   * [auth0 authentication library](https://github.com/auth0-samples/auth0-golang-web-app)
   * gin-jwt or just auth0's jwt support?
   * graphql
   * viper
   * [mgo](https://github.com/globalsign/mgo)
 * react-based frontend
   * bootstrap theme
 * Support for:
   * [graphql](https://graphql.org/)
   * mongo, something else? (what wouldn't need an external db?)
   * serverless support?
   * various local and remote authentiation frameworks via auth?
   * security?
 * CI/CD
   * build script for - circleci, or??
   * linters
   * unit tests
   * docker build
   * docker-compose for local R&D
   * k8s deployment?
 * release checklist - things to do/customize/test/document
 * offline dev is a plus

## Pre-requisites
To get started, the following items are needed for local R&D (osx reminders in parenthesis):
 * go
 * [dep](https://github.com/golang/dep) (`brew install dep`)
 * npm?
 * docker?

## Documentation
I fully expect to forget how to use this between uses, so some
reminders for myself, or starting thoughts for others:
 * Docs for usage

### GOPATH
Right now this is intended to be used for stand-alone usage, so
GOPATH should be set to under the server directory in this project:
```
export GOPATH={path to app directory}/server
```

## Intentions
 * keep db code separate so I can switch between databases fairly easily
 * Move golang things around to match structure at https://github.com/golang-standards/project-layout

# Thanks
Ideas for this came from previous experience, along with the following blog posts:
 * https://medium.com/@chrischuck35/how-to-build-a-simple-web-app-in-react-graphql-go-e71c79beb1d

