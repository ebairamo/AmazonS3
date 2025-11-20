# AmazonS3
# Triple-S

Упрощенная реализация объектного хранилища Amazon S3 на языке Go. Проект реализует базовые операции S3 через REST API для управления бакетами и объектами.

## 🚀 Возможности

- **Управление бакетами**: создание, просмотр списка и удаление бакетов
- **Операции с объектами**: загрузка, получение и удаление файлов
- **REST API**: полностью совместимый с S3 HTTP-интерфейс
- **XML ответы**: все ответы в формате XML согласно спецификации Amazon S3
- **Метаданные**: хранение информации о бакетах и объектах в CSV файлах
- **Валидация имен**: проверка имен бакетов согласно требованиям S3

## 🛠️ Технологии

- **Язык**: Go (только стандартная библиотека)
- **Протокол**: HTTP/REST
- **Формат данных**: XML для ответов, CSV для метаданных
- **Хранилище**: файловая система

## 📦 Установка

```bash
# Клонировать репозиторий
git clone https://github.com/ebairamo/triple-s.git
cd triple-s

# Собрать проект
go build -o triple-s .
```

## 🎯 Использование

### Запуск сервера

```bash
# Базовый запуск
./triple-s --port 8080 --dir ./data

# Справка
./triple-s --help
```

**Параметры:**
- `--port N` - номер порта (по умолчанию 8080)
- `--dir S` - путь к директории для хранения данных

### API Endpoints

#### Управление бакетами

**Создать бакет**
```bash
curl -X PUT http://localhost:8080/my-bucket
```

**Список всех бакетов**
```bash
curl http://localhost:8080/
```

**Удалить бакет**
```bash
curl -X DELETE http://localhost:8080/my-bucket
```

#### Операции с объектами

**Загрузить файл**
```bash
curl -X PUT http://localhost:8080/my-bucket/photo.jpg \
  -H "Content-Type: image/jpeg" \
  --data-binary @photo.jpg
```

**Получить файл**
```bash
curl http://localhost:8080/my-bucket/photo.jpg -o downloaded.jpg
```

**Удалить файл**
```bash
curl -X DELETE http://localhost:8080/my-bucket/photo.jpg
```

## 🏗️ Архитектура

### Структура данных

```
data/
├── buckets.csv              # Метаданные бакетов
└── my-bucket/              # Директория бакета
    ├── objects.csv         # Метаданные объектов
    ├── photo.jpg           # Файл объекта
    └── document.pdf        # Другой объект
```

### Формат метаданных

**buckets.csv**
- Name - уникальное имя бакета
- CreationTime - время создания
- LastModifiedTime - время последней модификации
- Status - статус бакета

**objects.csv**
- ObjectKey - ключ объекта
- Size - размер в байтах
- ContentType - MIME тип
- LastModified - время последнего изменения

## ✅ Правила именования бакетов

- Длина: 3-63 символа
- Разрешены: строчные буквы, цифры, дефисы (-), точки (.)
- Не должно быть форматом IP-адреса
- Не начинается и не заканчивается дефисом
- Нет двух последовательных точек или дефисов

## 📝 Примеры использования

```bash
# Создать бакет
curl -X PUT http://localhost:8080/photos

# Загрузить изображение
curl -X PUT http://localhost:8080/photos/sunset.png \
  -H "Content-Type: image/png" \
  --data-binary @sunset.png

# Получить изображение
curl http://localhost:8080/photos/sunset.png -o sunset.png

# Список бакетов
curl http://localhost:8080/

# Удалить объект
curl -X DELETE http://localhost:8080/photos/sunset.png

# Удалить бакет
curl -X DELETE http://localhost:8080/photos
```

## 🔧 Обработка ошибок

Сервер возвращает соответствующие HTTP коды состояния:

- **200 OK** - успешная операция
- **204 No Content** - успешное удаление
- **400 Bad Request** - невалидное имя бакета
- **404 Not Found** - бакет или объект не найден
- **409 Conflict** - бакет уже существует или не пуст

## 🎓 Цели обучения

Этот проект демонстрирует:
- Построение REST API на Go
- Работу с HTTP и net/http пакетом
- Управление файловой системой
- Валидацию данных
- Обработку XML
- Базовые принципы облачного хранилища

## 📚 Ссылки

- [Документация Amazon S3 API](https://docs.aws.amazon.com/s3/)
- [Go net/http пакет](https://pkg.go.dev/net/http)
- [REST API принципы](https://restfulapi.net/)

## 🙏 Автор задания

**Savva Savostyanov**
- Email: savvax@savvax.com
- [GitHub](https://github.com/savvax)
- [LinkedIn](https://www.linkedin.com/in/savvax/)

# Triple-S

A simplified Amazon S3 object storage implementation in Go. This project implements core S3 operations through a REST API for managing buckets and objects.

## 🚀 Features

- **Bucket Management**: create, list, and delete buckets
- **Object Operations**: upload, retrieve, and delete files
- **REST API**: fully S3-compatible HTTP interface
- **XML Responses**: all responses in XML format following Amazon S3 specification
- **Metadata Storage**: bucket and object information stored in CSV files
- **Name Validation**: bucket name validation according to S3 requirements

## 🛠️ Tech Stack

- **Language**: Go (standard library only)
- **Protocol**: HTTP/REST
- **Data Format**: XML for responses, CSV for metadata
- **Storage**: File system

## 📦 Installation

```bash
# Clone repository
git clone https://github.com/ebairamo/triple-s.git
cd triple-s

# Build project
go build -o triple-s .
```

## 🎯 Usage

### Starting the Server

```bash
# Basic usage
./triple-s --port 8080 --dir ./data

# Help
./triple-s --help
```

**Options:**
- `--port N` - port number (default 8080)
- `--dir S` - path to data directory

### API Endpoints

#### Bucket Management

**Create bucket**
```bash
curl -X PUT http://localhost:8080/my-bucket
```

**List all buckets**
```bash
curl http://localhost:8080/
```

**Delete bucket**
```bash
curl -X DELETE http://localhost:8080/my-bucket
```

#### Object Operations

**Upload file**
```bash
curl -X PUT http://localhost:8080/my-bucket/photo.jpg \
  -H "Content-Type: image/jpeg" \
  --data-binary @photo.jpg
```

**Get file**
```bash
curl http://localhost:8080/my-bucket/photo.jpg -o downloaded.jpg
```

**Delete file**
```bash
curl -X DELETE http://localhost:8080/my-bucket/photo.jpg
```

## 🏗️ Architecture

### Data Structure

```
data/
├── buckets.csv              # Bucket metadata
└── my-bucket/              # Bucket directory
    ├── objects.csv         # Object metadata
    ├── photo.jpg           # Object file
    └── document.pdf        # Another object
```

### Metadata Format

**buckets.csv**
- Name - unique bucket name
- CreationTime - creation timestamp
- LastModifiedTime - last modification timestamp
- Status - bucket status

**objects.csv**
- ObjectKey - object key
- Size - size in bytes
- ContentType - MIME type
- LastModified - last modification timestamp

## ✅ Bucket Naming Rules

- Length: 3-63 characters
- Allowed: lowercase letters, numbers, hyphens (-), dots (.)
- Must not be formatted as an IP address
- Must not start or end with a hyphen
- No consecutive dots or hyphens

## 📝 Usage Examples

```bash
# Create bucket
curl -X PUT http://localhost:8080/photos

# Upload image
curl -X PUT http://localhost:8080/photos/sunset.png \
  -H "Content-Type: image/png" \
  --data-binary @sunset.png

# Get image
curl http://localhost:8080/photos/sunset.png -o sunset.png

# List buckets
curl http://localhost:8080/

# Delete object
curl -X DELETE http://localhost:8080/photos/sunset.png

# Delete bucket
curl -X DELETE http://localhost:8080/photos
```

## 🔧 Error Handling

The server returns appropriate HTTP status codes:

- **200 OK** - successful operation
- **204 No Content** - successful deletion
- **400 Bad Request** - invalid bucket name
- **404 Not Found** - bucket or object not found
- **409 Conflict** - bucket already exists or not empty

## 🎓 Learning Objectives

This project demonstrates:
- Building REST APIs in Go
- Working with HTTP and net/http package
- File system management
- Data validation
- XML processing
- Basic cloud storage principles

## 📚 References

- [Amazon S3 API Documentation](https://docs.aws.amazon.com/s3/)
- [Go net/http package](https://pkg.go.dev/net/http)
- [REST API principles](https://restfulapi.net/)

## 🙏 Project Author

**Savva Savostyanov**
- Email: savvax@savvax.com
- [GitHub](https://github.com/savvax)
- [LinkedIn](https://www.linkedin.com/in/savvax/)
