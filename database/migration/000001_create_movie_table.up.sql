CREATE TABLE IF NOT EXISTS movie (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    year INTEGER NOT NULL,
    rated VARCHAR(10) NOT NULL,
    released DATE,
    runtime VARCHAR(32) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    director VARCHAR(255) NOT NULL,
    writer VARCHAR(255) NOT NULL,
    actors VARCHAR(255) NOT NULL,
    language VARCHAR(255),
    country VARCHAR(255) ,
    awards VARCHAR(255) ,
    metascore VARCHAR(32) NOT NULL,
    imdb_rating VARCHAR(32) NOT NULL,
    imdb_votes VARCHAR(32) NOT NULL,
    imdb_id VARCHAR(20),
    type VARCHAR(20),
    dvd DATE,
    box_office VARCHAR(20),
    production VARCHAR(255),
    website VARCHAR(255),
    response BOOLEAN NOT NULL
    );