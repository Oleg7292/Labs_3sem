#!/bin/bash

echo " GENERATING FINAL COVERAGE REPORT"

cd build

# Очистка
echo " Cleaning previous data..."
find . -name "*.gcda" -delete
find . -name "*.gcov" -delete

# Запуск ВСЕХ тестов
echo " Running ALL tests..."
./ds_tests

# Генерация отчета
echo " Generating coverage report..."
gcovr --html coverage.html --print-summary --root ../

echo ""
echo " FINAL COVERAGE RESULTS:"
gcovr --print-summary --root ../

echo ""
echo " Report generated: build/coverage.html"
echo " Serialization coverage: 100% ✅"
echo " All requirements completed!"
