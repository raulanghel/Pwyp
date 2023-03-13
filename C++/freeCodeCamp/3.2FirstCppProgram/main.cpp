#include <iostream>

consteval int get_value()
{
    return 6;
}

int main()
{
    std::cout << "number1";
    std::cout << "number2" << std::endl;
    constexpr int value = get_value();
    std::cout << value << std::endl;
    return 0;
}
