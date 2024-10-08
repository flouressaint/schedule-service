package repo

import (
	"fmt"
	"log"

	"github.com/flouressaint/schedule-service/config"
	"github.com/flouressaint/schedule-service/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type User interface {
// 	CreateUser(user entity.User) (int, error)
// }

// type Todo interface {
// 	CreateTodo(todo entity.Todo) (uuid.UUID, error)
// 	GetTodos(userId uuid.UUID) ([]entity.Todo, error)
// 	DeleteTodo(id, userId uuid.UUID) error
// 	UpdateTodo(id uuid.UUID, newTodo entity.Todo) error
// }

type Auditorium interface {
	CreateAuditorium(auditorium entity.Auditorium) (uint, error)
	GetAuditoriums() ([]entity.Auditorium, error)
	DeleteAuditorium(id uint) error
	UpdateAuditorium(id uint, auditorium entity.Auditorium) error
}

type Discipline interface {
	CreateDiscipline(discipline entity.Discipline) (uint, error)
	GetDisciplines() ([]entity.Discipline, error)
	DeleteDiscipline(id uint) error
	UpdateDiscipline(id uint, discipline entity.Discipline) error
}

type Hometask interface {
	CreateHometask(hometask entity.Hometask) (uint, error)
	GetHometasks() ([]entity.Hometask, error)
	DeleteHometask(id uint) error
	UpdateHometask(id uint, hometask entity.Hometask) error
}

type Lesson interface {
	CreateLesson(lesson entity.Lesson) (uint, error)
	GetLessons() ([]entity.Lesson, error)
	DeleteLesson(id uint) error
	UpdateLesson(id uint, lesson entity.Lesson) error
}

type Repositories struct {
	Auditorium
	Discipline
	Hometask
	Lesson
	// User
	// Todo
}

func NewRepositories(config config.Config) *Repositories {
	// connect to the database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	log.Println("Connected Successfully to the Database")

	return &Repositories{
		Auditorium: NewAuditoriumRepo(db),
		Discipline: NewDisciplineRepo(db),
		Hometask:   NewHometaskRepo(db),
		Lesson:     NewLessonRepo(db),
		// User: NewUserRepo(db),
		// Todo: NewTodoRepo(db),
	}
}
