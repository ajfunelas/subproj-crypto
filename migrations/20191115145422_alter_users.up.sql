DROP TABLE users;

CREATE TABLE users (
    ID text UNIQUE NOT NULL,
    USERNAME varchar(255),
    PASSWORD varchar(255),
    EMAIL varchar(255)
);

