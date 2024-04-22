CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       username TEXT NOT NULL UNIQUE,
                       password TEXT NOT NULL
);

CREATE TABLE timetable (
                           id INTEGER PRIMARY KEY AUTOINCREMENT,
                           user_id INTEGER,
                           day TEXT,
                           period TEXT,
                           course TEXT,
                           teacher TEXT,
                           FOREIGN KEY (user_id) REFERENCES users(id)
);
