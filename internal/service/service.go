package service

import (
	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/repo"
	"github.com/google/uuid"
)

type Auth interface {
	ParseToken(accessToken string) (uuid.UUID, []string, error)
}

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

// type User interface {
// 	CreateUser(user entity.User) (int, error)
// }

// type Todo interface {
// 	CreateTodo(todo entity.Todo) (uuid.UUID, error)
// 	GetTodos(userId uuid.UUID) ([]entity.Todo, error)
// 	DeleteTodo(id, userId uuid.UUID) error
// 	UpdateTodo(id uuid.UUID, newTodo entity.Todo) error
// }

type Services struct {
	Auth Auth
	// User       User
	// Todo       Todo
	Auditorium Auditorium
	Discipline Discipline
	Hometask   Hometask
	Lesson     Lesson
}

type ServicesDependencies struct {
	Repo    *repo.Repositories
	SignKey string
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Auth:       NewAuthService(deps.SignKey),
		Auditorium: NewAuditoriumService(deps.Repo.Auditorium),
		Discipline: NewDisciplineService(deps.Repo.Discipline),
		Hometask:   NewHometaskService(deps.Repo.Hometask),
		Lesson:     NewLessonService(deps.Repo.Lesson),
		// User: NewUserService(deps.Repo.User),
		// Todo: NewTodoService(deps.Repo.Todo),
	}
}
