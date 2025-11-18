
A simple RESTful API built with Go.

## Features
- User authentication (signup, login)
- CRUD operations for posts


 Run Migrations
This will create the necessary tables in your database.
bash go run migrate/migrate.go

Authentication
- POST /signup`: Register a new user.
- POST /login`: Log in a user and receive a JWT token.

## Posts (Requires Authentication)
- `POST /posts`: Create a new post.
- `GET /posts`: Get all posts.
- `GET /posts/:id`: Get a single post by ID.
- `PUT /posts/:id`: Update a post by ID.
- `DELETE /posts/:id`: Delete a post by ID.
