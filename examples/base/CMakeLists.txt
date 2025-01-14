cmake_minimum_required(VERSION 3.10)

# Project Name
project(base VERSION 1.0 LANGUAGES CXX)

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

# Add executable
add_executable(base ${SOURCE_FILES})
