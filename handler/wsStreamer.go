package handler

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type WsStreamer struct {
	buff []byte
	cond *sync.Cond

	sync.Mutex

	conn *websocket.Conn
}

func (h *WsStreamer) Run() {
	for {
		_, msg, err := h.conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		h.Lock()
		h.buff = append(h.buff, msg...)
		h.cond.Signal()
		h.Unlock()
	}
}

func (h *WsStreamer) Read(b []byte) (size int, err error) {
	h.Lock()
	for len(h.buff) == 0 {
		h.cond.Wait()
	}
	size = copy(b, h.buff)
	h.buff = h.buff[size:]
	h.Unlock()
	return
}

func (h *WsStreamer) Write(b []byte) (size int, err error) {
	size = len(b)
	fmt.Println(string(b))
	err = h.conn.WriteMessage(websocket.TextMessage, b)
	return
}
