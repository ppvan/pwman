BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
	"username"	TEXT NOT NULL DEFAULT 'guardian' UNIQUE,
	"password"	TEXT NOT NULL,
	"id"	INTEGER,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "passwords" (
	"id"	INTEGER,
	"website"	TEXT NOT NULL,
	"password"	TEXT NOT NULL,
	"username"	TEXT NOT NULL,
	"created_at"	DATETIME,
	"updated_at"	DATETIME,
	"user_id"	INTEGER,
	FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE INDEX IF NOT EXISTS "user_idx" ON "passwords" (
	"user_id"
);
COMMIT;
