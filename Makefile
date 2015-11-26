EXECUTABLE = whiteboard
DB = whiteboard.db

all: whiteboard.go
	go build whiteboard.go

run:
	./$(EXECUTABLE) web

clean:
	rm -rf $(EXECUTABLE) $(DB)
