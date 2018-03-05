-- Read https://stackoverflow.com/questions/21757722/how-to-use-sqlite-decimal-precision-notation
-- for an interesting explanation of SQLites datatypes, in particular affinities.

CREATE TABLE IF NOT EXISTS users (
    id          INTEGER PRIMARY KEY AUTOINCREMENT, 
    username    TEXT
);

CREATE TABLE IF NOT EXISTS transactions (
    id          INTEGER PRIMARY KEY AUTOINCREMENT, 
    userid      INTEGER,
    -- For easier filtering
    year        INTEGER,
    month       INTEGER,
    timestamp   DATETIME DEFAULT CURRENT_TIMESTAMP,

    category    TEXT, 
    amount      DECIMAL(8,2),
    FOREIGN KEY(userid) REFERENCES users(id));

-- Add default user.
INSERT INTO users VALUES (1, 'user');