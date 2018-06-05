package chat

import (
	"echo-boilerplate/helpers/jsonHelper"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

func Chat(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}

	return jsonHelper.Message(c, true, "완료")
}
