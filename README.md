# Brazilian Championship Table API

This project consists of two APIs related to the Brazilian football championship: a table data API (`api-tabela`) and a team query API (`api-futebol`).

## API Tabela

The Tabela API provides a list of teams and their positions in the Brazilian championship table. It allows queries to retrieve all teams or filter by team name.

### Endpoints

- `GET /tabela`: Returns all teams in the table.
  Example usage:
  ```bash
  curl http://localhost:8000/tabela
  ```

- `GET /tabela?time={team_name}`: Returns the specific team that matches the provided name.
  Example usage:
  ```bash
  curl http://localhost:8000/tabela?time=Flamengo
  ```

### Installation and Execution

1. **Clone the repository:**
   ```bash
   git clone https://github.com/WesleyBSa/api-futebol.git
   cd repository-name
   ```

2. **Install dependencies:**
   Make sure you have Go installed. Then, install dependencies using:
   ```bash
   go mod download
   ```

3. **Run the Tabela API:**
   ```bash
   go run cmd/main.go
   ```
   The API will be available at `http://localhost:8000`.

## API Futebol

The Futebol API is responsible for integrating with the Tabela API to provide detailed data about specific teams, including statistics, players, and other relevant details.


### Installation and Execution

1. **Clone the repository:**
   ```bash
   git clone https://github.com/WesleyBSa/api-futebol.git
   cd repository-name
   ```

2. **Install dependencies:**
   Make sure you have Go installed. Then, install dependencies using:
   ```bash
   go mod download
   ```

3. **Run the Futebol API:**
   ```bash
   go run cmd/main.go
   ```
   The API will be available at `http://localhost:8001`.

### Contribution

Feel free to contribute improvements to this project. Follow these steps to contribute:

1. Fork the project
2. Create a branch with your feature (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Adding a new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a new Pull Request

### License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
