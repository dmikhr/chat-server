package data

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

// ChatModel - модель чата
type ChatModel struct {
	Ctx context.Context
}

// MessageModel - модель сообщения
type MessageModel struct {
	Ctx context.Context
}

// Models - wrapper
type Models struct {
	Chat    ChatModel
	Message MessageModel
}

// Chat - структура данных чата
type Chat struct {
	Name      string
	Users     []int64
	CreatedAt time.Time
}

// Message - структура данных сообщения
type Message struct {
	From      string
	Msg       string
	CreatedAt time.Time
}

// Create - создать чат с пользователями
func (m ChatModel) Create(chat *Chat) (int64, error) {
	queryChat := `INSERT INTO chat (name, created_at) VALUES ($1, $2) RETURNING id`
	queryUser := `INSERT INTO chat_users (chat_id, user_id) VALUES ($1, $2)`
	// достаем БД соеднинение из контекста
	conn := m.Ctx.Value(DBConn).(*pgx.Conn)
	tx, err := conn.Begin(m.Ctx)
	if err != nil {
		return -1, err
	}
	defer func() {
		err = tx.Rollback(m.Ctx)
		if err != nil {
			fmt.Println("rollback err:", err)
		}
	}()

	var chatID int64
	err = tx.QueryRow(m.Ctx, queryChat, chat.Name, chat.CreatedAt).Scan(&chatID)
	if err != nil {
		return -1, err
	}

	for _, userID := range chat.Users {
		_, err = tx.Exec(m.Ctx, queryUser, chatID, userID)
		if err != nil {
			return -1, err
		}
	}

	if err = tx.Commit(m.Ctx); err != nil {
		return -1, err
	}
	return chatID, nil
}

// Delete - удаление чата
func (m ChatModel) Delete(id int64) error {
	query := `DELETE FROM chat WHERE id=$1`
	conn := m.Ctx.Value(DBConn).(*pgx.Conn)
	cmdTag, err := conn.Exec(m.Ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("nothing was deleted")
	}
	return nil
}

// Send message - отправить сообщение
func (m MessageModel) Send(message *Message) error {
	query := `INSERT INTO message (name, msg, created_at) VALUES ($1, $2, $3)`
	// достаем БД соеднинение из контекста
	conn := m.Ctx.Value(DBConn).(*pgx.Conn)
	cmdTag, err := conn.Exec(m.Ctx, query, message.From, message.Msg, message.CreatedAt)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("nothing was updated")
	}
	return nil
}

// InitModels - инициализация моделей данных
func InitModels(ctx context.Context) Models {
	return Models{
		Chat:    ChatModel{Ctx: ctx},
		Message: MessageModel{Ctx: ctx},
	}
}
