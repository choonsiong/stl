all: clean
	go build -o server .

clean:
	rm -f server