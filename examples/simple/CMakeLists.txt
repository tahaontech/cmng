cmake_minimum_required(VERSION 3.10)
project(MyProject)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)

# Include directories
include_directories(include)

# Add source files
set(SOURCE_FILES
    src/main.cpp
    src/my_class.cpp
)

# Create executable
add_executable(my_project ${SOURCE_FILES})

# Add tests
enable_testing()
add_subdirectory(tests)