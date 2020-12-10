


CREATE TABLE users (
     id SERIAL PRIMARY KEY,
     username VARCHAR ( 500 ) UNIQUE NOT NULL,  
     password VARCHAR ( 500 ) NOT NULL
);

CREATE TABLE notes(
    id SERIAL PRIMARY KEY,
    ownerid INTEGER NOT NULL,
    note TEXT NOT NULL
);