package setup

import "fmt"

func generateVsCodeConfig() string {
	return `
{
    "configurations": [
        {
            "name": "Linux",
            "includePath": [
                "${workspaceFolder}/include",
                "/usr/include",
                "/usr/local/include",
                "${env:VCPKG_ROOT}/installed/x64-linux/include"
            ],
            "defines": [],
            "compilerPath": "/usr/bin/g++",
            "cStandard": "c11",
            "cppStandard": "c++17",
            "intelliSenseMode": "linux-gcc-x64"
        }
    ],
    "version": 4
}
`
}

func generateCMakeLists(projectName string, withTests bool) string {
	cmake := fmt.Sprintf(`cmake_minimum_required(VERSION 3.10)

# Project Name
project(%s VERSION 1.0 LANGUAGES CXX)

# Specify the C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Include directories
include_directories(${PROJECT_SOURCE_DIR}/include)

# Add source files
set(SOURCE_FILES
    src/main.cpp
    src/my_class.cpp
)

# Add fmt from vcpkg
find_package(fmt CONFIG REQUIRED)

# Add executable
add_executable(%s ${SOURCE_FILES})

# link fmt from vcpkg
target_link_libraries(%s PRIVATE fmt::fmt)
`, projectName, projectName, projectName)

	if withTests {
		cmake += `
# Add tests
enable_testing()
add_subdirectory(tests)
`
	}

	return cmake
}

func generateReadme(projectName string) string {
	return fmt.Sprintf(`# %s
## Intro

A basic C++ project with "cmng" tool.

read Docs: [TODO]


`, projectName)
}

func generateGitignore() string {
	return `
cmake
build
`
}

func generateMainCpp() string {
	return `
#include <iostream>
#include "my_class.h"

int main() {
    MyClass obj;
    obj.greet();
    return 0;
}`
}

func generateClassH() string {
	return `
#ifndef MY_CLASS_H
#define MY_CLASS_H

class MyClass {
public:
    void greet();
};

#endif // MY_CLASS_H
`
}

func generateClassCpp() string {
	return `
#include "my_class.h"
#include <fmt/core.h>

void MyClass::greet() {
    fmt::print("Hello World!\n");
}

`
}

func generateTestsCmake() string {
	return `
find_package(GTest REQUIRED)
include_directories(${GTEST_INCLUDE_DIRS})

add_executable(test_my_class test_my_class.cpp)

target_link_libraries(test_my_class ${GTEST_LIBRARIES} pthread)

add_test(NAME MyClassTest COMMAND test_my_class)
	`
}

func generateClassTest() string {
	return `
#include "gtest/gtest.h"
#include "my_class.h"
	
TEST(MyClassTest, GreetTest) {
	MyClass obj;
	testing::internal::CaptureStdout();
	obj.greet();
	std::string output = testing::internal::GetCapturedStdout();
	EXPECT_EQ(output, "Hello, world!\n");
}
`
}
