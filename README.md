
# API Students
This is a simple API to manage student data.

## Project Structure

- **`main.go`**: Arquivo principal que inicializa o servidor e configura o roteamento da API.
  
- **`models/`**: Contém os modelos de dados usados pela aplicação, como a definição de estudantes e suas relações com o banco de dados.
  
- **`routes/`**: Aqui estão definidos os endpoints da API. Cada rota é mapeada para um manipulador de requisições que define a lógica de negócios.
  
- **`config/`**: Arquivos de configuração que lidam com o banco de dados e outras variáveis de ambiente, como variáveis de configuração para o ambiente de produção ou desenvolvimento.
  
- **`swagger/`**: Contém a documentação gerada automaticamente usando o Swagger, que fornece uma interface visual para testar as rotas da API.


## Technologies used:
- Go
- GORM
- SQLite
- Swagger

## Routes:
### `GET /students`
- **Description**:  Retrieve a list of all students.
- **Response**:     A list of students, including their details.

### `GET /students/:id`
- **Description**: Retrieve information of a specific student by their id.
- **Parameters**:  `id` (integer): The ID of the student.
- **Response**:    The student's details.
  
### `POST /students`
- **Description**: Create a new student.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "cpf": 12345678901,
    "email": "john.doe@example.com",
    "age": 25,
    "active": true
  }

### `UPDATE /students/:id`
- **Description**: Update a specific student's information.
- **Parameters**:  `id` (integer): The ID of the student.
- **Request Body**: The updated student's details.

### `DELETE /students/:id`
- **Description**: Delete a student by their id.
- **Parameters**:  `id` (integer): The ID of the student.
- **Request Body**: The deleted student's details.

### STEPS TO RUN
## Clonar o Repository

````
https://github.com/MayaraPerez/api-students.git
````

### Installing dependencies:
After cloning, enter the project directory and install the necessary dependencies.
Install Go dependencies: If you don't have Go installed
Download the project dependencies: Inside the project directory, run:

- [Go](https://go.dev/doc/install) 1.20
- [Swagger](https://github.com/swaggo/gin-swagger)

### How to run the project:
In the main directory, run the command:

```
make run
```

### API Routes:
API routes are documented using Swagger.<br>
To understand in detail about the routes, run the application and see the `/swagger/index.html` route.

### Next steps:
- Add tests
- Run the project with Docker
