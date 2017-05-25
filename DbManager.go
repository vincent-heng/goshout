package main

import (
    "database/sql"
    "time"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
    log.Printf("Connecting to DB")
    connectdb, err := sql.Open("sqlite3", "./goshout.db")
    checkErr(err)
    db = connectdb
    log.Printf("Connected to DB")
    CreateMessageTable()
    loadConfigurationFile()
}

func CreateMessageTable() {
    defer func() { // In case of existing table
        if recover() != nil {
            log.Printf("Table already exists, can't create");
        }
    }()
    log.Printf("DB - Creating Table Message")
    stmt, err := db.Prepare("CREATE TABLE message (" +
        "`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
        "`chatid` VARCHAR(64) NOT NULL," +
        "`nickname` VARCHAR(64) NOT NULL," +
        "`text` VARCHAR(64) NOT NULL," +
        "`timestamp` DATE NOT NULL)")
    checkErr(err)
    res, err := stmt.Exec()
    checkErr(err)
    log.Printf("DB - Created Table Message")
    _ = res // Ignore "res declared and not used"

    stmt2, err := db.Prepare("CREATE INDEX chatindex ON message (chatid)")
    checkErr(err)
    res2, err := stmt2.Exec()
    checkErr(err)
    _ = res2
}

func GetChats() []string {
    rows, err := db.Query("SELECT * FROM (SELECT chatid FROM message GROUP BY chatid ORDER BY chatid DESC LIMIT 100) ORDER BY chatid ASC") // TODO put the limit in config file
    checkErr(err)
    chats := []string{}
    for rows.Next() {
        var chatid string
        err = rows.Scan(&chatid)
        checkErr(err)
        chats = append(chats,chatid)
    }
    return chats
}

func PostMessage(chatId string, message Message) {
    stmt, err := db.Prepare("INSERT INTO message(chatid, nickname, text, timestamp) values(?,?,?,?)")
    checkErr(err)
    res, err := stmt.Exec(chatId, message.Nickname, message.Text, time.Now())
    checkErr(err)
    _ = res
}

func GetMessages(chatId string) []Message {
    rows, err := db.Query("SELECT * FROM (SELECT nickname, text, timestamp FROM message WHERE chatid=? ORDER BY timestamp DESC LIMIT 100) ORDER BY timestamp ASC", chatId) // TODO put the limit in config file
    checkErr(err)
    messages := []Message{}
    for rows.Next() {
        var message Message
        err = rows.Scan(&message.Nickname, &message.Text, &message.Timestamp)
        checkErr(err)
        messages = append(messages, message)
    }

    return messages
}
