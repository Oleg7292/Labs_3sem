
#!/bin/bash

echo "=== FULL COVERAGE SETUP ==="

# 1. Очистка
echo " Step 1: Cleaning build directory..."
rm -rf *

# 2. Конфигурация с coverage флагами
echo " Step 2: Configuring with coverage flags..."
cmake .. -DCMAKE_CXX_FLAGS="-fprofile-arcs -ftest-coverage" -DCMAKE_EXE_LINKER_FLAGS="-fprofile-arcs -ftest-coverage"

# 3. Сборка
echo " Step 3: Building tests..."
make ds_tests

# 4. Проверка файлов coverage ДО запуска тестов
echo " Step 4: Checking coverage files before tests..."
find . -name "*.gcno" | head -5

# 5. Запуск тестов
echo " Step 5: Running tests..."
./ds_tests

# 6. Проверка файлов coverage ПОСЛЕ запуска тестов
echo " Step 6: Checking coverage files after tests..."
find . -name "*.gcda" | head -5

# 7. Генерация отчетов
echo " Step 7: Generating coverage reports..."

echo "--- TEXT REPORT ---"
gcovr --print-summary

echo "--- HTML REPORT ---"
gcovr --html coverage.html --print-summary --html-details

echo " DONE! Coverage report: coverage.html"
echo " To view: open coverage.html"
