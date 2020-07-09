package handlers

import (
	"database/sql"
	"strconv"

	"cah/pkg/models"
	"cah/pkg/services"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (a app) handlePostUser(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handlePostUser",
	})

	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(400, nil)
	}

	data, err := services.InsertUser(u)
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(201, data)
}

func (a app) handleRoomPostUser(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handlePostUser",
	})

	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(400, err)
	}

	// TODO check if user already in room

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Errorf("%v", err)
		return c.JSON(400, "Retrieve room data")
	}

	u.RoomID = room.ID

	contextLogger.Debug(u)
	data, err := services.InsertRoomsUsers(u)
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(201, data)
}

func (a app) handleListUser(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleListUser",
	})

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(400, "Retrive room data")
	}

	data, err := services.ListUserByRoom(room.ID)
	if err != nil {
		contextLogger.Errorf("%v", err)
		return c.JSON(500, "ListUserByRoom")
	}

	return c.JSON(200, data)
}

func (a app) handleListUserCard(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleListUserCard",
	})

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 16)
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(400, "UserID incorrect")
	}

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(400, "Retrive room data")
	}

	data, err := services.ListUserCardDetailed(userID, room.ID)

	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, data)
}

func (a app) handlePostUserCard(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handlePostUserCard",
	})

	u := new(models.RoomsUsersCards)
	if err := c.Bind(u); err != nil {
		return c.JSON(400, "Bad request POST")
	}

	// TODO Check room exist
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if userID != u.UserID {
		return c.JSON(400, "Bad request user_id")
	}

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(400, "Retrive room data")
	}

	dataDB, err := services.GetRoomsUsersCardsSimpleByCardID(u.CardID, room.ID)
	if err != nil {
		contextLogger.Error("GetRoomsUsersCardsSimpleByCardID %v", err)
		return c.JSON(500, "Pouet")
	}

	if userID != dataDB.UserID || dataDB.Used || !dataDB.OnHand {
		return c.JSON(400, "Bad request (DB Check)")
	}

	dataDB.Turn = room.Turn
	dataDB.Used = true

	data, err := services.UpdateRoomsUsersCards(dataDB)
	if err != nil {
		contextLogger.Error("UpdateRoomsUsersCards %v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, data)
}

func (a app) handlePostVote(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handlePostVote",
	})

	u := new(models.RoomsUsersCards)
	if err := c.Bind(u); err != nil {
		return c.JSON(400, "Bad request")
	}

	room, err := services.GetRoomName(c.Param("name"))
	if err != nil {
		contextLogger.Error("%v", err)
		return c.JSON(400, "Retrive room data")
	}

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)

	if userID != u.UserID {
		return c.JSON(400, "Bad request")
	}

	card, err := services.GetRoomsUsersCardsByCardID(u.CardID, room.ID)
	if err != nil {
		contextLogger.Errorf("GetRoomsUsersCardsByCardID %v", err)
		return c.JSON(500, "Pouet")
	}

	if card.UserID == userID {
		return c.JSON(400, "Can not vote for your own card")
	}

	contextLogger.Infof("card turn", card.Turn)
	contextLogger.Infof("card used", card.Used)
	contextLogger.Infof("room", room.Turn)
	if !card.Used || room.Turn != card.Turn {
		return c.JSON(400, "Card does not meet requirements")
	}

	data, err := services.VodeRoomsUsersCards(&models.RoomsUsersCards{UserID: card.UserID, CardID: u.CardID, RoomID: room.ID})

	if err != nil {
		contextLogger.Errorf("VodeRoomsUsersCards %v", err)
		return c.JSON(500, "Pouet")
	}

	return c.JSON(200, data)
}

func (a app) handleGetUser(c echo.Context) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"route": "handleGetUser",
	})

	userID, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		contextLogger.Errorf("%v", err)
		return c.JSON(400, "User ID incorrect")
	}

	user, err := services.GetUser(userID)
	if err == sql.ErrNoRows {
		return c.JSON(404, "User not found")
	} else if err != nil {
		contextLogger.Errorf("GetUser %v", err)
		return c.JSON(500, "Check logs")
	} else {
		return c.JSON(200, user)
	}
}
