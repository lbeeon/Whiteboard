To run in Docker, we first need to build it:
```
go get github.com/hunterpraska/Whiteboard
cd $GOPATH/src/github.com/hunterpraska/Whiteboard
GOOS=linux GOARCH=amd64 go build
docker build -t whiteboard $GOPATH/src/github.com/hunterpraska/Whiteboard
```

Then just ``` docker run -it -d -p 8080:8080 --name whiteboard whiteboard```


### TODO
Add Docker build to one of the Docker repositories.
