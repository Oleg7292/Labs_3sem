#ifndef DBMANAGER_H
#define DBMANAGER_H

#include <QString>
#include <QSqlDatabase>
#include <QSqlQuery>
#include <QSqlError>
#include <QDebug>
#include <QVariant>

class DbManager
{
public:
    DbManager(const QString& path);
    ~DbManager();

    bool isOpen() const;
    void setConnectionDetails(const QString& host, int port, const QString& dbName, const QString& user, const QString& password);
    bool createConnection();
    QSqlQuery executeQuery(const QString& queryStr);
    QString getLastError() const;

private:
    QSqlDatabase m_db;
    QString m_path;
    
    QString m_host;
    int m_port;
    QString m_dbName;
    QString m_user;
    QString m_password;
};

#endif 
