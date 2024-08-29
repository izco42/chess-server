package models

import (
	"gorm.io/gorm"
	"time"
)

type Match struct {
	gorm.Model
	WinnerId  int
	Player2id int
	Player3id int
	Player4id int
	Time      string
	Winner    User `gorm:"foreignKey:WinnerId"`
	Player2   User `gorm:"foreignKey:Player2id"`
	Player3   User `gorm:"foreignKey:Player3id"`
	Player4   User `gorm:"foreignKey:Player4id"`
}

type MatchResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	WinnerId  int       `json:"winnerId,omitempty"`
	Player2id int       `json:"player2Id,omitempty"`
	Player3id int       `json:"player3Id,omitempty"`
	Player4id int       `json:"player4Id,omitempty"`
	Time      string    `json:"time,omitempty"`
}

type MatchResponseUsername struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	WinnerUsername  string    `json:"winner_username"`
	Player2Username string    `json:"player2_username"`
	Player3Username string    `json:"player3_username"`
	Player4Username string    `json:"player4_username"`
	Time            string    `json:"time,omitempty"`
}
