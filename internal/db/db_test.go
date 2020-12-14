package db_test

const userstablecreationquery = ` CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username VARCHAR ( 500 ) UNIQUE NOT NULL,  
	password VARCHAR ( 500 ) NOT NULL
);
`

const notestablecreationquery = ` CREATE TABLE IF NOT EXISTS notes (
	id SERIAL PRIMARY KEY,
	username VARCHAR ( 500 ) UNIQUE NOT NULL,  
	password VARCHAR ( 500 ) NOT NULL
);
`
