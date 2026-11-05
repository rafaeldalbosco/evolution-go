package message_model

import (
	"encoding/json"
	"time"
)

// HistoryMessage guarda as mensagens do histórico que o WhatsApp manda no pareamento (DATABASE_SAVE_HISTORY=true).
// O histórico só chega uma vez: guardado aqui, quem integra (ex.: o atrio-ia levando ao Chatwoot) importa depois,
// pelos últimos N dias. Message é o waE2E.Message em JSON, no mesmo formato do webhook.
type HistoryMessage struct {
	MessageID  string          `json:"message_id" gorm:"primaryKey"`
	InstanceID string          `json:"instance_id" gorm:"type:uuid;primaryKey;index:idx_history_instancia_data,priority:1"`
	ChatJID    string          `json:"chat_jid"`
	SenderJID  string          `json:"sender_jid"`
	PushName   string          `json:"push_name"`
	FromMe     bool            `json:"from_me"`
	IsGroup    bool            `json:"is_group"`
	Timestamp  time.Time       `json:"timestamp" gorm:"index:idx_history_instancia_data,priority:2"`
	Message    json.RawMessage `json:"message" gorm:"type:jsonb"`
}
