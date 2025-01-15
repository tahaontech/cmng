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

# Install vcpkg

# Function to install vcpkg
install_vcpkg() {
    echo "vcpkg not found. Installing..."
    # Clone the vcpkg repository
    git clone https://github.com/microsoft/vcpkg.git
    cd vcpkg || { echo "Failed to change directory"; exit 1; }
    
    # Bootstrap vcpkg
    ./bootstrap-vcpkg.sh || { echo "Failed to bootstrap vcpkg"; exit 1; }
    
    echo "vcpkg installed successfully."
}

# Function to update ~/.bashrc
update_bashrc() {
    local VCPKG_PATH=$(pwd)/vcpkg
    echo "Adding vcpkg environment variables to ~/.bashrc..."
    {
        echo ""
        echo "# vcpkg environment setup"
        echo "export VCPKG_ROOT=${VCPKG_PATH}"
        echo "export PATH=\$VCPKG_ROOT:\$PATH"
    } >> ~/.bashrc
    echo "Environment variables added. Reload your shell or run 'source ~/.bashrc' to apply changes."
}

# Check if vcpkg is already installed
if ! command -v ./vcpkg/vcpkg &> /dev/null; then
    if [ -d "vcpkg" ]; then
        echo "vcpkg directory exists but is not installed. Attempting to bootstrap."
        cd vcpkg || { echo "Failed to change directory"; exit 1; }
        ./bootstrap-vcpkg.sh || install_vcpkg
    else
        install_vcpkg
    fi
else
    echo "vcpkg is already installed."
fi

# Add environment variables to ~/.bashrc
update_bashrc

echo "dependencies Installation completed!"

echo "Installing cmng..."
# download build file
wget -O cmng https://github.com/tahaontech/cmng/releases/download/v1/cmng-debian

# make executeable
sudo chmod +x cmng

# mv to bin
sudo mv cmng /usr/local/bin/

echo "Installed successfully run (cmng help)"