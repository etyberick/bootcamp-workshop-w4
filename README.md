# Bootcamp Workshop W4 Mini Coding Challenge

## Background

You're the new developer in a Go backend team and you were given a small project, an API service about geometric figures.

It seems that the person who made the last commit forgot to commit the last changes and it accidentally broke the project. All you know about this project is that it was made with Go-Mockery.

## The challenge

You were given the task of making 3 things:
* You have to make it easy to build, your team lead said explicitly that he wants to clone the project and run it by using `go build`, right now when the lead is trying to run the project it breaks.
* You have to change all http status codes that were declared as integers and use instead the codes in the http library.
* You have to add a test case that requests an API endpoint that doesn't exist, it must return an error 404, please don't use integers for that.