CREATE TABLE [IF NOT EXISTS] User
(
    id SERIAL PRIMARY KEY NOT NULL,
    username VARCHAR (25) UNIQUE NOT NULL,
    first_name VARCHAR (50) ,
    last_name VARCHAR (50) ,
    mobile VARCHAR (11) ,
    national_code VARCHAR (10),
    address TEXT (500),
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,

    );

CREATE TABLE [IF NOT EXISTS] Role
(
    id SERIAL PRIMARY KEY NOT NULL,
    [name] VARCHAR (25) UNIQUE NOT NULL,
    );