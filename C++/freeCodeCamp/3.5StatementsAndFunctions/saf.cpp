#include <iostream>

int addTwoNumbers(int first_number, int second_number)
{
    return first_number + second_number;
}

int main(int argc, char const *argv[])
{
    int first_number = 22;
    int second_number = 37;

    std::cout << addTwoNumbers(first_number, second_number) << std::endl;

    return 0;
}
