package db

import (
    "fmt"
    "log"
    "sync"
    "time"

    "forum-app/backend/internal/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var (
    instance *gorm.DB
    once     sync.Once
)

func Connect(dsn string) (*gorm.DB, error) {
    var err error
    once.Do(func() {
        cfg := &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info),
        }
        db, dbErr := gorm.Open(mysql.Open(dsn), cfg)
        if dbErr != nil {
            err = dbErr
            return
        }
        sqlDB, sqlErr := db.DB()
        if sqlErr != nil {
            err = sqlErr
            return
        }
        sqlDB.SetConnMaxLifetime(5 * time.Minute)
        sqlDB.SetMaxOpenConns(10)
        sqlDB.SetMaxIdleConns(5)
        if migrateErr := db.AutoMigrate(&models.Board{}, &models.Topic{}, &models.User{}, &models.Post{}, &models.Comment{}); migrateErr != nil {
            err = migrateErr
            return
        }
        instance = db
    })
    if err != nil {
        return nil, fmt.Errorf("db init failed: %w", err)
    }
    return instance, nil
}

func Get() *gorm.DB {
    if instance == nil {
        log.Fatal("database not initialized")
    }
    return instance
}

