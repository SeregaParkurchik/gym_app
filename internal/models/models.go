package models

import "time"

// Определение ролей
const (
	RoleMember        = "member"
	RoleTrainer       = "trainer"
	RoleAdministrator = "administrator"
)

// Структура для представления члена клуба
type ClubMember struct {
	ID             int       // Уникальный идентификатор члена клуба
	Name           string    // Имя члена клуба
	MembershipType string    // Тип абонемента
	StartDate      time.Time // Дата начала абонемента
	EndDate        time.Time // Дата окончания абонемента
	Visits         []Visit   // Список посещений
	Role           string    // Роль пользователя
}

// Структура для представления посещения
type Visit struct {
	EntryTime time.Time // Время входа
	ExitTime  time.Time // Время выхода
}

// Структура для представления тренера
type Trainer struct {
	ID       int               // Уникальный идентификатор тренера
	Name     string            // Имя тренера
	Schedule []TrainingSession // Расписание тренировок
	Role     string            // Роль пользователя
}

// Структура для представления администратора
type Administrator struct {
	ID   int    // Уникальный идентификатор администратора
	Name string // Имя администратора
	Role string // Роль пользователя
}

// Структура для представления тренировки
type TrainingSession struct {
	ID          int       // Уникальный идентификатор тренировки
	TrainerID   int       // ID тренера
	Time        time.Time // Время тренировки
	Format      string    // Формат тренировки (групповая/индивидуальная)
	Description string    // Описание тренировки
}
