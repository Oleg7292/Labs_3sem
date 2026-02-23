#include "mainwindow.h"
#include <QHeaderView>
#include <QMessageBox>
#include <QStatusBar>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
{
    m_db = new DbManager("smeshariki");
    if (!m_db->createConnection()) {
        QMessageBox::critical(this, "Ошибка базы данных", "Не удалось подключиться к базе данных. Убедитесь, что Docker запущен.");
    }

    createConnectionMenu();
    setupUI();
    setupQueries();
}

MainWindow::~MainWindow()
{
    delete m_db;
}

void MainWindow::createConnectionMenu() // Создание меню для подключения к БД
{
    QMenu *dbMenu = menuBar()->addMenu("База данных");
    QAction *connectAction = new QAction("Подключение...", this);
    connect(connectAction, &QAction::triggered, this, &MainWindow::showConnectionDialog);
    dbMenu->addAction(connectAction);
}

void MainWindow::setupUI() 
{
    QWidget *centralWidget = new QWidget(this);
    setCentralWidget(centralWidget);
    
    QVBoxLayout *mainLayout = new QVBoxLayout(centralWidget);
    
    // Header
    QLabel *header = new QLabel("Менеджер Базы Данных Смешарики", this);
    header->setAlignment(Qt::AlignCenter);
    header->setStyleSheet("font-size: 24px; font-weight: bold; color: #2c3e50; margin: 10px;");
    mainLayout->addWidget(header);

    QTabWidget *tabWidget = new QTabWidget(this);
    mainLayout->addWidget(tabWidget);

    QWidget *tab1 = new QWidget();
    QVBoxLayout *tab1Layout = new QVBoxLayout(tab1);

    QHBoxLayout *controlLayout = new QHBoxLayout();
    m_querySelector = new QComboBox(this);
    m_querySelector->setMinimumWidth(300);
    
    QPushButton *execButton = new QPushButton("Выполнить запрос", this);
    execButton->setCursor(Qt::PointingHandCursor);
    
    controlLayout->addWidget(new QLabel("Выберите запрос:"));
    controlLayout->addWidget(m_querySelector);
    controlLayout->addWidget(execButton);
    controlLayout->addStretch();
    
    tab1Layout->addLayout(controlLayout);

    m_queryDescription = new QTextEdit(this);
    m_queryDescription->setReadOnly(true);
    m_queryDescription->setMaximumHeight(60);
    m_queryDescription->setStyleSheet("background-color: #f8f9fa; border: 1px solid #dee2e6; border-radius: 4px; color: #495057;");
    tab1Layout->addWidget(m_queryDescription);
    
    tabWidget->addTab(tab1, "Готовые запросы");

    QWidget *tab2 = new QWidget();
    QVBoxLayout *tab2Layout = new QVBoxLayout(tab2);
    
    QLabel *customLabel = new QLabel("Введите SQL запрос:", this);
    tab2Layout->addWidget(customLabel);

    m_customQueryEdit = new QTextEdit(this);
    m_customQueryEdit->setPlaceholderText("SELECT * FROM characters WHERE ...");
    m_customQueryEdit->setStyleSheet("font-family: monospace;");
    tab2Layout->addWidget(m_customQueryEdit);

    QPushButton *execCustomButton = new QPushButton("Выполнить SQL", this);
    execCustomButton->setCursor(Qt::PointingHandCursor);
    tab2Layout->addWidget(execCustomButton);

    tabWidget->addTab(tab2, "Свой запрос");

    m_tableView = new QTableView(this);
    m_tableView->horizontalHeader()->setSectionResizeMode(QHeaderView::Stretch);
    m_tableView->setAlternatingRowColors(true);
    m_tableView->setStyleSheet("QTableView { border: 1px solid #dee2e6; gridline-color: #e9ecef; } QHeaderView::section { background-color: #e9ecef; padding: 4px; border: 1px solid #dee2e6; }");
    mainLayout->addWidget(m_tableView);

    m_statusLabel = new QLabel("Готов", this);
    statusBar()->addWidget(m_statusLabel);

    connect(execButton, &QPushButton::clicked, this, &MainWindow::executeSelectedQuery);
    connect(execCustomButton, &QPushButton::clicked, this, &MainWindow::executeCustomQuery);
    connect(m_querySelector, QOverload<int>::of(&QComboBox::currentIndexChanged), [this](int index){
        if(index >= 0 && index < m_queries.size()) {
            m_queryDescription->setText(m_queries[index].description + "\nSQL: " + m_queries[index].sql);
        }
    });

    resize(800, 600);
    setWindowTitle("Лабораторная работа №6 - БД Смешарики");
    
    setStyleSheet(
        "QMainWindow { background-color: white; }"
        "QPushButton { background-color: #007bff; color: white; border: none; padding: 8px 16px; border-radius: 4px; font-weight: bold; }"
        "QPushButton:hover { background-color: #0056b3; }"
        "QPushButton:pressed { background-color: #004085; }"
        "QComboBox { padding: 6px; border: 1px solid #ced4da; border-radius: 4px; }"
    );
}

void MainWindow::setupQueries()
{
    m_queries = {
        {
            "1. Список персонажей и характеров",
            "SELECT c.name, c.species, p.main_trait FROM characters c JOIN personalities p ON c.personality_id = p.id;", 
            "Показывает всех персонажей и их главную черту характера, используя INNER JOIN."
        },
        {
            "2. Количество предметов у персонажа",
            "SELECT c.name, COUNT(i.id) as item_count FROM characters c LEFT JOIN items i ON c.id = i.owner_id GROUP BY c.name;", 
            "Считает, сколько предметов принадлежит каждому персонажу, используя LEFT JOIN и GROUP BY."
        },
        {
            "3. Эпизоды 1-го сезона",
            "SELECT title, release_date FROM episodes WHERE season = 1;", 
            "Выводит список всех эпизодов первого сезона, используя условие WHERE."
        },
        {
            "4. Персонажи в эпизоде 'Скамейка'",
            "SELECT c.name FROM characters c JOIN character_episodes ce ON c.id = ce.character_id JOIN episodes e ON ce.episode_id = e.id WHERE e.title = 'Скамейка';",
            "Находит всех персонажей, которые появлялись в эпизоде 'Скамейка', используя несколько JOIN."
        },
        {
            "5. Локации с предметами",
            "SELECT l.name, COUNT(i.id) as count FROM locations l JOIN items i ON l.id = i.location_id GROUP BY l.name HAVING COUNT(i.id) > 0;",
            "Показывает локации, в которых есть хотя бы один предмет, используя HAVING."
        },
        {
            "6. Всего предметов (Агрегация)",
            "SELECT COUNT(*) as total_items FROM items;",
            "Вычисляет общее количество предметов в базе данных, используя агрегатную функцию."
        },
        {
            "7. Друзья Кроша",
            "SELECT c.name as Friend, r.relationship_type FROM characters c JOIN relationships r ON (c.id = r.char2_id AND r.char1_id = 1) OR (c.id = r.char1_id AND r.char2_id = 1);",
            "Находит всех персонажей, связанных с Крошем (ID 1), используя сложную логику в JOIN."
        },
        {
            "8. Предметы в доме владельца",
            "SELECT c.name as Owner, i.name as Item, l.name as Location FROM characters c JOIN items i ON c.id = i.owner_id JOIN locations l ON i.location_id = l.id WHERE l.name LIKE '%' || c.name || '%';",
            "Находит предметы, которые находятся в месте, названном в честь их владельца (например, Скамейка Кроша в Доме Кроша)."
        },
        {
            "9. Обновление: Переместить Скамейку в Море",
            "UPDATE items SET location_id = (SELECT id FROM locations WHERE name = 'Море') WHERE name = 'Скамейка';",
            "Обновляет местоположение предмета 'Скамейка' на 'Море'. (Запустите запрос 10 или проверьте таблицу, чтобы увидеть эффект)."
        },
        {
            "10. Показать все предметы (Проверка обновления)",
            "SELECT i.name, l.name as location FROM items i LEFT JOIN locations l ON i.location_id = l.id;",
            "Выводит список всех предметов и их местоположение для проверки изменений."
        }
    };

    for (const auto& q : m_queries) {
        m_querySelector->addItem(q.name);
    }
    
    if (!m_queries.isEmpty()) {
        m_queryDescription->setText(m_queries[0].description + "\nSQL: " + m_queries[0].sql);
    }
}

void MainWindow::executeCustomQuery()
{
    QString sql = m_customQueryEdit->toPlainText().trimmed();
    if (sql.isEmpty()) return;

    if (sql.startsWith("UPDATE", Qt::CaseInsensitive) || 
        sql.startsWith("DELETE", Qt::CaseInsensitive) || 
        sql.startsWith("INSERT", Qt::CaseInsensitive) ||
        sql.startsWith("CREATE", Qt::CaseInsensitive) ||
        sql.startsWith("DROP", Qt::CaseInsensitive) ||
        sql.startsWith("ALTER", Qt::CaseInsensitive)) {
        
        QSqlQuery query = m_db->executeQuery(sql);
        if (query.lastError().isValid()) {
            m_statusLabel->setText("Ошибка: " + query.lastError().text());
            QMessageBox::warning(this, "Ошибка запроса", query.lastError().text());
        } else {
            m_statusLabel->setText("Запрос выполнен успешно. Затронуто строк: " + QString::number(query.numRowsAffected()));
            m_tableView->setModel(nullptr); 
        }
    } else {
        QSqlQueryModel *model = new QSqlQueryModel(this);
        model->setQuery(sql);
        
        if (model->lastError().isValid()) {
            m_statusLabel->setText("Ошибка: " + model->lastError().text());
            delete model;
        } else {
            m_tableView->setModel(model);
            m_statusLabel->setText("Запрос вернул " + QString::number(model->rowCount()) + " строк.");
        }
    }
}

void MainWindow::showConnectionDialog()
{
    QDialog dialog(this);
    dialog.setWindowTitle("Подключение к БД");
    
    QFormLayout form(&dialog);
    
    QLineEdit *hostEdit = new QLineEdit("localhost");
    QSpinBox *portEdit = new QSpinBox();
    portEdit->setRange(1, 65535);
    portEdit->setValue(5432);
    QLineEdit *dbEdit = new QLineEdit("smeshariki");
    QLineEdit *userEdit = new QLineEdit("user");
    QLineEdit *passEdit = new QLineEdit("password");
    passEdit->setEchoMode(QLineEdit::Password);
    
    form.addRow("Хост:", hostEdit);
    form.addRow("Порт:", portEdit);
    form.addRow("Имя БД:", dbEdit);
    form.addRow("Пользователь:", userEdit);
    form.addRow("Пароль:", passEdit);
    
    QDialogButtonBox buttonBox(QDialogButtonBox::Ok | QDialogButtonBox::Cancel, Qt::Horizontal, &dialog);
    form.addRow(&buttonBox);
    
    connect(&buttonBox, &QDialogButtonBox::accepted, &dialog, &QDialog::accept);
    connect(&buttonBox, &QDialogButtonBox::rejected, &dialog, &QDialog::reject);
    
    if (dialog.exec() == QDialog::Accepted) {
        m_db->setConnectionDetails(
            hostEdit->text(),
            portEdit->value(),
            dbEdit->text(),
            userEdit->text(),
            passEdit->text()
        );
        
        if (m_db->createConnection()) {
            m_statusLabel->setText("Подключено к " + dbEdit->text());
            QMessageBox::information(this, "Успех", "Подключение установлено успешно!");
        } else {
            m_statusLabel->setText("Ошибка подключения");
            QMessageBox::critical(this, "Ошибка", "Не удалось подключиться: " + m_db->getLastError());
        }
    }
}

void MainWindow::executeSelectedQuery()
{
    int index = m_querySelector->currentIndex();
    if (index < 0 || index >= m_queries.size()) return;

    QString sql = m_queries[index].sql;
    
    if (sql.startsWith("UPDATE") || sql.startsWith("DELETE") || sql.startsWith("INSERT")) {
        QSqlQuery query = m_db->executeQuery(sql);
        if (query.lastError().isValid()) {
            m_statusLabel->setText("Ошибка: " + query.lastError().text());
            QMessageBox::warning(this, "Ошибка запроса", query.lastError().text());
        } else {
            m_statusLabel->setText("Запрос выполнен успешно. Затронуто строк: " + QString::number(query.numRowsAffected()));
            m_tableView->setModel(nullptr); 
        }
    } else {
        QSqlQueryModel *model = new QSqlQueryModel(this);
        model->setQuery(sql);
        
        if (model->lastError().isValid()) {
            m_statusLabel->setText("Ошибка: " + model->lastError().text());
            delete model;
        } else {
            m_tableView->setModel(model);
            m_statusLabel->setText("Запрос вернул " + QString::number(model->rowCount()) + " строк.");
        }
    }
}
