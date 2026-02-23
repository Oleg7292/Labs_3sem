#include "dbmanager.h"

DbManager::DbManager(const QString& path) : m_path(path) // подключение к PostgreSQL
{
    m_host = "localhost";
    m_port = 5432;
    m_dbName = "smeshariki";
    m_user = "user";
    m_password = "password";
}

DbManager::~DbManager()
{
    if (m_db.isOpen())
    {
        m_db.close();
    }
}

void DbManager::setConnectionDetails(const QString& host, int port, const QString& dbName, const QString& user, const QString& password)
{ // Переподключение к БД с новыми параметрами
    m_host = host;
    m_port = port;
    m_dbName = dbName;
    m_user = user;
    m_password = password;
}

bool DbManager::createConnection() // Создание подключения
{
    if (m_db.isOpen()) {
        m_db.close();
    }
    
    m_db = QSqlDatabase::addDatabase("QPSQL");
    m_db.setHostName(m_host);
    m_db.setPort(m_port);
    m_db.setDatabaseName(m_dbName);
    m_db.setUserName(m_user);
    m_db.setPassword(m_password);

    if (!m_db.open())
    {
        qDebug() << "Ошибка: не удалось подключиться к базе данных";
        qDebug() << m_db.lastError().text();
        return false;
    }
    else
    {
        qDebug() << "База данных: подключение успешно";
        return true;
    }
}

bool DbManager::isOpen() const // Проверка состояния подключения
{
    return m_db.isOpen();
}

QSqlQuery DbManager::executeQuery(const QString& queryStr)
{
    QSqlQuery query;
    if (!query.exec(queryStr)) {
        qDebug() << "Запрос не выполнен: " << query.lastError().text();
    }
    return query;
}

QString DbManager::getLastError() const {
    return m_db.lastError().text();
}
