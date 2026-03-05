#include "serialization.h"

void Serializer::binarySerialize(const Stack& obj, const string& filename) {
    ofstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<string> data = obj.toVector();
    size_t size = data.size();
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    for (const auto& str : data) {
        size_t strSize = str.size();
        file.write(reinterpret_cast<const char*>(&strSize), sizeof(strSize));
        file.write(str.c_str(), strSize);
    }
    file.close();
}

void Serializer::binaryDeserialize(Stack& obj, const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    
    vector<string> data;
    for (size_t i = 0; i < size; ++i) {
        size_t strSize;
        file.read(reinterpret_cast<char*>(&strSize), sizeof(strSize));
        
        string str(strSize, '\0');
        file.read(&str[0], strSize);
        data.push_back(str);
    }
    file.close();
    
    obj.fromVector(data);
}

// Текстовая сериализация Stack
void Serializer::textSerialize(const Stack& obj, const string& filename) {
    ofstream file(filename);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<string> data = obj.toVector();
    file << data.size() << endl;
    for (const auto& str : data) {
        file << str << endl;
    }
    file.close();
}

void Serializer::textDeserialize(Stack& obj, const string& filename) {
    ifstream file(filename);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file >> size;
    file.ignore();
    
    vector<string> data;
    for (size_t i = 0; i < size; ++i) {
        string str;
        getline(file, str);
        data.push_back(str);
    }
    file.close();
    
    obj.fromVector(data);
}

void Serializer::binarySerialize(const Array& obj, const string& filename) {
    ofstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<string> data = obj.toVector();
    size_t size = data.size();
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    for (const auto& str : data) {
        size_t strSize = str.size();
        file.write(reinterpret_cast<const char*>(&strSize), sizeof(strSize));
        file.write(str.c_str(), strSize);
    }
    file.close();
}

void Serializer::binaryDeserialize(Array& obj, const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    
    vector<string> data;
    for (size_t i = 0; i < size; ++i) {
        size_t strSize;
        file.read(reinterpret_cast<char*>(&strSize), sizeof(strSize));
        
        string str(strSize, '\0');
        file.read(&str[0], strSize);
        data.push_back(str);
    }
    file.close();
    
    obj.fromVector(data);
}

void Serializer::textSerialize(const Array& obj, const string& filename) {
    ofstream file(filename);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<string> data = obj.toVector();
    file << data.size() << endl;
    for (const auto& str : data) {
        file << str << endl;
    }
    file.close();
}

void Serializer::textDeserialize(Array& obj, const string& filename) {
    ifstream file(filename);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file >> size;
    file.ignore();
    
    vector<string> data;
    for (size_t i = 0; i < size; ++i) {
        string str;
        getline(file, str);
        data.push_back(str);
    }
    file.close();
    
    obj.fromVector(data);
}

// =============== HASH TABLE ===============

void Serializer::binarySerialize(const OpenAddressingHashTable& obj, const string& filename) {
    ofstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<pair<char, int>> data = obj.toVector();
    size_t size = data.size();
    
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    for (const auto& item : data) {
        file.write(reinterpret_cast<const char*>(&item.first), sizeof(char));
        file.write(reinterpret_cast<const char*>(&item.second), sizeof(int));
    }
    
    file.close();
}

void Serializer::binaryDeserialize(OpenAddressingHashTable& obj, const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    // Получаем размер файла для проверки
    file.seekg(0, ios::end);
    streampos fileSize = file.tellg();
    file.seekg(0, ios::beg);
    
    // Проверяем минимальный размер файла
    if (fileSize < static_cast<streampos>(sizeof(size_t))) {
        file.close();
        throw runtime_error("File is too small or corrupted");
    }
    
    size_t size;
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    
    // Проверка на успешное чтение
    if (file.fail()) {
        file.close();
        throw runtime_error("Failed to read size information");
    }
    
    // Защита от слишком большого размера (предотвращение DoS)
    const size_t MAX_REASONABLE_SIZE = 10000;
    if (size > MAX_REASONABLE_SIZE) {
        file.close();
        throw runtime_error("Size seems unreasonably large, file may be corrupted");
    }
    
    // Проверка соответствия размера файла
    streampos expectedSize = static_cast<streampos>(sizeof(size) + size * (sizeof(char) + sizeof(int)));
    if (fileSize < expectedSize) {
        file.close();
        throw runtime_error("File size mismatch: expected more data");
    }
    
    obj.clear();
    
    // Чтение данных с проверками
    for (size_t i = 0; i < size; ++i) {
        // Проверка, не достигли ли конца файла преждевременно
        if (file.eof()) {
            file.close();
            throw runtime_error("Unexpected end of file at position " + to_string(i));
        }
        
        char key;
        file.read(reinterpret_cast<char*>(&key), sizeof(char));
        
        if (file.fail()) {
            file.close();
            throw runtime_error("Failed to read key at position " + to_string(i));
        }
        
        int value;
        file.read(reinterpret_cast<char*>(&value), sizeof(int));
        
        if (file.fail()) {
            file.close();
            throw runtime_error("Failed to read value at position " + to_string(i));
        }
        
        // Попытка вставки с проверкой
        try {
            obj.insert(key, value);
        } catch (const exception& e) {
            file.close();
            throw runtime_error("Failed to insert element: " + string(e.what()));
        }
    }
    
    // Проверка на лишние данные
    file.peek();
    if (!file.eof()) {
        // Это предупреждение, не ошибка
        cerr << "Warning: Extra data found in file after deserialization" << endl;
    }
    
    file.close();
}

void Serializer::textSerialize(const OpenAddressingHashTable& obj, const string& filename) {
    ofstream file(filename);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<pair<char, int>> data = obj.toVector();
    
    file << data.size() << endl;
    
    for (const auto& item : data) {
        file << item.first << " " << item.second << endl;
    }
    
    file.close();
}

void Serializer::textDeserialize(OpenAddressingHashTable& obj, const string& filename) {
    ifstream file(filename);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file >> size;
    
    obj.clear();
    
    for (size_t i = 0; i < size; ++i) {
        char key;
        int value;
        
        file >> key >> value;
        obj.insert(key, value);
    }
    
    file.close();
}

// =============== SINGLY LINKED LIST ===============

void Serializer::binarySerialize(const SinglyLinkedList& obj, const string& filename) {
    ofstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<string> data = obj.toVector();
    size_t size = data.size();
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    for (const auto& str : data) {
        size_t strSize = str.size();
        file.write(reinterpret_cast<const char*>(&strSize), sizeof(strSize));
        file.write(str.c_str(), strSize);
    }
    file.close();
}

void Serializer::binaryDeserialize(SinglyLinkedList& obj, const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    
    vector<string> data;
    for (size_t i = 0; i < size; ++i) {
        size_t strSize;
        file.read(reinterpret_cast<char*>(&strSize), sizeof(strSize));
        
        string str(strSize, '\0');
        file.read(&str[0], strSize);
        data.push_back(str);
    }
    file.close();
    
    obj.fromVector(data);
}

void Serializer::textSerialize(const SinglyLinkedList& obj, const string& filename) {
    ofstream file(filename);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<string> data = obj.toVector();
    file << data.size() << endl;
    for (const auto& str : data) {
        file << str << endl;
    }
    file.close();
}

void Serializer::textDeserialize(SinglyLinkedList& obj, const string& filename) {
    ifstream file(filename);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file >> size;
    file.ignore();
    
    vector<string> data;
    for (size_t i = 0; i < size; ++i) {
        string str;
        getline(file, str);
        data.push_back(str);
    }
    file.close();
    
    obj.fromVector(data);
}

// =============== DOUBLY LINKED LIST ===============

void Serializer::binarySerialize(const DoublyLinkedList& obj, const string& filename) {
    ofstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<string> data = obj.toVector();
    size_t size = data.size();
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    for (const auto& str : data) {
        size_t strSize = str.size();
        file.write(reinterpret_cast<const char*>(&strSize), sizeof(strSize));
        file.write(str.c_str(), strSize);
    }
    file.close();
}

void Serializer::binaryDeserialize(DoublyLinkedList& obj, const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    
    vector<string> data;
    for (size_t i = 0; i < size; ++i) {
        size_t strSize;
        file.read(reinterpret_cast<char*>(&strSize), sizeof(strSize));
        
        string str(strSize, '\0');
        file.read(&str[0], strSize);
        data.push_back(str);
    }
    file.close();
    
    obj.fromVector(data);
}

void Serializer::textSerialize(const DoublyLinkedList& obj, const string& filename) {
    ofstream file(filename);
    if (!file) throw runtime_error("Cannot open file for writing");
    
    vector<string> data = obj.toVector();
    file << data.size() << endl;
    for (const auto& str : data) {
        file << str << endl;
    }
    file.close();
}

void Serializer::textDeserialize(DoublyLinkedList& obj, const string& filename) {
    ifstream file(filename);
    if (!file) throw runtime_error("Cannot open file for reading");
    
    size_t size;
    file >> size;
    file.ignore();
    
    vector<string> data;
    for (size_t i = 0; i < size; ++i) {
        string str;
        getline(file, str);
        data.push_back(str);
    }
    file.close();
    
    obj.fromVector(data);
}