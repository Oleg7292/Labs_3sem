#include <iostream>
#include <string>
#include <vector>
#include <limits>
#include <memory>

using namespace std;

// Базовый класс для всех устройств
class Device {
public:
    // Конструктор для инициализации общих свойств
    Device(const string& partNumber, const string& brand, double price)
        : partNumber(partNumber), brand(brand), price(price) {}

    // Чисто виртуальная функция для вывода информации
    virtual void displayInfo() const = 0;

    // Виртуальный деструктор
    virtual ~Device() = default;

    // Вывод краткой информации (номер и бренд)
    void displayShortInfo() const {
        cout << partNumber << ", " << brand << endl;
    }

protected:
    string partNumber; // Номенклатурный номер
    string brand;      // Бренд
    double price;      // Стоимость
};


// КЛАССЫ КОНКРЕТНЫХ УСТРОЙСТВ


// Класс для наушников
class Headphones : public Device {
public:
    Headphones(const string& partNumber, const string& brand, double price,
               const string& designType, const string& mountingType)
        : Device(partNumber, brand, price), designType(designType), mountingType(mountingType) {}

    void displayInfo() const override {
        cout << "Наушники:" << endl;
        cout << "  Номенклатурный номер: " << partNumber << endl;
        cout << "  Бренд: " << brand << endl;
        cout << "  Цена: $" << price << endl;
        cout << "  Тип конструкции: " << designType << endl;
        cout << "  Способ крепления: " << mountingType << endl;
    }

private:
    string designType;  // Тип конструкции
    string mountingType; // Способ крепления
};

// Класс для микрофонов
class Microphone : public Device {
public:
    Microphone(const string& partNumber, const string& brand, double price,
               const string& frequencyRange, double sensitivity)
        : Device(partNumber, brand, price), frequencyRange(frequencyRange), sensitivity(sensitivity) {}

    void displayInfo() const override {
        cout << "Микрофон:" << endl;
        cout << "  Номенклатурный номер: " << partNumber << endl;
        cout << "  Бренд: " << brand << endl;
        cout << "  Цена: $" << price << endl;
        cout << "  Частотный диапазон: " << frequencyRange << endl;
        cout << "  Чувствительность: " << sensitivity << " дБ" << endl;
    }

private:
    string frequencyRange; // Частотный диапазон
    double sensitivity;    // Чувствительность
};

// Класс для клавиатур
class Keyboard : public Device {
public:
    Keyboard(const string& partNumber, const string& brand, double price,
             const string& switchType, const string& interfaceType)
        : Device(partNumber, brand, price), switchType(switchType), interfaceType(interfaceType) {}

    void displayInfo() const override {
        cout << "Клавиатура:" << endl;
        cout << "  Номенклатурный номер: " << partNumber << endl;
        cout << "  Бренд: " << brand << endl;
        cout << "  Цена: $" << price << endl;
        cout << "  Тип переключателей: " << switchType << endl;
        cout << "  Интерфейс: " << interfaceType << endl;
    }

private:
    string switchType;    // Тип переключателей
    string interfaceType; // Интерфейс
};


// ФАБРИЧНЫЙ ПАТТЕРН


// Абстрактная фабрика устройств
class DeviceFactory {
public:
    virtual Device* createDevice() const = 0;
    virtual string getDeviceType() const = 0;
    virtual ~DeviceFactory() = default;
};

// Фабрика для создания наушников
class HeadphonesFactory : public DeviceFactory {
private:
    string partNumber, brand, designType, mountingType;
    double price;

public:
    HeadphonesFactory(const string& pn, const string& b, double pr,
                      const string& dt, const string& mt)
        : partNumber(pn), brand(b), price(pr), designType(dt), mountingType(mt) {}

    Device* createDevice() const override {
        return new Headphones(partNumber, brand, price, designType, mountingType);
    }

    string getDeviceType() const override {
        return "Headphones";
    }
};

// Фабрика для создания микрофонов
class MicrophoneFactory : public DeviceFactory {
private:
    string partNumber, brand, frequencyRange;
    double price, sensitivity;

public:
    MicrophoneFactory(const string& pn, const string& b, double pr,
                      const string& fr, double sens)
        : partNumber(pn), brand(b), price(pr), frequencyRange(fr), sensitivity(sens) {}

    Device* createDevice() const override {
        return new Microphone(partNumber, brand, price, frequencyRange, sensitivity);
    }

    string getDeviceType() const override {
        return "Microphone";
    }
};

// Фабрика для создания клавиатур
class KeyboardFactory : public DeviceFactory {
private:
    string partNumber, brand, switchType, interfaceType;
    double price;

public:
    KeyboardFactory(const string& pn, const string& b, double pr,
                    const string& st, const string& it)
        : partNumber(pn), brand(b), price(pr), switchType(st), interfaceType(it) {}

    Device* createDevice() const override {
        return new Keyboard(partNumber, brand, price, switchType, interfaceType);
    }

    string getDeviceType() const override {
        return "Keyboard";
    }
};


// МЕНЕДЖЕР УСТРОЙСТВ


// Класс для управления устройствами через фабрики
class DeviceManager {
private:
    vector<DeviceFactory*> factories; // Храним фабрики
    vector<Device*> devices;          // Храним созданные устройства

    // Очистка потока ввода при ошибке
    void clearInput() {
        cin.clear();
        cin.ignore(numeric_limits<streamsize>::max(), '\n');
    }

    // Безопасный ввод целого числа
    int safeIntInput() {
        int value;
        while (true) {
            cin >> value;
            if (cin.fail()) {
                cout << "Ошибка ввода. Введите число: ";
                clearInput();
            } else {
                clearInput();
                return value;
            }
        }
    }

public:
    // Деструктор - освобождаем память
    ~DeviceManager() {
        for (auto factory : factories) delete factory;
        for (auto device : devices) delete device;
    }

    // Добавление фабрики в менеджер
    void addFactory(DeviceFactory* factory) {
        factories.push_back(factory);
    }

    // Создание устройств из всех фабрик
    void createAllDevices() {
        for (auto factory : factories) {
            devices.push_back(factory->createDevice());
        }
    }

    // Вывод нумерованного списка устройств
    void displayNumberedList() {
        cout << "\nСписок устройств:" << endl;
        for (size_t i = 0; i < devices.size(); ++i) {
            cout << i + 1 << ". ";
            devices[i]->displayShortInfo();
        }
        cout << devices.size() + 1 << ". Вернуться в меню" << endl;
    }

    // Вывод информации о выбранном устройстве
    void displayDeviceInfo(int index) {
        if (index >= 1 && index <= devices.size()) {
            cout << "\nИнформация об устройстве:" << endl;
            devices[index - 1]->displayInfo();

            int action;
            cout << "\nВыберите действие:" << endl;
            cout << "1. Вернуться в меню" << endl;
            cout << "2. Выход" << endl;
            cout << "Введите номер: ";
            action = safeIntInput();

            if (action == 2) {
                cout << "Завершение программы." << endl;
                exit(0);
            }
        }
    }

    // Запуск меню выбора устройств
    void runMenu() {
        while (true) {
            displayNumberedList();
            cout << "\nВведите номер устройства: ";
            
            int choice = safeIntInput();

            if (choice == devices.size() + 1) {
                break; // Возврат в предыдущее меню
            } else if (choice >= 1 && choice <= devices.size()) {
                displayDeviceInfo(choice);
            } else {
                cout << "Неверный номер. Выберите из списка." << endl;
            }
        }
    }
};


// МЕНЮ КАТЕГОРИЙ


// Меню для выбора категории устройств
void categoryMenu() {
    // Создаем менеджеры для каждой категории
    DeviceManager headphonesManager;
    DeviceManager microphonesManager;
    DeviceManager keyboardsManager;

    // Добавляем фабрики для наушников (3 шт)
    headphonesManager.addFactory(new HeadphonesFactory("H123", "Logitech G435", 99.99, "Накладные", "Оголовье"));
    headphonesManager.addFactory(new HeadphonesFactory("H456", "Razer Kraken Pro", 79.99, "Накладные", "Оголовье"));
    headphonesManager.addFactory(new HeadphonesFactory("H789", "HyperX Cloud II", 149.99, "Накладные", "Оголовье"));
    headphonesManager.createAllDevices();

    // Добавляем фабрики для микрофонов (3 шт)
    microphonesManager.addFactory(new MicrophoneFactory("M123", "Blue Yeti", 129.99, "20Гц-20кГц", 120));
    microphonesManager.addFactory(new MicrophoneFactory("M456", "Razer Seiren", 149.99, "30Гц-20кГц", 118));
    microphonesManager.addFactory(new MicrophoneFactory("M789", "Audio-Technica AT2020", 99.99, "20Гц-20кГц", 114));
    microphonesManager.createAllDevices();

    // Добавляем фабрики для клавиатур (3 шт)
    keyboardsManager.addFactory(new KeyboardFactory("K123", "Logitech G Pro", 129.99, "Механические", "USB"));
    keyboardsManager.addFactory(new KeyboardFactory("K456", "Razer BlackWidow", 159.99, "Оптические", "USB"));
    keyboardsManager.addFactory(new KeyboardFactory("K789", "Corsair K95", 199.99, "Механические", "USB"));
    keyboardsManager.createAllDevices();

    while (true) {
        cout << "\nВыберите категорию устройств:" << endl;
        cout << "1. Наушники" << endl;
        cout << "2. Микрофоны" << endl;
        cout << "3. Клавиатуры" << endl;
        cout << "4. Выход" << endl;
        cout << "Введите номер категории: ";

        int choice;
        cin >> choice;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        switch (choice) {
            case 1:
                headphonesManager.runMenu();
                break;
            case 2:
                microphonesManager.runMenu();
                break;
            case 3:
                keyboardsManager.runMenu();
                break;
            case 4:
                cout << "Завершение программы." << endl;
                return;
            default:
                cout << "Неверный номер. Выберите 1-4." << endl;
        }
    }
}

int main() {
    categoryMenu(); // Запускаем меню
    return 0;
}