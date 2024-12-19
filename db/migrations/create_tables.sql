
CREATE TABLE IF NOT EXISTS Users (
                                     id BIGSERIAL PRIMARY KEY,
                                     username VARCHAR(50) UNIQUE NOT NULL,
                                     password_hash VARCHAR(255) NOT NULL,
                                     email VARCHAR(100) NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Clients (
                                       id BIGSERIAL PRIMARY KEY,
                                       user_id BIGINT NOT NULL UNIQUE REFERENCES Users (id) ON DELETE CASCADE,
                                       name VARCHAR(100) NOT NULL,
                                       passport VARCHAR(11) NOT NULL,
                                       phone VARCHAR(15) NOT NULL,
                                       driving_experience INT CHECK (driving_experience >= 0)
);