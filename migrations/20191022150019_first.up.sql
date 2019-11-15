CREATE TABLE users (
    ID INT,
    USERNAME varchar(255),
    PASSWORD varchar(255),
    EMAIL varchar(255)
);

CREATE TABLE tickers (
    id TEXT PRIMARY KEY NOT NULL,
    price TEXT NOT NULL,
    time TEXT NOT NULL,
    bid TEXT NOT NULL,
    ask TEXT NOT NULL,
    volume TEXT NOT NULL,
    size TEXT
);

CREATE TABLE user_favourites (
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    coin_id TEXT NOT NULL
);

