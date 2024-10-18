package data

type contextKey string

// DBConn - ключ по которому хранится соединение с БД в контексте
const DBConn contextKey = "dbConn"
