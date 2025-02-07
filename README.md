# Student Management Rest APIS

## Overview

Student Management RESI Api is a backend service built with Go (GoLang) to Manage student records. it provides endpoints to perform basic CURD (Create, Update , Read , Delete) Operations on student Data and is designed to  be lightweight, scalable, and easy to use.

### Features
  -  Add Student    : Create New Student Record.             
  -  View Student   : Featch Details Of all Students or Specific Student By ID. Add Student.
  -  Update Student : Update The Existing Student Details.
  -  Delete Student : Remove Student Records From The System.

### Teck Stack
-   Programming Language : Go Language(GoLang).
-   Database : PostgreSQL.
-   Frameworks/Libraries : .
    -  Gin For Http Request Handling.
    -  GORM for ORM- based Database .
    -  Viper For Configuration Management.
 
###  Api Endpoints :
####  Base URL : http://loaclhost:8080/api/v1

#### Endponits: 

  | Method   | Endpoint           | Description                 | Request Body                     |  Response (Success)                           |
  |----------|--------------------|-----------------------------|----------------------------------|-----------------------------------------------|
  | GET      |`\students`         | Fetch All Students          | None                             | `[{id , name , age , grade , created_at}]`    |
  | GET      |`\students/{id}`    | Fetch Student By ID         | None                             | `[{id , name , age, grade , created_at }]`    |
  | POST     |`\students`         | Add a new Student           | `{"name":"john Doe", "age":20, "Grade":"A"}`| `{"message":"student added"}`      |
  | PUT      |`\students\{id}`    | Update an Existing Student  | `{"name":"john smith", "age":21,"Grade":"A+"}`| `{"message":"Student Updated"}`  |
  | DELETE   |`\students\{id}`    | Delete a Student By ID      | None                             | `{"message":"student deleted"}`
  
### Setup and Installation

#### Prerequisites

  - Go (Version 1.19 or later)
  - Docker
  - Postgres DB
  - Git
  - Postman

 #### Steps

   1. Clone The Repository

      ```bash
      git clone https://github.com/KhomaneNandkumar/student-management-rest.git
      cd student-management-rest
      ```

   2. Install Dependencies

      ```bash
      go mod init tidy
      ```

   3. Set enviroment Variables: create a.env file at the project root:

      ```bash
       DB_NAME=<>
       DB_HOST=localhost
       DB_PORT=5432
       DB_USER=<>
       DB_PASSWORD=<>
       SERVER_ADDRESS=0.0.0.0:8085
      ```

  4. Run Database Migration

      ```bash
      go run  migration/main.go
      ```

  5. Build Server

      ```bash
      go build .
      ```

  6. Start Server

     ```bash
     ./student-management-rest
     ```
       
