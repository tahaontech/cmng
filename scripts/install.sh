#!/bin/bash

echo "Installing cmng dependency"

# Update package list
echo "Updating package list..."
sudo apt update

# Install build essentials (GCC, Make, etc.)
echo "Installing build-essential..."
sudo apt install -y build-essential

# Install GDB (GNU Debugger)
echo "Installing GDB..."
sudo apt install -y gdb

# Install CMake (cross-platform build system)
echo "Installing CMake..."
sudo apt install -y cmake

# Install Clang (alternative compiler)
echo "Installing Clang..."
sudo apt install -y clang

# Install Git (version control system)
echo "Installing Git..."
sudo apt install -y git

# Install additional development libraries and headers
echo "Installing additional libraries..."
sudo apt install -y libc-dev

# Install Valgrind (for memory leak detection)
echo "Installing Valgrind..."
sudo apt install -y valgrind

# Install GProf (performance profiling tool)
echo "Installing GProf..."
sudo apt install -y gprof

# Install Google test
echo "Installing Google Test..."
sudo apt install libgtest-dev

# Install Doxygen (for documentation generation)
echo "Installing Doxygen..."
sudo apt install -y doxygen

# Optional: Install other tools (if desired)
# Install Boost (C++ libraries)
echo "Installing Boost libraries..."
sudo apt install -y libboost-all-dev

# Install SDL2 (for game/multimedia development)
echo "Installing SDL2 libraries..."
sudo apt install -y libsdl2-dev

echo "dependency Installation completed!"

echo "Installing cmng..."
# download build file
wget "TODO"

# make executeable
sudo chmod +x cmng

# mv to bin
sudo mv cmng /usr/local/bin/

echo "Installed successfully run (cmng help)"