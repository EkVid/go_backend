### go.mod

This file is used for dependency management and stores the dependencies needed to build and run the program. It can be initialized with the command go mod init <module_name>. It identifies dependencies from the import statements and creates an entry for them in this file.

### go.sum

This file is also part of the dependency management in Golang. It is essentially a checksum for the dependencies to ensure that dependency is installed only once and with the correct version.

### main.go

This is the main entry point of our code, responsible for setting up the backend server, configuring logging, applying middleware, and defining API routes.

### 🚀 How to Run

1. **Clone the repository**

```bash
git clone -b main https://github.com/EkVid/go_backend.git
cd go_backend
```

2. **Install Dependencies**

```bash
go mod tidy
```

3. **Set Up PostgreSQL DB Connection (example)**

```bash
psql -U postgres
CREATE USER test_user WITH PASSWORD 'yourpassword';
CREATE DATABASE go_backend_db;
GRANT ALL PRIVILEGES ON DATABASE go_backend_db TO test_user;
```

4. **Create a .env file**

```bash
DB_USER=your_user_name
DB_PASSWORD=your_password
DB_NAME=your_db_name
```

5. **Run the app**

```bash
go run main.go
```

**The server should now be running at http://localhost:8080**
