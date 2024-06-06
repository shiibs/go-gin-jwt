Go Gin-Gonic JWT User Authentication with PostgreSQL and GORM
Introduction

This project is a Go-based web application that uses the Gin-Gonic framework to implement user authentication with JWT (JSON Web Tokens). The backend is powered by a PostgreSQL database, and GORM is used for ORM (Object Relational Mapping). The application provides endpoints for user registration, login, authentication, and token refresh.

Features

    User Registration: Allows new users to register.
    User Login: Authenticates registered users and issues JWT tokens.
    Token Authentication: Middleware to protect routes and verify JWT tokens.
    Token Refresh: Allows refreshing of JWT tokens to maintain session validity.

Technologies Used

    Go: Programming language.
    Gin: Web framework.
    JWT: JSON Web Tokens for secure token-based authentication.
    PostgreSQL: Relational database.
    GORM: ORM library for Go.
