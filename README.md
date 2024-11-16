
# API Students
This is a simple API to manage student data.

## Estrutura do Projeto
main.go:    Arquivo principal que inicializa o servidor.
models:     Contém as definições dos modelos de dados.
routes:     Define as rotas da API.
config:     Configurações do banco de dados e outras variáveis de ambiente.
swagger:    Documentação gerada automaticamente usando o Swagger.

### Tecnologias utilizadas
- Go
- GORM
- SQLite
- Swagger

## Routes
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

## Instalando dependências:
Depois de clonar, entre no diretório do projeto e instale as dependências necessárias.
Instale as dependências do Go: Se você não tiver o Go instalado
Baixe as dependências do projeto: Dentro do diretório do projeto, rode:

- [Go](https://go.dev/doc/install) 1.20
- [Swagger](https://github.com/swaggo/gin-swagger)

## Como rodar o projeto
No diretório principal, rode o comando:

```
make run
```

## Rotas da API
As rotas da API estão documentadas usando o Swagger. Para entender com detalhes sobre as rotas, rode a aplicação e consulte a rota `/swagger/index.html`.


## Próximos passos
- Adicionar testes
- Rodar o projeto com Docker