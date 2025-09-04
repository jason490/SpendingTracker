DROP TABLE IF EXISTS "User";
DROP TABLE IF EXISTS "Tags";
DROP TABLE IF EXISTS "Expenses";

CREATE TABLE IF NOT EXISTS "Users" (
    "id" INTEGER NOT NULL UNIQUE PRIMARY KEY,
    "session_id" uuid UNIQUE,
    "username" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "email" TEXT UNIQUE NOT NULL,
    "total_spending" REAL,
    "created_at" INTEGER -- INSERT account created_at 
);

CREATE TABLE IF NOT EXISTS "Tags" (
    "id" INTEGER NOT NULL PRIMARY KEY UNIQUE,
    "user_id" TEXT,
    "name" TEXT,
    FOREIGN KEY ("user_id") REFERENCES "Users"("id")
);

CREATE TABLE IF NOT EXISTS "Expenses" (
    "id" INTEGER NOT NULL PRIMARY KEY UNIQUE,
    "user_id" TEXT,
    "tag_id" INTEGER,
    "name" TEXT,
    "description" TEXT,
    "cost" REAL,
    FOREIGN KEY ("user_id") REFERENCES "Users"("id"),
    FOREIGN KEY ("tag_id") REFERENCES "Tags"("id")
);

