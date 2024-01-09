package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const userTable = `CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT UNIQUE,
	username TEXT UNIQUE,
	password TEXT,
	session_token TEXT DEFAULT NULL,
	expiresAt DATETIME DEFAULT NULL
);`

const postTable = `CREATE TABLE IF NOT EXISTS post (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	author TEXT,
	title TEXT,
	description TEXT,
	likes INT DEFAULT 0,
	dislikes INT DEFAULT 0,
	category TEXT,
	created_at DATE DEFAULT (datetime('now','localtime')),
	FOREIGN KEY (author) REFERENCES user(username)
);`

const commentTable = `CREATE TABLE IF NOT EXISTS comment (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	id_post INT,
	author TEXT,
	comment TEXT,
	likes INT DEFAULT 0,
	dislikes INT DEFAULT 0,
	created_at DATE DEFAULT (datetime('now','localtime')),
	FOREIGN KEY (author) REFERENCES user(username)
);`

const reactionCommentTable = `CREATE TABLE IF NOT EXISTS likesComment(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	userId INTEGER,
	commentsId INTEGER,
	like1 INT,
    FOREIGN KEY (userId) REFERENCES user(id),
    FOREIGN KEY (commentsId) REFERENCES comment(id)
);`

const reactionPostTable = `CREATE TABLE IF NOT EXISTS likesPost(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	userId INTEGER,
	postId INTEGER,
	like1 INT,
    FOREIGN KEY (userId) REFERENCES user(id),
    FOREIGN KEY (postId) REFERENCES post(id)
);`

const categoriesTable = `CREATE TABLE IF NOT EXISTS category (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	tag TEXT,
	id_post INT,
    FOREIGN KEY (id_post) REFERENCES post(id)
);`

var tables = []string{userTable, postTable, commentTable, reactionPostTable, reactionCommentTable, categoriesTable, categoriesTable}

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "forum.db?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	for _, dbname := range tables {
		if _, err = db.Exec(dbname); err != nil {
			log.Fatal(err)
		}
	}
	return db
}
