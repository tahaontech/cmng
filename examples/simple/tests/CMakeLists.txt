# Add GoogleTest as a subdirectory (assuming it's installed)
find_package(GTest REQUIRED)
include_directories(${GTEST_INCLUDE_DIRS})

add_executable(test_my_class test_my_class.cpp)

target_link_libraries(test_my_class ${GTEST_LIBRARIES} pthread)

add_test(NAME MyClassTest COMMAND test_my_class)
