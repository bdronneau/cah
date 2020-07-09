package handlers

import (
	"database/sql"
	"fmt"

	"cah/pkg/models"
	"cah/pkg/services"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (a app) handleJudgeCards(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleJudgeCards",
	})

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Error("GetRoomName %v", err)
		return c.JSON(500, "Pouet")
	}

	judgeHand, err := services.ListCurrentPlayedCards(*room)
	if err != nil {
		contextLogger.Error("judgeHand %v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, judgeHand)
}

func (a app) handleListRoom(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleListRoom",
	})

	data, err := services.ListRoom()

	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, data)
}

func (a app) handlePostRoom(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handlePostRoom",
	})

	r := new(models.Room)
	if err := c.Bind(r); err != nil {
		return c.JSON(400, nil)
	}

	if len(r.Description) == 0 {
		return c.JSON(400, "Description can not be empty")
	}

	data, err := services.InsertRoom(r.Description)
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(201, data)
}

func (a app) handlePostTurn(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handlePostTurn",
	})

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Error("GetRoomName %v", err)
		return c.JSON(500, "Pouet")
	}

	users, err := services.ListUserByRoom(room.ID)
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(500, "Pouet")
	}

	err = services.DistributeRoomsUsersCards(users)
	if err != nil {
		contextLogger.Error("Distribute cards", err)
		return c.JSON(500, "Pouet")
	}

	bCard, err := services.PickNewBlackCard(*room)
	if err != nil {
		contextLogger.Error("Distribute rooms cards", err)
		return c.JSON(500, "Pouet")
	}

	err = a.nextJudge(*room)
	if err != nil {
		contextLogger.Error("Election of judge %v", err)
		return c.JSON(500, "Pouet")
	}

	room.Turn = room.Turn + 1
	// TODO export in service
	roomsCards := models.RoomsCards{
		RoomID: room.ID,
		CardID: bCard.ID,
		Turn:   room.Turn,
	}

	_, err = services.InsertRoomsCards(&roomsCards)
	if err != nil {
		contextLogger.Error("InsertRoomsCards %v", err)
		return c.JSON(500, "Pouet")
	}

	data, err := services.UpdateRoom(room)
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, data)
}

func (a app) handleRoomStop(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleRoomStop",
	})

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Errorf("GetRoomName %v", err)
		return c.JSON(500, "Pouet")
	}

	room.Status = "finished"

	room, err = services.UpdateRoom(room)
	if err != nil {
		contextLogger.Errorf("UpdateRoom %v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, nil)
}

func (a app) handleRoomStats(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleRoomHistory",
	})

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Error("GetRoomName %v", err)
		return c.JSON(500, "Pouet")
	}
	// status := make(map[string]interface{})
	status := stats{}

	status.History, err = services.GetHistory(*room)
	if err != nil {
		contextLogger.Error("GetHistory %v", err)
		return c.JSON(500, "Pouet")
	}

	status.Classement, err = services.GetClassementByRoom(*room)
	if err != nil {
		contextLogger.Error("GetClassementByRoom %v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, status)
}

func (a app) handleRoomStart(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleRoomStart",
	})

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Error("GetRoomName %v", err)
		return c.JSON(500, "Pouet")
	}

	users, err := services.ListUserByRoom(room.ID)
	if err != nil {
		contextLogger.Error("ListUserByRoom %v", err)
		return c.JSON(500, "Pouet")
	}

	// TODO limit cards by turn & users
	wCards, err := services.ListCardsByType("white")
	if err != nil {
		contextLogger.Error("ListCardsByType white %v", err)
		return c.JSON(500, "Pouet")
	}

	// TODO limit cards by turn
	bCards, err := services.ListCardsByType("black")
	if err != nil {
		contextLogger.Error("ListCardsByType blacck %v", err)
		return c.JSON(500, "Pouet")
	}

	a.shuffleDeck(wCards)
	a.shuffleDeck(bCards)

	err = a.judgeElection(users)
	if err != nil {
		contextLogger.Error("Election of judge %v", err)
		return c.JSON(500, "Pouet")
	}

	for i, v := range users {
		contextLogger.Debugf("%d takes card from %d to %d", v.Name, i*5, i*5+5-1)
		for j, cards := range wCards[i*5 : i*5+5] {
			onHand := false
			if j < 5 {
				onHand = true
			}
			contextLogger.Debugf("index %d with cardID %d", j, cards.ID)

			roomsUsersCards := &models.RoomsUsersCards{
				RoomID: room.ID,
				CardID: cards.ID,
				UserID: v.ID,
				OnHand: onHand,
			}

			_, err = services.InsertRoomsUsersCards(roomsUsersCards)
			if err != nil {
				contextLogger.Error("InsertRoomsUsersCards %v", err)
				return c.JSON(500, "Pouet")
			}
		}

	}

	// TODO export in service
	roomsCards := models.RoomsCards{
		RoomID: room.ID,
		CardID: bCards[0].ID,
		Turn:   1,
	}

	_, err = services.InsertRoomsCards(&roomsCards)
	if err != nil {
		contextLogger.Error("InsertRoomsCards %v", err)
		return c.JSON(500, "Pouet")
	}

	room.Status = "started"
	room.Turn = 1

	room, err = services.UpdateRoom(room)
	if err != nil {
		contextLogger.Error("UpdateRoom %v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, nil)
}

func (a app) handleRoomStatus(c echo.Context) error {
	logrus.Info("handling /handleRoomStatus request")
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleRoomStatus",
	})
	roomStatus := &RoomStatus{}

	room, err := services.GetRoomName(c.Param("name"))
	if err == sql.ErrNoRows {
		return c.JSON(404, "Room does not exist")
	} else if err != nil {
		contextLogger.Errorf("GetRoomName (%s) %v", room.ID, err)
		return c.JSON(500, "Retrieve room data")
	}

	roomStatus.ID = room.ID
	logrus.Info(fmt.Sprint(room))
	roomStatus.Status = room.Status
	roomStatus.Turn = room.Turn
	roomStatus.Name = room.Name

	if room.Status == "created" {
		return c.JSON(200, roomStatus)
	}

	roomCard, err := services.GetRoomsCardsByTurn(room)
	if err == sql.ErrNoRows {
		roomStatus.CurrentCard = &models.RoomsCardsDetailled{}
	} else if err != nil {
		contextLogger.Errorf("GetRoomsCardsByTurn %v", err)
		return c.JSON(500, "Retrieve roomCard data")
	} else {
		roomStatus.CurrentCard = roomCard
	}

	turnCards, err := services.GetCountCardsByTurn(room)
	if err == sql.ErrNoRows {
		roomStatus.CurrentTurnCards = &models.CountCards{}
	} else if err != nil {
		contextLogger.Errorf("GetCountCardsByTurn %v", err)
		return c.JSON(500, "Check logs")
	} else {
		roomStatus.CurrentTurnCards = turnCards
	}

	turnVotes, err := services.GetVotes(room)
	if err == sql.ErrNoRows {
		roomStatus.CurrentVotes = &models.CountVotes{}
	} else if err != nil {
		contextLogger.Errorf("GetVotes %v", err)
		return c.JSON(500, "Check logs")
	} else {
		roomStatus.CurrentVotes = turnVotes
	}

	responseCard, err := services.GetElected(room)
	if err == sql.ErrNoRows {
		roomStatus.CurrentResponse = &models.CardsDetailled{}
	} else if err != nil {
		contextLogger.Errorf("GetVotes %v", err)
		return c.JSON(500, "Check logs")
	} else {
		roomStatus.CurrentResponse = responseCard
	}

	users, err := services.ListUserByRoom(room.ID)
	if err != nil {
		contextLogger.Error("ListUserByRoom %v", err)
		return c.JSON(500, "Pouet")
	}

	roomStatus.UsersCount = len(users)

	judge, err := services.GetJudge(room)
	if err != nil {
		contextLogger.Errorf("GetJudge %v", err)
		return c.JSON(500, "Check logs")
	}

	roomStatus.UserJudge = judge

	return c.JSON(200, roomStatus)
}

func (a app) handleHeadRoomStatus(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleHeadRoomStatus",
	})

	_, err := services.GetRoomName(c.Param("name"))

	if err == sql.ErrNoRows {
		return c.JSON(404, "Room does not exist")
	} else if err != nil {
		contextLogger.Errorf("GetRoomName (%s) %v", c.Param("name"), err)
		return c.JSON(500, "Retrieve room data")
	}

	return c.JSON(200, nil)
}

// handleRoomSCards rerieve current black card
func (a app) handleRoomSCards(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleRoomSCards",
	})

	room, err := services.GetRoomName(c.Param("name"))

	if err != nil {
		contextLogger.Errorf("GetRoomName %v", err)
		return c.JSON(400, "Retrive room data")
	}

	roomCard, err := services.GetRoomsCardsByTurn(room)

	if err != nil {
		contextLogger.Errorf("GetRoomsCardsByTurn %v", err)
		return c.JSON(400, "Retrive roomCard data")
	}

	return c.JSON(200, roomCard)
}
