package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ConnectDB() (*pgxpool.Pool, error) {
	connStr := "host=localhost port=5432 user=ykwais password=1111 dbname=afdb sslmode=disable search_path=cw_test"

	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	// Проверка подключения
	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: %v", err)
	}
	defer conn.Release()

	err = conn.Conn().Ping(context.Background())
	if err != nil {
		return nil, err
	}

	log.Println("Успешное подключение к базе данных.")
	return dbPool, nil
}

func RunMigrations(dbPool *pgxpool.Pool) error {
	migrationsPath := "db/migrations"
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("не удалось прочитать папку с миграциями: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(migrationsPath, file.Name())
			if err := executeSQLFile(dbPool, filePath); err != nil {
				return fmt.Errorf("ошибка выполнения миграции %s: %v", file.Name(), err)
			}
			log.Printf("Миграция %s успешно выполнена", file.Name())
		}
	}

	return nil
}

// executeSQLFile выполняет SQL-скрипт из файла
func executeSQLFile(dbPool *pgxpool.Pool, filePath string) error {
	sqlContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл %s: %v", filePath, err)
	}

	ctx := context.Background()
	_, err = dbPool.Exec(ctx, string(sqlContent))
	if err != nil {
		return fmt.Errorf("ошибка выполнения SQL: %v", err)
	}

	return nil
}
