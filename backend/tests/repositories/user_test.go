package repositories_tests

import (
	"testing"

	"hackathons-app/internal/models"
	"hackathons-app/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Создаем новую функцию для получения mock репозитория
func newMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	// Создаем mock для PostgreSQL
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// Преобразуем в *gorm.DB
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

// Тест для метода GetAllUsers
func TestGetAllUsers(t *testing.T) {
	// Создаем mock DB
	db, mock, err := newMockDB()
	assert.NoError(t, err)

	// Создаем репозиторий
	repo := repositories.NewUserRepository(db)

	// Мокаем запрос к базе, учитывая "deleted_at IS NULL"
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."deleted_at" IS NULL`).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "John Doe").
			AddRow(2, "Jane Doe"),
	)

	// Вызов метода репозитория
	users, err := repo.GetAllUsers()

	// Проверяем, что ошибок нет
	assert.NoError(t, err)

	// Проверяем, что возвращенные данные правильные
	assert.Len(t, users, 2)
	assert.Equal(t, "John Doe", users[0].FirstName)
	assert.Equal(t, "Jane Doe", users[1].FirstName)

	// Проверяем что запрос был выполнен
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// Тест для метода GetAllUsersWithHackathons
func TestGetAllUsersWithHackathons(t *testing.T) {
	// Создаем mock DB
	db, mock, err := newMockDB()
	assert.NoError(t, err)

	// Создаем репозиторий
	repo := repositories.NewUserRepository(db)

	// Мокаем запрос к базе
	mock.ExpectQuery("SELECT \\* FROM \"users\"").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "John Doe").
			AddRow(2, "Jane Doe"),
	)
	mock.ExpectQuery("SELECT \\* FROM \"hackathons\" WHERE \"user_id\" = ?").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "user_id"}).
			AddRow(1, "Hackathon1", 1).
			AddRow(2, "Hackathon2", 1),
	)

	// Вызов метода репозитория
	users, err := repo.GetAllUsersWithHackathons()

	// Проверяем, что ошибок нет
	assert.NoError(t, err)

	// Проверяем, что количество пользователей правильное
	assert.Len(t, users, 2)
	assert.Len(t, users[0].Hackathons, 2) // Проверяем, что у первого пользователя два хакатона

	// Проверяем, что запрос был выполнен
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// Тест для метода GetUserInfoById
func TestGetUserInfoById(t *testing.T) {
	// Создаем mock DB
	db, mock, err := newMockDB()
	assert.NoError(t, err)

	// Создаем репозиторий
	repo := repositories.NewUserRepository(db)

	// Мокаем запрос к базе
	mock.ExpectQuery("SELECT \\* FROM \"users\" WHERE \"id\" = ?").WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "John Doe"),
	)

	// Вызов метода репозитория
	user, err := repo.GetUserInfoById(1)

	// Проверяем, что ошибок нет
	assert.NoError(t, err)

	// Проверяем, что данные пользователя правильные
	assert.Equal(t, "John Doe", user.FirstName)

	// Проверяем, что запрос был выполнен
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// Тест для метода CreateUser
func TestCreateUser(t *testing.T) {
	// Создаем mock DB
	db, mock, err := newMockDB()
	assert.NoError(t, err)

	// Создаем репозиторий
	repo := repositories.NewUserRepository(db)

	// Настраиваем моки для выполнения INSERT
	mock.ExpectQuery(`INSERT INTO "users" \("created_at","updated_at","deleted_at","first_name","second_name","email","telegram_id","hashed_password"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7,\$8\) RETURNING "id"`).
		WithArgs(
			sqlmock.AnyArg(), // created_at
			sqlmock.AnyArg(), // updated_at
			nil,              // deleted_at
			"John Doe",       // first_name
			"",               // second_name
			"",               // email
			"",               // telegram_id
			"",               // hashed_password
		).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1), // Возвращаемый ID
		)

	// Создаем нового пользователя
	user := models.User{
		FirstName: "John Doe",
		// Другие поля пустые
	}

	// Вызов метода репозитория
	err = repo.CreateUser(&user)

	// Проверяем, что ошибок нет
	assert.NoError(t, err)

	// Проверяем, что ID пользователя установлен
	assert.Equal(t, 1, user.ID)

	// Проверяем, что все ожидания выполнены
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// Тест для метода DeleteUserById
func TestDeleteUserById(t *testing.T) {
	// Создаем mock DB
	db, mock, err := newMockDB()
	assert.NoError(t, err)

	// Создаем репозиторий
	repo := repositories.NewUserRepository(db)

	// Мокаем запрос для проверки существования пользователя
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(1, 1). // ID пользователя и LIMIT
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "John Doe"),
		)

	// Мокаем начало транзакции
	mock.ExpectBegin()

	// Мокаем запрос на удаление (обновление deleted_at)
	mock.ExpectExec(`UPDATE "users" SET "deleted_at"=`).
		WithArgs(sqlmock.AnyArg(), 1).            // timestamp и ID пользователя
		WillReturnResult(sqlmock.NewResult(1, 1)) // Успешное удаление

	// Мокаем завершение транзакции
	mock.ExpectCommit()

	// Вызов метода репозитория
	err = repo.DeleteUserById(1)

	// Проверяем, что ошибок нет
	assert.NoError(t, err)

	// Проверяем, что все ожидания выполнены
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
