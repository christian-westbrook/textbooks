blah.exe: blah.o
	gcc blah.o -o blah.exe # Runs third

blah.o: blah.c
	gcc -c blah.c -o blah.o # Runs second

# Typically blah.c would already exist, but I want to limit any additional required files
blah.c:
	echo "int main() { return 0; }" > blah.c # Runs first