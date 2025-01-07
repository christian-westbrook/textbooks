#include <iostream> // Import declarations for iostream

using namespace std; // Make names from std visible without std::

double square(double x)
{
    return x*x;
}

void print_square(double x)
{
    cout << "the square of " << x << " is " << square(x) << "\n";
}

int main()
{
    print_square(1.234); // The square of 1.234 is 1.52276
}