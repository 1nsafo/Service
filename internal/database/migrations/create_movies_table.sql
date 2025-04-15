CREATE TABLE films(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    releaseYear TEXT,
	director TEXT,    
	genre TEXT,       
	rating DOUBLE PRECISION
);