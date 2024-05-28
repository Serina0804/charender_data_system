CREATE TABLE schedule (
    id INTEGER PRIMARY KEY,
    month     INTEGER,
	date      INTEGER,
	day       TEXT,
	eventName TEXT,
	startHour INTEGER,
	startMin  INTEGER,
	endHour   INTEGER,
	endMin    INTEGER,
	memo      TEXT,
	record    TEXT,
);