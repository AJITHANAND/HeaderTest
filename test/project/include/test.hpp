#ifndef TEST_HPP
#define TEST_HPP

#include <iostream>
#include <vector>
#include <string>
#include "common/common.hpp"

class Test{
    public:
        Test();
        ~Test();
        void print();
        void print(std::string str);
        void print(std::vector<std::string> str);
        void exampleFunction();
};


#endif // TEST_HPP