# CPPManager

A lightweight and minimal C++ project manager tool written in Go. `CPPManager` simplifies creating, building, and managing C++ projects, enabling developers to focus on coding rather than repetitive setup tasks.

## Features

- Generate boilerplate C++ projects.
- Support for CMake-based build systems.
- Manage multiple project templates (simple and with unit tests).
- Automatically configure include paths and build folders.
- Easily integrate with GitHub for version control.
- Cross-platform support.

## Installation

1. Clone the repository:
   ```bash
   wget -O - https://raw.githubusercontent.com/tahaontech/cmng/refs/tags/v1/scripts/install.sh | sh
   ```

2. Verify installation:
   ```bash
   cmng version
   ```

## Usage

### 1. Create a New Project
Generate a new C++ project in the current directory:
```bash
cmng new
```
- **Interactive Setup**: Prompts for the project name and template selection.
- Available templates:
  - `simple`: Basic C++ project.
  - `with-test`: Includes Google Test setup for unit testing.

---

### 2. Build the Project
Build the C++ project:
```bash
cmng build
```
- Automatically runs `cmake` and `make` in the designated build folder.

---

### 3. Clean the Project
Remove all generated build files:
```bash
cmng clean
```

---

### 4. Run Unit Tests
Run the unit tests (if applicable):
```bash
cmng test
```

---

## Example Workflow

```bash
# Step 1: Create a new project
cmng new

# Step 2: Build the project
cmng build

# Step 3: Run the executable
cmng run # or ./build/my_project

# Step 4: Run tests (if the template includes tests)
cmng test
```

## Directory Structure

Here’s an example structure for a project created with `CPPManager`:

```plaintext
my_project/
├── CMakeLists.txt
├── include/
│   ├── my_class.h
├── src/
│   ├── main.cpp
│   ├── my_class.cpp
├── tests/  # Included only with `with-test` template
│   ├── test_my_class.cpp
├── build/
```

## Requirements

- GCC, Clang, or any compatible C++ compiler.
- CMake (version 3.10 or later).
- Go (for building this tool).

## Contribution

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature/new-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add new feature"
   ```
4. Push to your fork:
   ```bash
   git push origin feature/new-feature
   ```
5. Create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---

## Acknowledgments

- Inspired by the simplicity of other C++ build tools.
- Built with ❤️ using Go.

## Contact

For questions or suggestions, please open an issue or contact us at [thhk2831@gmail.com](mailto:thhk2831@gmail.com).
