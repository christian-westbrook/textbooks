.PHONY: build clean default run

default: build run clean

build: out.exe

out.exe: square.cpp
	g++ -o out.exe square.cpp

run:
	./out.exe

clean:
	rm ./*.exe