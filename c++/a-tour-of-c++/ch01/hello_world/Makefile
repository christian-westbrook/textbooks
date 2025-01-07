.PHONY: build clean default run

default: build run clean

build: out.exe

out.exe: hello_world.cpp
	g++ -o out.exe hello_world.cpp

run:
	./out.exe

clean:
	rm ./*.exe