# 🧠 Question Manager API - One Table CRUD

This project is a simple backend to manage quiz questions. It uses a single collection `questions` in MongoDB and provides standard CRUD operations.

**Server URL:** `http://localhost:6060`

---

## 🛠️ Question APIs

### 1. Create a Question
- **Method:** `POST`
- **Full URL:** `http://localhost:6060/api/questions`
- **Request Body:**
```json
{
  "question": "What is the capital of France?",
  "options": ["Paris", "London", "Berlin", "Madrid"]
}
```

### 2. Get All Questions
- **Method:** `GET`
- **Full URL:** `http://localhost:6060/api/questions`
- **Description:** List every question in the database.

### 3. Get Single Question
- **Method:** `GET`
- **Full URL:** `http://localhost:6060/api/questions/:id` (e.g., `/api/questions/65eb2b4e8c5b6b1234567890`)
- **Description:** Get details of one specific question.
- **Success Response:**
```json
{
  "id": "65eb2b4e8c5b6b1234567890",
  "question": "What is the capital of France?",
  "options": ["Paris", "London", "Berlin", "Madrid"],
  "created_at": "..."
}
```

### 4. Update a Question
- **Method:** `PUT`
- **Full URL:** `http://localhost:6060/api/questions/:id` (e.g., `/api/questions/65eb2b4e8c5b6b1234567890`)
- **Request Body:**
```json
{
  "question": "Updated question text?",
  "options": ["Option 1", "Option 2", "Option 3", "Option 4"]
}
```

### 5. Delete a Question
- **Method:** `DELETE`
- **Full URL:** `http://localhost:6060/api/questions/:id` (e.g., `/api/questions/65eb2b4e8c5b6b1234567890`)
- **Response:**
```json
{
  "message": "Question deleted successfully"
}
```

---

## ⚙️ Features
- **Single Collection**: Only uses the `questions` collection.
- **MongoDB Database**: Stores data in a NoSQL MongoDB instance.
- **JSON Format**: Easy to integrate with any frontend.
- **MVC Structure**: Clean, organized code.
