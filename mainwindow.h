#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QTableView>
#include <QPushButton>
#include <QVBoxLayout>
#include <QHBoxLayout>
#include <QComboBox>
#include <QLabel>
#include <QTextEdit>
#include <QSqlQueryModel>
#include <QTabWidget>
#include <QMenuBar>
#include <QMenu>
#include <QAction>
#include <QDialog>
#include <QLineEdit>
#include <QFormLayout>
#include <QSpinBox>
#include <QDialogButtonBox>
#include "dbmanager.h"

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();

private slots:
    void executeSelectedQuery();
    void executeCustomQuery();
    void showConnectionDialog();

private:
    void setupUI();
    void setupQueries();
    void createConnectionMenu();

    DbManager *m_db;
    QTableView *m_tableView;
    QComboBox *m_querySelector;
    QLabel *m_statusLabel;
    QTextEdit *m_queryDescription;
    QTextEdit *m_customQueryEdit;
    
    struct QueryInfo {
        QString name;
        QString sql;
        QString description;
    };
    
    QList<QueryInfo> m_queries;
};

#endif 
