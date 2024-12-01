package config_tests

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"hackathons-app/internal/config"

	"github.com/google/uuid"
)

var configFile string

func getProjectRoot() (string, error) {
	// Получаем текущую рабочую директорию
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("ошибка при определении текущей директории: %w", err)
	}

	// Проходим вверх по дереву каталогов, чтобы найти go.mod
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil // Нашли корень проекта
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("корень проекта (с go.mod) не найден")
		}
		dir = parent
	}
}
func fillStruct(cfgType reflect.Type, cfgValue reflect.Value) {
	// Проходим по всем полям структуры
	for i := 0; i < cfgValue.NumField(); i++ {
		field := cfgValue.Field(i)

		// Проверяем, что поле экспортируемое
		if field.CanSet() {
			// Заполняем поле значением. Для примера можно просто передать строку.
			// Это можно заменить на более сложную логику, например, получение значений из окружения.
			switch field.Kind() {
			case reflect.String:
				field.SetString(uuid.NewString())
			case reflect.Int:
				field.SetInt(rand.Int63())
			case reflect.Bool:
				field.SetBool(true)
			}
		}
	}

	// Для вложенных структур тоже рекурсивно заполняем
	for i := 0; i < cfgValue.NumField(); i++ {
		field := cfgValue.Field(i)

		// Если это вложенная структура, вызываем рекурсивно fillStruct
		if field.Kind() == reflect.Struct {
			fillStruct(cfgType, field)
		}
	}
}

func serializeConfig(cfg interface{}) (string, error) {
	var sb strings.Builder
	cfgValue := reflect.ValueOf(cfg)

	if cfgValue.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected a struct, got %s", cfgValue.Kind())
	}

	for i := 0; i < cfgValue.NumField(); i++ {
		field := cfgValue.Field(i)
		fieldType := cfgValue.Type().Field(i)

		envTag := fieldType.Tag.Get("env")
		if envTag == "" {
			envTag = fieldType.Name
		}

		if field.Kind() == reflect.Struct {
			subConfig, err := serializeConfig(field.Interface())
			if err != nil {
				return "", err
			}
			sb.WriteString(subConfig)
			continue
		}

		sb.WriteString(fmt.Sprintf("%s=%v\n", envTag, field.Interface()))
	}

	return sb.String(), nil
}

func TestMain(m *testing.M) {
	path, err := getProjectRoot()
	if err != nil {
		panic(err)
	}

	configFile = filepath.Join(path, "test.env")
	_, err = os.Create(configFile)
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()
	os.Remove(configFile)
	os.Exit(exitVal)
}

func TestWriteConfig(t *testing.T) {
	cfg := config.Config{}

	// Заполняем структуру значениями
	cfgType := reflect.TypeOf(cfg)
	fillStruct(cfgType, reflect.ValueOf(&cfg).Elem())

	// Генерируем строку для записи в файл, рекурсивно обрабатывая вложенные структуры
	content, err := serializeConfig(cfg)
	if err != nil {
		t.Fatalf("Ошибка при сериализации конфигурации: %v", err)
	}

	// Записываем в файл
	err = os.WriteFile(configFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Ошибка при записи в файл: %v", err)
	}

	// Читаем конфиг из файла
	readCfg := config.GetConfig(configFile)

	// Сравниваем оригинальный и прочитанный конфиги
	if !reflect.DeepEqual(cfg, readCfg) {
		t.Errorf("Ожидается: %+v, получено: %+v", cfg, readCfg)
	}
}
