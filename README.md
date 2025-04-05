🔧 Технологический стек:

Язык: Go

Фреймворк: Gin

ORM: GORM

База данных: PostgreSQL

Аутентификация: JWT

Миграции: GORM

Контейнеризация: Docker + Docker Compose

🏗️ Архитектура проекта:

1. Роли пользователей:

Пациент

Врач

Администратор

3. Основные сущности:

User (Пациент/Врач/Админ)

Hospital (Тип: частная / государственная)

Address (для проверки прописки)

Doctor (Привязан к больнице и специализации)

Specialization (терапевт, дерматолог и т.д.)

Appointment (Запись к врачу)

Prescription (Назначенные лекарства и визиты)

✨ Фичи:

🏥 Больницы:

Частные: доступны всем

Государственные: только по адресу прописки

📆 Запись:

Список всех больниц

Список доступных врачей по специализации

Свободные временные слоты

Проверка наличия регистрации по адресу для гос. больниц

👨‍⚕️ Врач:

Регистрирует пациента к себе

Назначает лечение и следующую дату визита

Информация сохраняется и видна пациенту

👤 Личный кабинет:

Уведомления

История записей

Лекарства и назначения

🧱 REST API (CRUD):

/auth/register, /auth/login

/hospitals (GET, POST, PUT, DELETE)

/doctors, /appointments, /patients, /prescriptions

/specializations

🔐 Аутентификация и авторизация:

JWT (доступ по ролям)

Middleware:

AuthMiddleware

RoleBasedAccessMiddleware

📦 Docker:

Dockerfile для Go-приложения

docker-compose.yml:

api (приложение)

postgres

pgadmin

🧠 Паттерны проектирования:

Repository pattern – доступ к данным

Service layer – бизнес-логика

DTO / Mapper – разделение слоев

Factory pattern – создание сущностей

Dependency Injection
