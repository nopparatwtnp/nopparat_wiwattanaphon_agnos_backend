## Setup and Running

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-repo/password-strength-api.git
   cd password-strength-api

1. **Create Initial Database:**
    ```bash
    docker exec -it postgres psql -U user -d password_db -c "CREATE TABLE logs (id SERIAL PRIMARY KEY, request TEXT, response TEXT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"


1. **Run:**
   docker-compose up --build