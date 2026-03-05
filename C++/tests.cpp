#include <gtest/gtest.h>
#include <fstream>
#include <cstdio>
#include "data_structures.h"
#include "stack.h"
#include "array.h"
#include "singly_linked_list.h"
#include "doubly_linked_list.h"
#include "binary_tree.h"
#include "hash_table.h"
#include "serialization.h"

using namespace std;

// Вспомогательная функция для удаления тестовых файлов
void removeTestFile(const string& filename) {
    remove(filename.c_str());
}

// ==================== БАЗОВЫЕ ТЕСТЫ ====================

TEST(StackTest, Basic) {
    Stack stack(5);
    stack.push("test");
    EXPECT_EQ(stack.pop(), "test");
}

TEST(ArrayTest, Basic) {
    Array arr;
    arr.addToEnd("test");
    EXPECT_EQ(arr.getIndex(0), "test");
}

// ==================== ТЕСТЫ ДЛЯ ARRAY::TOVECTOR() ====================

TEST(ArrayToVectorTest, EmptyArrayReturnsEmptyVector) {
    Array emptyArray;
    vector<string> result = emptyArray.toVector();
    EXPECT_TRUE(result.empty());
    EXPECT_EQ(0, result.size());
}

TEST(ArrayToVectorTest, SingleElementArray) {
    Array arr;
    arr.addToEnd("hello");
    vector<string> result = arr.toVector();
    ASSERT_EQ(1, result.size());
    EXPECT_EQ("hello", result[0]);
}

TEST(ArrayToVectorTest, MultipleElementsArray) {
    Array arr;
    vector<string> expected = {"apple", "banana", "cherry", "date"};
    for (const auto& elem : expected) {
        arr.addToEnd(elem);
    }
    vector<string> result = arr.toVector();
    ASSERT_EQ(expected.size(), result.size());
    for (size_t i = 0; i < expected.size(); i++) {
        EXPECT_EQ(expected[i], result[i]) << "Mismatch at index " << i;
    }
}

TEST(ArrayToVectorTest, PreservesElementOrder) {
    Array arr;
    vector<string> input = {"first", "second", "third", "fourth"};
    for (const auto& elem : input) {
        arr.addToEnd(elem);
    }
    vector<string> result = arr.toVector();
    ASSERT_EQ(input.size(), result.size());
    for (size_t i = 0; i < input.size(); i++) {
        EXPECT_EQ(input[i], result[i]);
    }
}

TEST(ArrayToVectorTest, HandlesEmptyStrings) {
    Array arr;
    arr.addToEnd("");
    arr.addToEnd("non-empty");
    arr.addToEnd("");
    vector<string> result = arr.toVector();
    ASSERT_EQ(3, result.size());
    EXPECT_EQ("", result[0]);
    EXPECT_EQ("non-empty", result[1]);
    EXPECT_EQ("", result[2]);
}

TEST(ArrayToVectorTest, HandlesSpecialCharacters) {
    Array arr;
    vector<string> special = {
        "hello\nworld",
        "tab\there",
        "space here",
        "!@#$%^&*()",
        "unicode_тест",
        "very_long_string_that_might_cause_issues_but_should_be_handled_correctly_1234567890"
    };
    for (const auto& elem : special) {
        arr.addToEnd(elem);
    }
    vector<string> result = arr.toVector();
    ASSERT_EQ(special.size(), result.size());
    for (size_t i = 0; i < special.size(); i++) {
        EXPECT_EQ(special[i], result[i]);
    }
}

TEST(ArrayToVectorTest, DeepCopyIndependentOfOriginal) {
    Array arr;
    arr.addToEnd("original");
    vector<string> result = arr.toVector();
    arr.addToEnd("modified");
    ASSERT_EQ(1, result.size());
    EXPECT_EQ("original", result[0]);
    EXPECT_NE(arr.getSize(), result.size());
}

TEST(ArrayToVectorTest, LargeArrayPerformance) {
    const int LARGE_SIZE = 10000;
    Array arr;
    for (int i = 0; i < LARGE_SIZE; i++) {
        arr.addToEnd("element_" + to_string(i));
    }
    vector<string> result = arr.toVector();
    ASSERT_EQ(LARGE_SIZE, result.size());
    EXPECT_EQ("element_0", result[0]);
    EXPECT_EQ("element_" + to_string(LARGE_SIZE - 1), result[LARGE_SIZE - 1]);
}

TEST(ArrayToVectorTest, ConsistencyAfterMultipleConversions) {
    Array arr;
    arr.addToEnd("consistent");
    arr.addToEnd("data");
    vector<string> firstConversion = arr.toVector();
    vector<string> secondConversion = arr.toVector();
    EXPECT_EQ(firstConversion, secondConversion);
}

TEST(ArrayToVectorTest, ToVectorFromVectorCycle) {
    Array originalArr;
    vector<string> testData = {"cycle", "test", "data"};
    for (const auto& elem : testData) {
        originalArr.addToEnd(elem);
    }
    vector<string> serialized = originalArr.toVector();
    Array newArr;
    newArr.fromVector(serialized);
    vector<string> result = newArr.toVector();
    EXPECT_EQ(testData, result);
}

TEST(ArrayToVectorTest, NoMemoryLeaks) {
    Array* arr = new Array();
    arr->addToEnd("test1");
    arr->addToEnd("test2");
    arr->addToEnd("test3");
    vector<string> result = arr->toVector();
    delete arr;
    ASSERT_EQ(3, result.size());
    EXPECT_EQ("test1", result[0]);
    EXPECT_EQ("test2", result[1]);
    EXPECT_EQ("test3", result[2]);
}

// ==================== ТЕСТЫ СЕРИАЛИЗАЦИИ STACK ====================

TEST(SerializationStackTest, BinaryBasic) {
    Stack stack1(5);
    stack1.push("hello");
    stack1.push("world");

    Serializer::binarySerialize(stack1, "test_binary.dat");
    
    Stack stack2(5);
    Serializer::binaryDeserialize(stack2, "test_binary.dat");
    
    EXPECT_EQ(stack2.size(), 2);
    EXPECT_EQ(stack2.pop(), "world");
    EXPECT_EQ(stack2.pop(), "hello");
    
    removeTestFile("test_binary.dat");
}

TEST(SerializationStackTest, TextBasic) {
    Stack stack1(5);
    stack1.push("data1");
    stack1.push("data2");

    Serializer::textSerialize(stack1, "test_text.txt");
    
    Stack stack2(5);
    Serializer::textDeserialize(stack2, "test_text.txt");
    
    EXPECT_EQ(stack2.size(), 2);
    EXPECT_EQ(stack2.pop(), "data2");
    EXPECT_EQ(stack2.pop(), "data1");
    
    removeTestFile("test_text.txt");
}

TEST(SerializationStackTest, BinaryEmpty) {
    Stack stack1(5);
    Serializer::binarySerialize(stack1, "test_empty.dat");
    
    Stack stack2(5);
    Serializer::binaryDeserialize(stack2, "test_empty.dat");
    
    EXPECT_TRUE(stack2.isEmpty());
    
    removeTestFile("test_empty.dat");
}

TEST(SerializationStackTest, TextEmpty) {
    Stack stack1(5);
    Serializer::textSerialize(stack1, "test_empty.txt");
    
    Stack stack2(5);
    Serializer::textDeserialize(stack2, "test_empty.txt");
    
    EXPECT_TRUE(stack2.isEmpty());
    
    removeTestFile("test_empty.txt");
}

TEST(SerializationStackTest, BinaryMultiple) {
    Stack stack1(10);
    for (int i = 0; i < 5; i++) {
        stack1.push("item_" + to_string(i));
    }

    Serializer::binarySerialize(stack1, "test_multiple.dat");
    
    Stack stack2(10);
    Serializer::binaryDeserialize(stack2, "test_multiple.dat");
    
    EXPECT_EQ(stack2.size(), 5);
    for (int i = 4; i >= 0; i--) {
        EXPECT_EQ(stack2.pop(), "item_" + to_string(i));
    }
    
    removeTestFile("test_multiple.dat");
}

TEST(SerializationStackTest, TextMultiple) {
    Stack stack1(10);
    for (int i = 0; i < 3; i++) {
        stack1.push("elem_" + to_string(i));
    }

    Serializer::textSerialize(stack1, "test_multiple.txt");
    
    Stack stack2(10);
    Serializer::textDeserialize(stack2, "test_multiple.txt");
    
    EXPECT_EQ(stack2.size(), 3);
    for (int i = 2; i >= 0; i--) {
        EXPECT_EQ(stack2.pop(), "elem_" + to_string(i));
    }
    
    removeTestFile("test_multiple.txt");
}

TEST(SerializationStackTest, BinarySpecialChars) {
    Stack stack1(5);
    stack1.push("hello\nworld");
    stack1.push("tab\tdata");
    stack1.push("quote\"test");

    Serializer::binarySerialize(stack1, "test_special.dat");
    
    Stack stack2(5);
    Serializer::binaryDeserialize(stack2, "test_special.dat");
    
    EXPECT_EQ(stack2.size(), 3);
    EXPECT_EQ(stack2.pop(), "quote\"test");
    EXPECT_EQ(stack2.pop(), "tab\tdata");
    EXPECT_EQ(stack2.pop(), "hello\nworld");
    
    removeTestFile("test_special.dat");
}

// ==================== ТЕСТЫ СЕРИАЛИЗАЦИИ ARRAY ====================

TEST(SerializationArrayTest, BinaryBasic) {
    Array arr1;
    arr1.addToEnd("hello");
    arr1.addToEnd("world");

    Serializer::binarySerialize(arr1, "test_array_binary.dat");
    
    Array arr2;
    Serializer::binaryDeserialize(arr2, "test_array_binary.dat");
    
    vector<string> result = arr2.toVector();
    ASSERT_EQ(2, result.size());
    EXPECT_EQ("hello", result[0]);
    EXPECT_EQ("world", result[1]);
    
    removeTestFile("test_array_binary.dat");
}

TEST(SerializationArrayTest, TextBasic) {
    Array arr1;
    arr1.addToEnd("data1");
    arr1.addToEnd("data2");

    Serializer::textSerialize(arr1, "test_array_text.txt");
    
    Array arr2;
    Serializer::textDeserialize(arr2, "test_array_text.txt");
    
    vector<string> result = arr2.toVector();
    ASSERT_EQ(2, result.size());
    EXPECT_EQ("data1", result[0]);
    EXPECT_EQ("data2", result[1]);
    
    removeTestFile("test_array_text.txt");
}

TEST(SerializationArrayTest, BinaryEmpty) {
    Array arr1;
    Serializer::binarySerialize(arr1, "test_array_empty.dat");
    
    Array arr2;
    Serializer::binaryDeserialize(arr2, "test_array_empty.dat");
    
    EXPECT_TRUE(arr2.toVector().empty());
    
    removeTestFile("test_array_empty.dat");
}

TEST(SerializationArrayTest, TextEmpty) {
    Array arr1;
    Serializer::textSerialize(arr1, "test_array_empty.txt");
    
    Array arr2;
    Serializer::textDeserialize(arr2, "test_array_empty.txt");
    
    EXPECT_TRUE(arr2.toVector().empty());
    
    removeTestFile("test_array_empty.txt");
}

// ==================== ТЕСТЫ СЕРИАЛИЗАЦИИ HASH TABLE ====================

TEST(HashTableSerializationTest, BinaryBasic) {
    OpenAddressingHashTable ht1(10);
    ht1.insert('a', 100);
    ht1.insert('b', 200);
    ht1.insert('c', 300);
    
    Serializer::binarySerialize(ht1, "test_hash_binary.dat");
    
    OpenAddressingHashTable ht2;
    Serializer::binaryDeserialize(ht2, "test_hash_binary.dat");
    
    EXPECT_EQ(ht2.getSize(), 3);
    
    int value;
    EXPECT_TRUE(ht2.search('a', value));
    EXPECT_EQ(value, 100);
    EXPECT_TRUE(ht2.search('b', value));
    EXPECT_EQ(value, 200);
    EXPECT_TRUE(ht2.search('c', value));
    EXPECT_EQ(value, 300);
    
    removeTestFile("test_hash_binary.dat");
}

TEST(HashTableSerializationTest, TextBasic) {
    OpenAddressingHashTable ht1(5);
    ht1.insert('x', 10);
    ht1.insert('y', 20);
    ht1.insert('z', 30);
    
    Serializer::textSerialize(ht1, "test_hash_text.txt");
    
    OpenAddressingHashTable ht2;
    Serializer::textDeserialize(ht2, "test_hash_text.txt");
    
    EXPECT_EQ(ht2.getSize(), 3);
    
    int value;
    EXPECT_TRUE(ht2.search('x', value));
    EXPECT_EQ(value, 10);
    EXPECT_TRUE(ht2.search('y', value));
    EXPECT_EQ(value, 20);
    EXPECT_TRUE(ht2.search('z', value));
    EXPECT_EQ(value, 30);
    
    removeTestFile("test_hash_text.txt");
}

TEST(HashTableSerializationTest, BinaryEmpty) {
    OpenAddressingHashTable ht1;
    Serializer::binarySerialize(ht1, "test_hash_empty.dat");
    
    OpenAddressingHashTable ht2;
    Serializer::binaryDeserialize(ht2, "test_hash_empty.dat");
    
    EXPECT_TRUE(ht2.isEmpty());
    
    removeTestFile("test_hash_empty.dat");
}

TEST(HashTableSerializationTest, TextEmpty) {
    OpenAddressingHashTable ht1;
    Serializer::textSerialize(ht1, "test_hash_empty.txt");
    
    OpenAddressingHashTable ht2;
    Serializer::textDeserialize(ht2, "test_hash_empty.txt");
    
    EXPECT_TRUE(ht2.isEmpty());
    
    removeTestFile("test_hash_empty.txt");
}

// НОВЫЕ ТЕСТЫ ДЛЯ ПОКРЫТИЯ OpenAddressingHashTable::binaryDeserialize

TEST(HashTableSerializationTest, BinaryDeserializeFileTooSmall) {
    ofstream file("too_small.dat", ios::binary);
    char data[1] = {0};
    file.write(data, 1);
    file.close();
    
    OpenAddressingHashTable ht;
    EXPECT_THROW(Serializer::binaryDeserialize(ht, "too_small.dat"), runtime_error);
    
    removeTestFile("too_small.dat");
}

TEST(HashTableSerializationTest, BinaryDeserializeFailReadSize) {
    ofstream file("fail_size.dat", ios::binary);
    char data[2] = {0, 0};
    file.write(data, 2);
    file.close();
    
    OpenAddressingHashTable ht;
    EXPECT_THROW(Serializer::binaryDeserialize(ht, "fail_size.dat"), runtime_error);
    
    removeTestFile("fail_size.dat");
}

TEST(HashTableSerializationTest, BinaryDeserializeUnreasonableSize) {
    ofstream file("huge_size.dat", ios::binary);
    size_t hugeSize = 20000;
    file.write(reinterpret_cast<const char*>(&hugeSize), sizeof(hugeSize));
    file.close();
    
    OpenAddressingHashTable ht;
    EXPECT_THROW(Serializer::binaryDeserialize(ht, "huge_size.dat"), runtime_error);
    
    removeTestFile("huge_size.dat");
}

TEST(HashTableSerializationTest, BinaryDeserializeSizeMismatch) {
    ofstream file("mismatch.dat", ios::binary);
    size_t size = 5;
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    char key = 'a';
    int value = 100;
    file.write(&key, sizeof(char));
    file.write(reinterpret_cast<const char*>(&value), sizeof(int));
    file.close();
    
    OpenAddressingHashTable ht;
    EXPECT_THROW(Serializer::binaryDeserialize(ht, "mismatch.dat"), runtime_error);
    
    removeTestFile("mismatch.dat");
}

TEST(HashTableSerializationTest, BinaryDeserializeEofDuringRead) {
    ofstream file("eof_during.dat", ios::binary);
    size_t size = 2;
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    char key = 'a';
    file.write(&key, sizeof(char));
    file.close();
    
    OpenAddressingHashTable ht;
    EXPECT_THROW(Serializer::binaryDeserialize(ht, "eof_during.dat"), runtime_error);
    
    removeTestFile("eof_during.dat");
}

TEST(HashTableSerializationTest, BinaryDeserializeFailReadKey) {
    ofstream file("fail_key.dat", ios::binary);
    size_t size = 1;
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    file.close();
    
    OpenAddressingHashTable ht;
    EXPECT_THROW(Serializer::binaryDeserialize(ht, "fail_key.dat"), runtime_error);
    
    removeTestFile("fail_key.dat");
}

TEST(HashTableSerializationTest, BinaryDeserializeFailReadValue) {
    ofstream file("fail_value.dat", ios::binary);
    size_t size = 1;
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    char key = 'a';
    file.write(&key, sizeof(char));
    file.close();
    
    OpenAddressingHashTable ht;
    EXPECT_THROW(Serializer::binaryDeserialize(ht, "fail_value.dat"), runtime_error);
    
    removeTestFile("fail_value.dat");
}

TEST(HashTableSerializationTest, BinaryDeserializeExtraData) {
    ofstream file("extra_data.dat", ios::binary);
    size_t size = 2;
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    char key = 'a';
    int value = 100;
    file.write(&key, sizeof(char));
    file.write(reinterpret_cast<const char*>(&value), sizeof(int));
    
    key = 'b';
    value = 200;
    file.write(&key, sizeof(char));
    file.write(reinterpret_cast<const char*>(&value), sizeof(int));
    
    char extra = 'X';
    file.write(&extra, sizeof(char));
    file.close();
    
    OpenAddressingHashTable ht;
    testing::internal::CaptureStderr();
    EXPECT_NO_THROW(Serializer::binaryDeserialize(ht, "extra_data.dat"));
    string output = testing::internal::GetCapturedStderr();
    EXPECT_TRUE(output.find("Warning: Extra data") != string::npos);
    
    int val;
    EXPECT_TRUE(ht.search('a', val));
    EXPECT_EQ(100, val);
    EXPECT_TRUE(ht.search('b', val));
    EXPECT_EQ(200, val);
    
    removeTestFile("extra_data.dat");
}

// ==================== ТЕСТЫ СЕРИАЛИЗАЦИИ SINGLY LINKED LIST ====================

TEST(SerializationSinglyLinkedListTest, BinaryBasic) {
    SinglyLinkedList list1;
    list1.push_back("first");
    list1.push_back("second");
    list1.push_back("third");

    Serializer::binarySerialize(list1, "test_sll_binary.dat");
    
    SinglyLinkedList list2;
    Serializer::binaryDeserialize(list2, "test_sll_binary.dat");
    
    vector<string> result = list2.toVector();
    ASSERT_EQ(3, result.size());
    EXPECT_EQ("first", result[0]);
    EXPECT_EQ("second", result[1]);
    EXPECT_EQ("third", result[2]);
    
    removeTestFile("test_sll_binary.dat");
}

TEST(SerializationSinglyLinkedListTest, TextBasic) {
    SinglyLinkedList list1;
    list1.push_back("one");
    list1.push_back("two");
    list1.push_back("three");

    Serializer::textSerialize(list1, "test_sll_text.txt");
    
    SinglyLinkedList list2;
    Serializer::textDeserialize(list2, "test_sll_text.txt");
    
    vector<string> result = list2.toVector();
    ASSERT_EQ(3, result.size());
    EXPECT_EQ("one", result[0]);
    EXPECT_EQ("two", result[1]);
    EXPECT_EQ("three", result[2]);
    
    removeTestFile("test_sll_text.txt");
}

TEST(SerializationSinglyLinkedListTest, BinaryEmpty) {
    SinglyLinkedList list1;
    Serializer::binarySerialize(list1, "test_sll_empty.dat");
    
    SinglyLinkedList list2;
    Serializer::binaryDeserialize(list2, "test_sll_empty.dat");
    
    EXPECT_TRUE(list2.toVector().empty());
    
    removeTestFile("test_sll_empty.dat");
}

TEST(SerializationSinglyLinkedListTest, TextEmpty) {
    SinglyLinkedList list1;
    Serializer::textSerialize(list1, "test_sll_empty.txt");
    
    SinglyLinkedList list2;
    Serializer::textDeserialize(list2, "test_sll_empty.txt");
    
    EXPECT_TRUE(list2.toVector().empty());
    
    removeTestFile("test_sll_empty.txt");
}

// НОВЫЕ ТЕСТЫ ДЛЯ ПОКРЫТИЯ SinglyLinkedList

TEST(SinglyLinkedListSerializationTest, BinarySerializeNormal) {
    SinglyLinkedList list;
    list.push_back("one");
    list.push_back("two");
    list.push_back("three");
    
    EXPECT_NO_THROW(Serializer::binarySerialize(list, "sll_binary_normal.dat"));
    removeTestFile("sll_binary_normal.dat");
}

TEST(SinglyLinkedListSerializationTest, BinarySerializeEmpty) {
    SinglyLinkedList list;
    EXPECT_NO_THROW(Serializer::binarySerialize(list, "sll_binary_empty.dat"));
    removeTestFile("sll_binary_empty.dat");
}

TEST(SinglyLinkedListSerializationTest, BinaryDeserializeFileNotFound) {
    SinglyLinkedList list;
    EXPECT_THROW(Serializer::binaryDeserialize(list, "nonexistent.dat"), runtime_error);
}

TEST(SinglyLinkedListSerializationTest, TextSerializeNormal) {
    SinglyLinkedList list;
    list.push_back("one");
    list.push_back("two");
    
    EXPECT_NO_THROW(Serializer::textSerialize(list, "sll_text_normal.txt"));
    removeTestFile("sll_text_normal.txt");
}

TEST(SinglyLinkedListSerializationTest, TextSerializeEmpty) {
    SinglyLinkedList list;
    EXPECT_NO_THROW(Serializer::textSerialize(list, "sll_text_empty.txt"));
    removeTestFile("sll_text_empty.txt");
}

TEST(SinglyLinkedListSerializationTest, TextDeserializeFileNotFound) {
    SinglyLinkedList list;
    EXPECT_THROW(Serializer::textDeserialize(list, "nonexistent.txt"), runtime_error);
}

// ==================== ТЕСТЫ СЕРИАЛИЗАЦИИ DOUBLY LINKED LIST ====================

TEST(SerializationDoublyLinkedListTest, BinaryBasic) {
    DoublyLinkedList list1;
    list1.push_back("first");
    list1.push_back("second");
    list1.push_back("third");

    Serializer::binarySerialize(list1, "test_dll_binary.dat");
    
    DoublyLinkedList list2;
    Serializer::binaryDeserialize(list2, "test_dll_binary.dat");
    
    vector<string> result = list2.toVector();
    ASSERT_EQ(3, result.size());
    EXPECT_EQ("first", result[0]);
    EXPECT_EQ("second", result[1]);
    EXPECT_EQ("third", result[2]);
    
    removeTestFile("test_dll_binary.dat");
}

TEST(SerializationDoublyLinkedListTest, TextBasic) {
    DoublyLinkedList list1;
    list1.push_back("one");
    list1.push_back("two");
    list1.push_back("three");

    Serializer::textSerialize(list1, "test_dll_text.txt");
    
    DoublyLinkedList list2;
    Serializer::textDeserialize(list2, "test_dll_text.txt");
    
    vector<string> result = list2.toVector();
    ASSERT_EQ(3, result.size());
    EXPECT_EQ("one", result[0]);
    EXPECT_EQ("two", result[1]);
    EXPECT_EQ("three", result[2]);
    
    removeTestFile("test_dll_text.txt");
}

TEST(SerializationDoublyLinkedListTest, BinaryEmpty) {
    DoublyLinkedList list1;
    Serializer::binarySerialize(list1, "test_dll_empty.dat");
    
    DoublyLinkedList list2;
    Serializer::binaryDeserialize(list2, "test_dll_empty.dat");
    
    EXPECT_TRUE(list2.toVector().empty());
    
    removeTestFile("test_dll_empty.dat");
}

TEST(SerializationDoublyLinkedListTest, TextEmpty) {
    DoublyLinkedList list1;
    Serializer::textSerialize(list1, "test_dll_empty.txt");
    
    DoublyLinkedList list2;
    Serializer::textDeserialize(list2, "test_dll_empty.txt");
    
    EXPECT_TRUE(list2.toVector().empty());
    
    removeTestFile("test_dll_empty.txt");
}

// НОВЫЕ ТЕСТЫ ДЛЯ ПОКРЫТИЯ DoublyLinkedList

TEST(DoublyLinkedListSerializationTest, BinarySerializeNormal) {
    DoublyLinkedList list;
    list.push_back("one");
    list.push_back("two");
    list.push_back("three");
    
    EXPECT_NO_THROW(Serializer::binarySerialize(list, "dll_binary_normal.dat"));
    removeTestFile("dll_binary_normal.dat");
}

TEST(DoublyLinkedListSerializationTest, BinarySerializeEmpty) {
    DoublyLinkedList list;
    EXPECT_NO_THROW(Serializer::binarySerialize(list, "dll_binary_empty.dat"));
    removeTestFile("dll_binary_empty.dat");
}

TEST(DoublyLinkedListSerializationTest, BinaryDeserializeFileNotFound) {
    DoublyLinkedList list;
    EXPECT_THROW(Serializer::binaryDeserialize(list, "nonexistent.dat"), runtime_error);
}

TEST(DoublyLinkedListSerializationTest, TextSerializeNormal) {
    DoublyLinkedList list;
    list.push_back("one");
    list.push_back("two");
    
    EXPECT_NO_THROW(Serializer::textSerialize(list, "dll_text_normal.txt"));
    removeTestFile("dll_text_normal.txt");
}

TEST(DoublyLinkedListSerializationTest, TextSerializeEmpty) {
    DoublyLinkedList list;
    EXPECT_NO_THROW(Serializer::textSerialize(list, "dll_text_empty.txt"));
    removeTestFile("dll_text_empty.txt");
}

TEST(DoublyLinkedListSerializationTest, TextDeserializeFileNotFound) {
    DoublyLinkedList list;
    EXPECT_THROW(Serializer::textDeserialize(list, "nonexistent.txt"), runtime_error);
}

// ==================== ДОПОЛНИТЕЛЬНЫЕ ТЕСТЫ ====================

TEST(SerializationErrorTest, FileNotFound) {
    Stack stack(5);
    EXPECT_THROW(Serializer::binaryDeserialize(stack, "nonexistent_file.dat"), runtime_error);
    EXPECT_THROW(Serializer::textDeserialize(stack, "nonexistent_file.txt"), runtime_error);
}

TEST(SerializationComplexTest, MultipleStacks) {
    const int NUM_STACKS = 3;
    vector<string> filenames;
    
    // Сериализация нескольких стеков
    for (int i = 0; i < NUM_STACKS; i++) {
        Stack stack(10);  // Создаем на стеке, не в куче
        for (int j = 0; j < 5; j++) {
            stack.push("stack" + to_string(i) + "_item" + to_string(j));
        }
        
        string filename = "stack_" + to_string(i) + ".dat";
        filenames.push_back(filename);
        Serializer::binarySerialize(stack, filename);
    }
    
    // Десериализация и проверка
    for (int i = 0; i < NUM_STACKS; i++) {
        Stack deserializedStack(10);
        Serializer::binaryDeserialize(deserializedStack, filenames[i]);
        
        EXPECT_EQ(deserializedStack.size(), 5);
        
        // Проверяем в обратном порядке (LIFO)
        for (int j = 4; j >= 0; j--) {
            EXPECT_EQ(deserializedStack.pop(), "stack" + to_string(i) + "_item" + to_string(j));
        }
        
        // Удаляем файл после использования
        removeTestFile(filenames[i]);
    }
}

// НОВЫЙ ТЕСТ ДЛЯ ПОЛНОГО ПОКРЫТИЯ ВСЕХ ФУНКЦИЙ
TEST(SerializationCompleteTest, AllFunctions) {
    // Stack
    Stack stack(5);
    stack.push("test");
    Serializer::binarySerialize(stack, "all_stack.dat");
    Stack stackDeserialized(5);
    Serializer::binaryDeserialize(stackDeserialized, "all_stack.dat");
    removeTestFile("all_stack.dat");
    
    // Array
    Array array;
    array.addToEnd("test");
    Serializer::binarySerialize(array, "all_array.dat");
    Array arrayDeserialized;
    Serializer::binaryDeserialize(arrayDeserialized, "all_array.dat");
    removeTestFile("all_array.dat");
    
    // HashTable
    OpenAddressingHashTable hash;
    hash.insert('t', 1);
    Serializer::binarySerialize(hash, "all_hash.dat");
    OpenAddressingHashTable hashDeserialized;
    Serializer::binaryDeserialize(hashDeserialized, "all_hash.dat");
    removeTestFile("all_hash.dat");
    
    // SinglyLinkedList
    SinglyLinkedList sll;
    sll.push_back("test");
    Serializer::binarySerialize(sll, "all_sll.dat");
    SinglyLinkedList sllDeserialized;
    Serializer::binaryDeserialize(sllDeserialized, "all_sll.dat");
    removeTestFile("all_sll.dat");
    
    // DoublyLinkedList
    DoublyLinkedList dll;
    dll.push_back("test");
    Serializer::binarySerialize(dll, "all_dll.dat");
    DoublyLinkedList dllDeserialized;
    Serializer::binaryDeserialize(dllDeserialized, "all_dll.dat");
    removeTestFile("all_dll.dat");
    
    SUCCEED();
}

// Array дополнительные тесты
TEST(ArrayTest, Comprehensive) {
    Array arr;
    
    for (int i = 0; i < 20; i++) {
        arr.addToEnd("item_" + to_string(i));
    }
    EXPECT_EQ(arr.getSize(), 20);
    
    arr.add(5, "inserted");
    EXPECT_EQ(arr.getIndex(5), "inserted");
    
    arr.replace(10, "replaced");
    EXPECT_EQ(arr.getIndex(10), "replaced");
    
    arr.remove(7);
    EXPECT_EQ(arr.getSize(), 20);
}

// BinaryTree дополнительные тесты
TEST(FullBinaryTreeTest, Comprehensive) {
    FullBinaryTree tree;
    
    vector<int> values = {50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45};
    for (int val : values) {
        tree.insert(val);
    }
    
    vector<int> elements = tree.toVector();
    EXPECT_EQ(elements.size(), values.size());
    EXPECT_TRUE(is_sorted(elements.begin(), elements.end()));
}

// SinglyLinkedList дополнительные тесты
TEST(SinglyLinkedListTest, Comprehensive) {
    SinglyLinkedList list;
    
    list.push_back("a");
    list.push_front("start");
    list.push_back("b");
    list.insert(2, "middle");
    list.push_back("end");
    
    EXPECT_EQ(list.size(), 5);
    EXPECT_EQ(list.get(0), "start");
    EXPECT_EQ(list.get(2), "middle");
    EXPECT_EQ(list.get(4), "end");
    
    list.clear();
    EXPECT_TRUE(list.isEmpty());
}

// DoublyLinkedList дополнительные тесты
TEST(DoublyLinkedListTest, Comprehensive) {
    DoublyLinkedList list;
    
    list.push_back("first");
    list.push_front("very_first");
    list.push_back("last");
    list.insert(2, "middle");
    
    EXPECT_EQ(list.size(), 4);
    EXPECT_EQ(list.get_front(), "very_first");
    EXPECT_EQ(list.get_back(), "last");
    EXPECT_EQ(list.get(2), "middle");
    
    list.pop_front();
    EXPECT_EQ(list.get_front(), "first");
    
    list.pop_back();
    EXPECT_EQ(list.get_back(), "middle");
}

// HashTable дополнительные тесты
TEST(HashTableTest, Comprehensive) {
    OpenAddressingHashTable ht(50);
    
    for (char c = 'a'; c <= 'z'; c++) {
        ht.insert(c, int(c));
    }
    
    int value;
    for (char c = 'a'; c <= 'z'; c++) {
        EXPECT_TRUE(ht.search(c, value));
        EXPECT_EQ(value, int(c));
    }
    
    vector<pair<char, int>> data = ht.toVector();
    EXPECT_FALSE(data.empty());
}

// Stack дополнительные тесты
TEST(StackTest, Comprehensive) {
    Stack stack(10);
    
    for (int i = 0; i < 8; i++) {
        stack.push("val_" + to_string(i));
    }
    
    EXPECT_EQ(stack.size(), 8);
    
    vector<string> data = stack.toVector();
    Stack stack2(10);
    stack2.fromVector(data);
    
    EXPECT_EQ(stack2.size(), 8);
    for (int i = 7; i >= 0; i--) {
        EXPECT_EQ(stack2.pop(), "val_" + to_string(i));
    }
}

// ==================== ФИНАЛЬНЫЕ ТЕСТЫ ====================

TEST(FullBinaryTreeTest, PrintCoverage) {
    FullBinaryTree tree;
    
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
    
    tree.insert(5);
    tree.insert(3);
    tree.insert(7);
    
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
}

TEST(ArrayTest, EdgeCases) {
    Array arr;
    
    EXPECT_NO_THROW(arr.ShowArray());
    EXPECT_EQ(arr.getSize(), 0);
    
    arr.add(0, "first");
    EXPECT_EQ(arr.getSize(), 1);
    EXPECT_EQ(arr.getIndex(0), "first");
    
    vector<string> empty_vec;
    arr.fromVector(empty_vec);
    EXPECT_EQ(arr.getSize(), 0);
}

TEST(HashTableTest, EdgeCases) {
    OpenAddressingHashTable ht(5);
    
    ht.clear();
    EXPECT_TRUE(ht.isEmpty());
    
    int value;
    EXPECT_FALSE(ht.search('x', value));
    EXPECT_FALSE(ht.remove('x'));
    
    ht.insert('a', 1);
    ht.insert('b', 2);
    ht.insert('c', 3);
    ht.insert('d', 4);
    ht.insert('e', 5);
    ht.insert('f', 6);
    
    vector<pair<char, int>> data = ht.toVector();
    OpenAddressingHashTable ht2;
    ht2.fromVector(data);
    
    EXPECT_TRUE(ht2.search('a', value));
    EXPECT_EQ(value, 1);
}

TEST(SinglyLinkedListTest, EdgeCases) {
    SinglyLinkedList list;
    
    vector<string> empty = list.toVector();
    EXPECT_TRUE(empty.empty());
    
    vector<string> data = {"from", "vector", "test"};
    list.fromVector(data);
    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.get(1), "vector");
    
    SinglyLinkedList empty_list;
    EXPECT_NO_THROW(empty_list.remove(0));
    EXPECT_NO_THROW(empty_list.pop_front());
}

TEST(DoublyLinkedListTest, EdgeCases) {
    DoublyLinkedList list;
    
    vector<string> empty = list.toVector();
    EXPECT_TRUE(empty.empty());
    
    vector<string> data = {"p", "q", "r"};
    list.fromVector(data);
    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.get_front(), "p");
    EXPECT_EQ(list.get_back(), "r");
    
    EXPECT_THROW(list.get(10), out_of_range);
    EXPECT_THROW(list.get(-1), out_of_range);
}

TEST(StackTest, EdgeCases) {
    Stack stack(2);
    
    vector<string> empty = stack.toVector();
    EXPECT_TRUE(empty.empty());
    
    stack.fromVector(empty);
    EXPECT_TRUE(stack.isEmpty());
    
    vector<string> data = {"cycle", "test"};
    stack.fromVector(data);
    vector<string> result = stack.toVector();
    EXPECT_EQ(data, result);
}

TEST(DoublyLinkedListTest, FullCoverage) {
    DoublyLinkedList list;
    
    list.push_front("start");
    list.push_back("end");
    list.insert(1, "middle1");
    list.insert(2, "middle2");
    list.push_front("new_start");
    list.push_back("new_end");
    
    list.pop_front();
    list.pop_back();
    list.remove(2);
    
    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.get_front(), "start");
    EXPECT_EQ(list.get_back(), "end");
    EXPECT_EQ(list.get(1), "middle1");
}

TEST(SinglyLinkedListTest, FullCoverage) {
    SinglyLinkedList list;
    
    list.push_back("a");
    list.push_front("before_a");
    list.push_back("c");
    list.insert(2, "b");
    list.push_front("very_first");
    list.push_back("very_last");
    
    EXPECT_EQ(list.size(), 6);
    
    list.pop_front();
    list.remove(3);
    list.remove(0);
    
    while (!list.isEmpty()) {
        list.pop_front();
    }
    EXPECT_TRUE(list.isEmpty());
    
    vector<string> big_data;
    for (int i = 0; i < 20; i++) {
        big_data.push_back("big_" + to_string(i));
    }
    list.fromVector(big_data);
    EXPECT_EQ(list.size(), 20);
}

TEST(StackTest, FullCoverage) {
    Stack stack(5);
    
    stack.push("first");
    stack.push("second");
    stack.push("third");
    
    EXPECT_EQ(stack.size(), 3);
    EXPECT_EQ(stack.peek(), "third");
    
    vector<string> data1 = stack.toVector();
    Stack stack2(5);
    stack2.fromVector(data1);
    
    EXPECT_EQ(stack2.size(), 3);
    EXPECT_EQ(stack2.peek(), "third");
    
    EXPECT_EQ(stack2.pop(), "third");
    EXPECT_EQ(stack2.pop(), "second");
    EXPECT_EQ(stack2.pop(), "first");
    EXPECT_TRUE(stack2.isEmpty());
    
    Stack empty_stack(3);
    vector<string> empty_data = empty_stack.toVector();
    EXPECT_TRUE(empty_data.empty());
    
    empty_stack.fromVector(empty_data);
    EXPECT_TRUE(empty_stack.isEmpty());
    
    Stack small_stack(2);
    small_stack.push("one");
    small_stack.push("two");
    EXPECT_THROW(small_stack.push("three"), overflow_error);
}

TEST(FullBinaryTreeTest, PrintAndEdgeCases) {
    FullBinaryTree tree;
    
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
    
    tree.insert(42);
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
    
    tree.insert(20);
    tree.insert(60);
    tree.insert(10);
    tree.insert(30);
    tree.insert(50);
    tree.insert(70);
    
    EXPECT_FALSE(tree.isEmpty());
    vector<int> elements = tree.toVector();
    EXPECT_EQ(elements.size(), 7);
    EXPECT_TRUE(is_sorted(elements.begin(), elements.end()));
    
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
}

TEST(HashTableTest, StressAndEdgeCases) {
    OpenAddressingHashTable ht(50);
    
    for (int i = 0; i < 26; i++) {
        char key = 'a' + i;
        ht.insert(key, i * 10);
    }
    
    int value;
    for (int i = 0; i < 26; i++) {
        char key = 'a' + i;
        EXPECT_TRUE(ht.search(key, value));
        EXPECT_EQ(value, i * 10);
    }
    
    ht.insert('a', 999);
    EXPECT_TRUE(ht.search('a', value));
    EXPECT_EQ(value, 999);
    
    EXPECT_TRUE(ht.remove('b'));
    EXPECT_FALSE(ht.search('b', value));
    
    ht.insert('b', 888);
    EXPECT_TRUE(ht.search('b', value));
    EXPECT_EQ(value, 888);
    
    vector<pair<char, int>> data = ht.toVector();
    OpenAddressingHashTable ht2;
    ht2.fromVector(data);
    
    EXPECT_TRUE(ht2.search('z', value));
    EXPECT_EQ(value, 25 * 10);
}

int main(int argc, char** argv) {
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}