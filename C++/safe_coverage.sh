#!/bin/bash

echo "=== SAFE COVERAGE SETUP ==="

# Очистка ТОЛЬКО build файлов, не исходников
echo " Cleaning build files..."
find . -maxdepth 1 -type f \( -name "*.gcda" -o -name "*.gcno" -o -name "ds_tests" -o -name "ds_demo" -o -name "coverage.html" \) -delete
rm -rf CMakeFiles/ CMakeCache.txt Makefile cmake_install.cmake

# Конфигурация с coverage
echo " Configuring with coverage..."
cmake .. -DCMAKE_CXX_FLAGS="-fprofile-arcs -ftest-coverage" -DCMAKE_EXE_LINKER_FLAGS="-fprofile-arcs -ftest-coverage"

# Сборка тестов
echo "🔨 Building tests..."
make ds_tests

# Проверка что собралось
echo " Checking build results..."
ls -la ds_tests

# Запуск тестов для генерации coverage данных
echo " Running tests..."
./ds_tests

# Проверка coverage файлов
echo " Checking coverage files..."
find . -name "*.gcda" | head -5

# Генерация отчета
echo " Generating coverage report..."
gcovr --html coverage.html --print-summary

echo " DONE! Coverage report: coverage.html"
echo " To view: open coverage.html"
