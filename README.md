# Calculator Service

${\textsf{\color{red}🟥 Project Description}}$

This project is a simple web service calculator that supports basic operations such as addition, subtraction, multiplication, and division of single and multi-digit numbers. It is designed to process mathematical expressions, providing results in JSON format.

---

${\textsf{\color{orange}🟧 How Does It Work?}}$

```
HTTP request to the server
       |
       v
+-------------------+
|    Handler        |
|    (handler)      |
+-------------------+
       |
       v
+-------------------+
|   Service         |
|   (service)       |
+-------------------+
       |
       v
+-------------------+
|   Caching         |
|   (cache)         |
+-------------------+
       |
       v
+-------------------+
|   Models          |
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

${\textsf{\color{green}🟩 Example Requests}}$

| Request Type          | Example Request         | Result                                |
|-----------------------|-------------------------|---------------------------------------|
| ${\textsf{\color{lightgreen}Correct}}$              | `"expression": "2 + 2"` | Returns `{"result": "4"}`            |
| ${\textsf{\color{lightgreen}Correct}}$              | `"expression": "10 - 5"` | Returns `{"result": "5"}`            |
| ${\textsf{\color{orange}Incorrect}}$            | `"expression": "5 / 0"`  | Returns `{"error": "division by zero"}` |
| ${\textsf{\color{orange}Incorrect}}$            | `"expression": "invalid"` | Returns `{"error": "invalid expression"}` |
| ${\textsf{\color{darkred}Invalid}}$            | `"expression": ""`     | Returns `{"error": "Invalid JSON format"}` |

---

${\textsf{\color{blue}🟦 HTTP Statuses and Returned Values}}$

| Status | Description                                 |
|--------|---------------------------------------------|
| 200    | Success, request processed                  |
| 400    | Request error, invalid JSON                 |
| 404    | Not found, if the path is incorrect        |
| 405    | Method not allowed for this path           |
| 422    | Processing error, invalid expression        |
| 500    | Internal server error                       |

---

${\textsf{\color{purple}🟣 Conclusion: How to Install}}$

1. Clone the repository:
   
   ```bash
   git clone https://github.com/yourusername/calc_service.git
   cd calc_service
   ```

2. Make sure you have Go installed (version 1.18 or higher).

3. Install dependencies:
   
   ```bash
   go mod tidy
   ```

4. Create a configuration file `config.json` in the root of the project:

   ```json
   {
       "log_level": "debug",
       "port": 8080
   }
   ```

You can change the `port` value to configure the port on which the server will run.

5. Run the project:

   ```bash
   go run cmd/main.go
   ```

The server will be available at [http://localhost:8080](http://localhost:8080).
