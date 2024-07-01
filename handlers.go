package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

func createUser(c *gin.Context, db *sql.DB) {
	var user User
	// ShouldBindJSON() аналог функции Unmarshal() из пакета encoding/json.
	// берет json тело ответа и записывает его в переданную переменную.
	if err := c.ShouldBindJSON(&user); err != nil {
		// JSON() принимает статус код, объект, который нужно передать в ответе клиенту, преобразовывает его в json и отправляет.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("insert into users (username, password) values (:username, :password)",
		sql.Named("username", user.Username), sql.Named("password", user.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}

func login(c *gin.Context, db *sql.DB) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := db.QueryRow("select id from users where username = :username and password = :password",
		sql.Named("username", user.Username), sql.Named("password", user.Password))

	var id int
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func createChannel(c *gin.Context, db *sql.DB) {
	var channel Channel
	if err := c.ShouldBindJSON(&channel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := db.Exec("insert into channels (name) values (:name)", sql.Named("name", channel.Name))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func listChannels(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("select id, name from channels")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var channels []Channel

	for rows.Next() {
		var channel Channel

		err := rows.Scan(&channel.ID, &channel.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		channels = append(channels, channel)
	}

	c.JSON(http.StatusOK, channels)
}

func createMessage(c *gin.Context, db *sql.DB) {
	var message Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("insert into messages (channel_id, user_id, message) values (:channel_id, :user_id, :message)",
		sql.Named("channel_id", message.ChannelID), sql.Named("user_id", message.UserID), sql.Named("message", message.Text))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func listMessages(c *gin.Context, db *sql.DB) {
	channelID, err := strconv.Atoi(c.Query("channelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 100
	}

	lastMessageID, err := strconv.Atoi(c.Query("lastMessageID"))
	if err != nil {
		lastMessageID = 0
	}

	rows, err := db.Query("SELECT m.id, channel_id, user_id, u.username AS user_name, message FROM messages m LEFT JOIN users u ON u.id = m.user_id WHERE channel_id = :channelID AND m.id > lastMessageID ORDER BY m.id ASC LIMIT :limit",
		sql.Named("channelID", channelID), sql.Named("lastMessageID", lastMessageID), sql.Named("limit", limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var messages []Message

	for rows.Next() {
		var message Message

		err := rows.Scan(&message.ID, &message.ChannelID, &message.UserID, &message.UserName, &message.Text)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		messages = append(messages, message)
	}

	c.JSON(http.StatusOK, messages)
}
