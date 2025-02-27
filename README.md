# Golang JWT Authentication API

🚀 A secure, scalable, and production-ready authentication system built with **Golang**, **Mux**, and **MongoDB**. This project implements JWT-based authentication with user registration, login, and protected routes.

## ✨ Features
- 🔐 **JWT Authentication** - Secure user authentication with access & refresh tokens
- 🛡 **Middleware Protection** - Restrict access to authorized users
- 💾 **MongoDB Integration** - Store user data with indexed email authentication
- 🚀 **RESTful API** - Fully functional API with structured endpoints
- ⚡ **Fast & Lightweight** - Optimized with the Mux router

## 📂 Project Structure
```
📁 golang-jwt-authentication
│-- 📂 controllers       # API logic for authentication
│-- 📂 database          # MongoDB connection handler
│-- 📂 helpers           # Utility functions (hashing, JWT handling)
│-- 📂 middleware        # Authentication middleware
│-- 📂 models            # Data models for MongoDB
│-- 📂 routers           # Route definitions for endpoints
│-- 📜 main.go           # Entry point for the server
│-- 📜 go.mod            # Go module dependencies
│-- 📜 README.md         # Project documentation
```

## 📌 API Endpoints
| Method | Endpoint         | Description              | Auth Required |
|--------|-----------------|--------------------------|--------------|
| POST   | `/users/signup` | Register a new user      | ❌ No        |
| POST   | `/users/login`  | Authenticate user & get token | ❌ No   |
| GET    | `/users/{id}`   | Get user details         | ✅ Yes       |

## 🛠 Installation
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/yourusername/golang-jwt-authentication.git
cd golang-jwt-authentication
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Set Up Environment Variables
Create a `.env` file and add:
```env
MONGODB_URI=mongodb+srv://your_username:your_password@cluster.mongodb.net/
SECRET_KEY=your_secret_key
PORT=8000
```

### 4️⃣ Run the Application
```sh
go run main.go
```

## 🚀 Deployment
You can deploy this project on:
- **Heroku** *(Recommended for beginners)*
- **AWS Lambda + API Gateway** *(For serverless architecture)*
- **Docker + Kubernetes** *(For containerized deployment)*

## 🛡 Security Best Practices
✔ Use **hashed passwords** (bcrypt)  
✔ Store **tokens securely**  
✔ Validate **JWT tokens** with expiration  
✔ Restrict **API access with middleware**

## 🤝 Contributing
Pull requests are welcome! Feel free to fork and improve the project.

## 📜 License
This project is licensed under the MIT License. Feel free to use and modify! 🚀
