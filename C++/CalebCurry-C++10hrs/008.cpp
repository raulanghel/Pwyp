#include <iostream>

int main(int argc, char const *argv[])
{
    int slices;

    std::cout << "How many slices do you want? ";
    std::cin >> slices;
    std::cout << "You have " << slices << " slices of pizza." << std::endl;

    return 0;
}
