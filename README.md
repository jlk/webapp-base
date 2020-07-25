# Webapp Base
This is my base framework for doing quick webapp development - a
[vue-based](https://vuejs.org/) webapp with the [CoreUI admin template](https://coreui.io/),
and a golang-based backend supporting Rest API with a postgres database.

It doesn't sound like a big deal, but when I have an idea I want to hack on, I don't want to waste time setting up a go project, getting it to build right in VSCode, adding node, remembering what modules I liked using from previous experiments. I want to write some business logic, add some data, put a front end on it and play with the idea. And preferrably do all that in a relatively secure way.

# README TODO* Add notes for getting client going

# Idea and design
When creating this project, the following ideas were put together for a dream list of 
 * go-based backend
   * [gin](https://github.com/gin-gonic/gin)
   * [auth0 authentication library](https://github.com/auth0-samples/auth0-golang-web-app) (in time...)
   * gin-jwt or just auth0's jwt support?
   * ~~graphql~~ (decided to go back to REST)
   * viper
   * ~~mgo~~ (nope - back to postgresql)
   * [SQLX](github.com/jmoiron/sqlx) and [PGX](github.com/jackc/pgx) for Postgres support
   * [go-cache](https://github.com/patrickmn/go-cache) for caching objects from db
   * 
 * vue-based frontend
   * bootstrap via CoreUI
 * Support for:
   * REST API
   * Postgres
   * serverless support? (maybe some day...)
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

![Image of browser screenshot](webapp%20base%20screenshot.png)

## Decisions
To give a little insight into some of the choices that have been made here:

### Typescript vs JavaScript
JavaScript is broken. I'm trying TypeScript on this project in hopes of improving appsec, after I was reminded of it by my Application Security Weekly co-host Mike Shema.

This isn't a free security upgrade - I'm still going through the pains of figuring out how to navigate typescript's errors when working with a language I was never super strong with (javascript/emcascript/whatever the kids call it this week). But from all I've seen it'll be worth it.

### React vs Vue
The first attempts with this project used [React](https://reactjs.org/). I put the work aside for a few months while work/etc kept me busy. Along that path I remember a few times looking at blog posts about [Vue](https://vuejs.org/) being better for new projects and thinking "yeah...I should probably do that instead" so when I picked this back up in 2020 I nuked the React (Don't think it was ever checked in) and moved to Vue. There's probably still more react devs out there, but Vue's clean, has decent Typescript support, and I think will fit my needs nicely.

### "Admin" template/framework
For me, context switching between back-end/front-end is super expensive. So I looked around for not just a bootstrap template that would look good, but an existing framework for Vue. After looking at a few, I really liked [CoreUI](https://coreui.io/vue/). It works well with Vue (doesn't feel like it was shoehorned in), is [Bootstrap](https://getbootstrap.com) based so TONS of documentation and developer knowledge, has a great set of components ready to go for the things I want (Tables, charts, widgets).

If you use this, I **strongly suggest you pay for a license**. Support is really great - I had a typescript issue on a Friday afternoon US time, and they were responding from Poland - saving me from having to wait through the weekend.

### ReST vs GraphQL
If you look at older versions of this README, you'll see discussion of Apollo vs Relay for GraphQL support. That just wasn't as nice/easy/clean as I wanted, and at the end of the day if I want to bring more developers into a project, REST is a no-brainer. KISS.

### Mongo vs SQL
This will get a laugh from compatriots at my last gig where I fought like crazy to not have Layered Insight ported from Mongo to Cassandra. I still believe that was the right decision there, but moving forward, over the last decade Postgres has picked up features and performance. Mysql is great for some quick hacking, but Postgres seems to have an edge in production environments.

I never got really really good at developing complex Mongo queries. I've used SQL since...1997 or so. Seems silly to fight that.

I still would require a hole in my head to choose to use Cassandra. :)

## Pre-requisites
To get started, the following items are needed for local R&D (osx reminders in parenthesis):
 * go
 * [go modules](https://github.com/golang/go/wiki/Modules) (this used to say "trying go modules" - in 2020 that ship's sailed)
 * npm?
 * docker

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

# Read the commit notes!
When I've done heavier lifts, like bringing in CoreUI, there's a good amount of commentary in the git commit messages. I strongly recommend reviewing these to see what was done and why.

# Thanks
Ideas for this came from previous experience, along with the following blog posts:
 * https://medium.com/@chrischuck35/how-to-build-a-simple-web-app-in-react-graphql-go-e71c79beb1d (not using, but I'll still give props as that guided me for a while)
 * https://medium.com/avitotech/how-to-work-with-postgres-in-go-bad2dabd13e4 interesting read for pgsql in go
 * https://github.com/tmrts/go-patterns good list of Go patterns

