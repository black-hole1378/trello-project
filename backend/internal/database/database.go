package database

import (
	"backend/internal/config"
	"backend/internal/models"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

var singleInstance *gorm.DB

func connectDatabase() *gorm.DB {
	cfg := config.GetInstance()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DB.Exec(stateQuery("role_states", []string{cfg.Roles.Admin, cfg.Roles.User})).Error; err != nil {
		log.Fatalf("failed to create enum type: %v", err)
	}

	status := []string{cfg.Status.Completed, cfg.Status.Planned, cfg.Status.Progress}

	if err = DB.Exec(stateQuery("task_status", status)).Error; err != nil {
		log.Fatalf("failed to create enum type: %v", err)
	}

	if err = DB.Exec(stateQuery("completed_state", []string{"YES", "NO"})).Error; err != nil {
		log.Fatalf("failed to create enum type: %v", err)
	}
	err = DB.AutoMigrate(&models.Task{}, &models.User{}, &models.SubTask{}, &models.UserWorkSpace{}, &models.WorkSpace{}, &models.Comment{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return DB
}

func GetInstance() *gorm.DB {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = connectDatabase()
		}
	}
	return singleInstance
}
