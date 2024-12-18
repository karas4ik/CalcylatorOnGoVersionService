# Calculator Service

${\textsf{\color{red}🟥 Project Description}}$

This project is a simple web service calculator that supports basic operations such as addition, subtraction, multiplication, and division of single and multi-digit numbers. It is designed to process mathematical expressions, providing results in JSON format.

---

${\textsf{\color{orange}🟧 How Does It Work?}}$

```
  HTTP запрос к серверу
           |
           v
  +-------------------+
  |   Обработчик      |
  |   (handler)       |
  +-------------------+
           |
           v
  +-------------------+
  |   Сервис          |
  |   (service)       |
  +-------------------+
           |
           v
  +-------------------+
  |   Кэширование     |
  |   (cache)         |
  +-------------------+
           |
           v
  +-------------------+
  |   Модели          |
  |   (model)         |
  +-------------------+
```

### Component Descriptions:
- **Handler (handler)**: Processes incoming HTTP requests and returns responses.
- **Service (service)**: Contains the logic for solving tasks.
- **Caching (cache)**: Manages caching of results for improved performance.
- **Models (model)**: Defines data structures for requests and responses.

---

${\textsf{\color{yellow}🟨 Interesting Points:}}$

#### 🏃 r.Run()

The `r.Run()` method is responsible for starting the HTTP server and processing incoming requests. It uses the Gin framework, which makes it easy to set up routing and handle requests. When `Run()` is called, the server starts listening on the specified port (in our case, 8080) and awaits incoming requests. The handlers defined earlier in the code will be invoked in response to these requests.

#### 📜 Logging

Logging is implemented using the Logrus library. We initialize the logger in the `InitLogger` function, where we set the logging level and format. All significant events, such as server startup, request handling, and errors, are logged. This allows us to monitor the application's behavior and identify issues in real-time.

#### 🔒 Error Handling

The project features its own error handling system. We create various error types, such as `InvalidExpressionError` and `DivisionByZeroError`, to provide more precise feedback to users about issues with their requests. This makes the code more readable and helps identify the causes of errors more easily.

#### ➗ Arithmetic Expression Handler

The arithmetic expression handler takes the input expression in JSON format and uses the `CalculateArithmeticExpression` method from the service for processing. It first checks if the JSON is correctly formatted. If the data is valid, the expression is passed for processing. The service tokenizes the expression, converts it to postfix notation, and performs the calculations. As a result, the client receives either a result or an error message.

---

${\textsf{\color{green}🟩 Примеры запросов}}$

| Тип запроса              | Пример запроса         | Результат                              |
|-------------------------|-----------------------|----------------------------------------|
| ${\textsf{\color{lightgreen}Корректный}}$              | `"expression": "2 + 2"` | Вернёт `{"result": "4"}`              |
| ${\textsf{\color{lightgreen}Корректный}}$              | `"expression": "10 - 5"` | Вернёт `{"result": "5"}`              |
| ${\textsf{\color{orange}Некорректный}}$            | `"expression": "5 / 0"`  | Вернёт `{"error": "division by zero"}` |
| ${\textsf{\color{orange}Некорректный}}$            | `"expression": "invalid"` | Вернёт `{"error": "invalid expression"}` |
| ${\textsf{\color{darkred}Неправильный}}$            | `"expression": ""`     | Вернёт `{"error": "Invalid JSON format"}` |

---

${\textsf{\color{blue}🟦 Статусы HTTP и возвращаемые значения}}$

| Статус | Описание                                     |
|--------|----------------------------------------------|
| 200    | Успех, запрос обработан                     |
| 400    | Ошибка запроса, некорректный JSON          |
| 404    | Не найдено, если путь неправилен            |
| 405    | Метод не разрешен для данного пути         |
| 422    | Ошибка обработки, неверное выражение        |
| 500    | Внутренняя ошибка сервера                   |


---

${\textsf{\color{purple}🟣 Заключение: Как установить}}$

1. Клонируйте репозиторий:
   
   ```bash
   git clone https://github.com/yourusername/calc_s>ervice.git
   cd calc_service
   ```

2. Убедитесь, что у вас установлен Go
(версия 1.18 и выше).

3. Установите зависимости:
   
   ```bash
   go mod tidy
   ```

4. Создайте файл конфигурации `config.json` в корне проекта:

   ```json
   {
       "log_level": "debug"
   }
   ```

5. Запустите проект:

   ```bash
   go run cmd/main.go
   ```

Сервер будет доступен по адресу [http://localhost:8080](http://localhost:8080).
