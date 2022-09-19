package messages

import (
	"github.com/deltamc/otus-social-networks-chat/db"
	"github.com/deltamc/otus-social-networks-chat/models/users"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type Message struct {
	Id         string    `db:"id" json:"id"`
	ShardId    uint8     `db:"shard_id" json:"shard_id"`
	UserIdFrom int64     `db:"user_id_from" json:"user_id_from"`
	UserIdTo   int64     `db:"user_id_to" json:"user_id_to"`
	Message    string    `db:"message" json:"message"`
	CreatedAt  time.Time `json:"created_at"  db:"created_at"`
}

var shards = db.ShardNodes{}

func getShards() db.ShardNodes {
	if len(shards.Nodes) == 0 {
		shards.Add(1, "node-1")
		shards.Add(2, "node-2")
	}

	return shards
}

func (m *Message) New(user users.User) (lastID string, err error) {
	dbPool := db.OpenDB(db.MessagesShard1)

	m.Id = uuid.New().String()
	m.UserIdFrom = user.Id
	sh := getShards()
	m.ShardId = sh.GetShardNodeByKey(user.City).Id

	stmt, err := dbPool.Prepare(
		"INSERT INTO messages (shard_id,id,user_id_from,user_id_to,message,created_at) VALUES (" + strconv.Itoa(int(m.ShardId)) + ", ?, ?, ?, ?, NOW())")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.Id, m.UserIdFrom, m.UserIdTo, m.Message)
	if err != nil {
		return
	}

	lastID = m.Id

	return
}

func GetMessages(user users.User) (messages []Message, err error) {
	dbPool := db.OpenDB(db.MessagesShard1)

	sh := getShards()
	shardId := sh.GetShardNodeByKey(user.City).Id

	sqlStmt := `SELECT 
					*
				FROM 
					messages 
				WHERE 
					shard_id = ? AND user_id_from = ?`

	rows, err := dbPool.Query(sqlStmt, user.Id, shardId)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var message Message

		if err = rows.Err(); err != nil {
			return
		}

		err = rows.Scan(
			&message.Id,
			&message.ShardId,
			&message.UserIdFrom,
			&message.UserIdTo,
			&message.Message,
			&message.CreatedAt)

		messages = append(messages, message)
	}
	return
}
