CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR (255) NOT NULL,
    last_name VARCHAR (255) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL, 
    birthday DATE NOT NULL,
    
    username VARCHAR (50) UNIQUE NOT NULL, 
    passwd TEXT NOT NULL, 
    active BOOLEAN DEFAULT FALSE,
    last_login TIMESTAMP,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP
);