package chat

import (
	"net/http"

	"gopkg.in/olahol/melody.v1"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

// ChannelSelect 채팅
func ChannelSelect(c echo.Context) error {
	return c.Render(http.StatusOK, "channelSelect.html", nil)
}

// Chat 채팅 화면
func Chat(c echo.Context) error {
	name := c.Param("name")
	println(name)

	data := struct {
		name string
	}{
		name: name,
	}

	return c.Render(http.StatusOK, "chat.html", data)
}

// HandleRequest 요청 핸들
func HandleRequest(c echo.Context) error {
	m.HandleRequest(c.Response().Writer, c.Request())
	return nil
}

// HandleMessage 메시지 핸들
func HandleMessage(s *melody.Session, msg []byte) {
	m.BroadcastFilter(msg, func(q *melody.Session) bool {
		return q.Request.URL.Path == s.Request.URL.Path
	})
}
