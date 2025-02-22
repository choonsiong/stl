all: clean
	go build -o stl-server .

clean:
	rm -f stl-server