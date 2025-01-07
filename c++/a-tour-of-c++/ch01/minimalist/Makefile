.PHONY: build clean default run

default: build run clean

build: out.exe

out.exe: minimalist.cpp
	g++ -o out.exe minimalist.cpp

run:
	./out.exe

clean:
	rm ./*.exe