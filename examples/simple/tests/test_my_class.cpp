
#include "gtest/gtest.h"
#include "my_class.h"

TEST(MyClassTest, GreetTest) {
    MyClass obj;
    testing::internal::CaptureStdout();
    obj.greet();
    std::string output = testing::internal::GetCapturedStdout();
    EXPECT_EQ(output, "Hello, world!\n");
}
