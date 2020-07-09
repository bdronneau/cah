package handlers

import (
	"math/rand"
	"time"

	"cah/pkg/models"
	"cah/pkg/services"
)

// RoomStatus contain all main informatiions about a room
type RoomStatus struct {
	ID               uint64                      `json:"id" form:"id" query:"id"`
	Name             string                      `json:"name"`
	Status           string                      `json:"status" form:"status" query:"status"`
	Turn             uint64                      `json:"turn" form:"turn" query:"turn"`
	CurrentCard      *models.RoomsCardsDetailled `json:"current_card" form:"current_card" query:"current_card"`
	CurrentTurnCards *models.CountCards          `json:"current_turn_cards" form:"current_turn_cards" query:"current_turn_cards"`
	CurrentVotes     *models.CountVotes          `json:"current_votes" form:"current_votes" query:"current_votes"`
	CurrentResponse  *models.CardsDetailled      `json:"current_response" form:"current_response" query:"current_response"`
	UsersCount       int                         `json:"users_count" form:"users_count" query:"users_count"`
	UserJudge        *models.User                `json:"user_judge" form:"user_judge" query:"user_judge"`
}

const (
	numberOfTurns = 10
)

func (a app) shuffleDeck(cards []*models.Cards) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func (a app) judgeElection(users []*models.User) error {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })

	// Here we already check if there is enought players
	err := services.UserElected(*users[0])
	if err != nil {
		return err
	}

	return nil
}

func (a app) nextJudge(room models.Room) error {
	err := services.SetNextJudge(room)
	if err != nil {
		return err
	}

	return nil
}
