package test

import (
	"Chess/api"
	"Chess/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	expect       int    //测试预期结果
}

type room struct {
	RoomID int `json:"info"`
}

var (
	testToken token
	testRoom  room
)

// 注册测试
func TestRegister(t *testing.T) {
	database.InitDB()
	// 创建一个模拟的gin.Context上下文
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// 设置POST请求的参数
	username := "testuser"
	password := "testpassword"
	c.Request, _ = http.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username="+username+"&password="+password))
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 调用被测试的函数
	api.Register(c)

	// 检查响应状态码是否为200 OK
	if c.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, c.Writer.Status())
	}

}

// 登录测试
func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	database.InitDB()
	//创建上下文
	c, _ := gin.CreateTestContext(w)
	//设置get参数
	username := "testuser"
	password := "testpassword"
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	c.Request, _ = http.NewRequest(http.MethodPost, "/user/login", strings.NewReader(form.Encode()))
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//调用函数
	api.Login(c)
	//检查状态码
	if c.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, c.Writer.Status())
	} else {
		t.Log("login pass")
	}
	resp := w.Body
	respBody, _ := io.ReadAll(resp)
	err := json.Unmarshal(respBody, &testToken)
	if err != nil {
		t.Errorf("json反序列化错误，%v", err)
	}
}

// token刷新测试，同时测试token验证机制能否正常运行
func TestRefreshToken(t *testing.T) {
	testBadToken := token{
		RefreshToken: "test",
		expect:       400,
	}
	var testGroup []token
	testToken.expect = http.StatusOK
	testGroup = append(testGroup, testToken)    //加入生成的token
	testGroup = append(testGroup, testBadToken) //加入随便写的token
	//遍历测试集
	for i, test := range testGroup {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/user/login/refresh?refresh_token="+test.RefreshToken, nil)
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		//调用函数进行测试
		api.RefreshToken(c)
		//检查状态码
		if c.Writer.Status() != test.expect {
			t.Errorf("test error on the %d,Expected status %d but got %d", i+1, http.StatusOK, c.Writer.Status())
		}
	}

}

// 创建房间测试
func TestCreateRoom(t *testing.T) {
	testToken.expect = http.StatusOK //设置期望
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/room/create", nil)
	c.Request.Header.Set("Authorization", testToken.Token)
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	api.CreateRoom(c)
	if c.Writer.Status() != testToken.expect {
		t.Errorf("Expected status %d but got %d", http.StatusOK, c.Writer.Status())
	}
	resp := w.Body
	respBody, _ := io.ReadAll(resp)
	err := json.Unmarshal(respBody, &testRoom)
	if err != nil {
		t.Errorf("json反序列化错误，%v", err)
	}
}

// websocket与棋盘逻辑测试
func TestWebsocketAndChessLogic(t *testing.T) {
	////s := httptest.NewServer(testWebsocket())
	//
	//token1 := testToken.Token
	//token2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA3LTEyIDIwOjIzOjE1LjY3NzM3NDEgKzA4MDAgQ1NUIG09KzQzMzE2LjIwOTY0NjgwMSIsImlkIjoiMSJ9.oLoTxiJ9lGYMYdNGdlhe86eXbSiOWiGIPC268fnlbak"
	////模仿两名用户
	//
	//w1 := httptest.NewRecorder()
	//w2 := httptest.NewRecorder()
	////w3 := newRecorder{
	////	ResponseRecorder: nil,
	////	Hijacker:         nil,
	////}
	////w3.HeaderMap = make(http.Header)
	////w3.Body = new(bytes.Buffer)
	////w3.Code = 200
	////w3.Hijacker=
	//
	//c1, _ := gin.CreateTestContext(w1)
	//c2, _ := gin.CreateTestContext(w2)
	//
	//c1.Request, _ = http.NewRequest("GET", "ws://127.0.0.1:8080/room/connect?room_id="+strconv.Itoa(testRoom.RoomID), nil)
	//c1.Request.Header.Set("Authorization", token1)
	//c1.Request.Header.Set("Upgrade", "websocket")
	//c1.Request.Header.Set("Connection", "Upgrade")
	//c1.Request.Header.Set("Sec-WebSocket-Version", "13")
	////c1.Request.Header.Set("Sec-WebSocket-Key", "1")
	//c2.Request, _ = http.NewRequest("GET", "ws://127.0.0.1:8080/room/connect?room_id="+strconv.Itoa(testRoom.RoomID), nil)
	//c2.Request.Header.Set("Authorization", token2)
	//c2.Request.Header.Set("Upgrade", "websocket")
	//c2.Request.Header.Set("Connection", "Upgrade")
	//c2.Request.Header.Set("Sec-WebSocket-Version", "13")
	////c1.Request.Header.Set("Sec-WebSocket-Key", "2")
	////进行websocket连接
	//api.ConnectRoom(c1)
	//api.ConnectRoom(c2)
	//wsURL := "ws://127.0.0.1:8080/room/connect?room_id=" + strconv.Itoa(testRoom.RoomID)
	//conn1, err := websocket2.Dial(wsURL, "", "")
	//if err != nil {
	//	t.Fatalf("WebSocket connection failed: %v", err)
	//}
	//conn2, err := websocket2.Dial(wsURL, "", "")
	//if err != nil {
	//	t.Fatalf("WebSocket connection failed: %v", err)
	//}
	//defer conn1.Close()
	//defer conn2.Close()
	////messageReady := struct {
	////	Type int `json:"type"`
	////}{
	////	Type: 1,
	////}
	////err = conn1.WriteJSON(messageReady)
	//if err != nil {
	//	t.Errorf("WebSocket connection failed: %v", err)
	//}
	////err = conn2.WriteJSON(messageReady)
	//if err != nil {
	//	t.Errorf("WebSocket connection failed: %v", err)
	//}
	r := gin.Default()
	r.GET("/room/connect", api.ConnectRoom)
	r.Run()
}

//func TestWebsocketAndChessLogic(t *testing.T) {
//	wsURL := "ws://127.0.0.1:8080/room/connect?room_id=" + strconv.Itoa(testRoom.RoomID)
//	router1 := gin.New()
//	router1.GET("/room/connect", func(c *gin.Context) {
//		// 使用websocket.Upgrader进行WebSocket升级
//		c.Request.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA3LTEyIDIwOjIzOjE1LjY3NzM3NDEgKzA4MDAgQ1NUIG09KzQzMzE2LjIwOTY0NjgwMSIsImlkIjoiMSJ9.oLoTxiJ9lGYMYdNGdlhe86eXbSiOWiGIPC268fnlbak")
//		c.Request.Header.Set("Upgrade", "websocket")
//		c.Request.Header.Set("Connection", "Upgrade")
//		c.Request.Header.Set("Sec-WebSocket-Version", "13")
//		c.Request.Header.Set("Sec-WebSocket-Key", "1")
//		api.ConnectRoom(c)
//		conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
//		if err != nil {
//			t.Errorf(err.Error())
//		}
//		log.Println(conn)
//	})
//	router2 := gin.New()
//	go router2.GET("/room/connect", func(c *gin.Context) {
//		// 使用websocket.Upgrader进行WebSocket升级
//		c.Request.Header.Set("Authorization", testToken.Token)
//		c.Request.Header.Set("Upgrade", "websocket")
//		c.Request.Header.Set("Connection", "Upgrade")
//		c.Request.Header.Set("Sec-WebSocket-Version", "13")
//		c.Request.Header.Set("Sec-WebSocket-Key", "2")
//		api.ConnectRoom(c)
//	})
//	//token2 := testToken.Token
//	//token1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA3LTEyIDIwOjIzOjE1LjY3NzM3NDEgKzA4MDAgQ1NUIG09KzQzMzE2LjIwOTY0NjgwMSIsImlkIjoiMSJ9.oLoTxiJ9lGYMYdNGdlhe86eXbSiOWiGIPC268fnlbak"
//	//// 创建测试上下文和请求
//	//w1 := httptest.NewRecorder()
//	//w2 := httptest.NewRecorder()
//	//c1, _ := gin.CreateTestContext(w1)
//	//c2, _ := gin.CreateTestContext(w2)
//	//c1.Request, _ = http.NewRequest("GET", "/room/connect?room_id="+strconv.Itoa(testRoom.RoomID), nil)
//	//c1.Request.Header.Set("Authorization", token1)
//	//c2.Request, _ = http.NewRequest("GET", "/room/connect?room_id="+strconv.Itoa(testRoom.RoomID), nil)
//	//c2.Request.Header.Set("Authorization", token2)
//	//
//	//// 启动HTTP服务器
//	server1 := httptest.NewServer(router1)
//	server2 := httptest.NewServer(router2)
//	defer server1.Close()
//	defer server2.Close()
//
//	// 连接WebSocket
//	wsURL1 := "ws" + strings.TrimPrefix(server1.URL, "http") + "/room/connect?room_id=" + strconv.Itoa(testRoom.RoomID)
//	wsURL2 := "ws" + strings.TrimPrefix(server2.URL, "http") + "/room/connect?room_id=" + strconv.Itoa(testRoom.RoomID)
//	conn1, _, err := websocket.DefaultDialer.Dial(wsURL1, nil)
//	if err != nil {
//		t.Fatalf("WebSocket connection failed: %v", err)
//	}
//	conn2, _, err := websocket.DefaultDialer.Dial(wsURL2, nil)
//	if err != nil {
//		t.Fatalf("WebSocket connection failed: %v", err)
//	}
//	defer conn1.Close()
//	defer conn2.Close()
//
//	// 发送准备消息
//	messageReady := model.WebsocketMessage{
//		Type: 1,
//	}
//	err = conn1.WriteJSON(messageReady)
//	if err != nil {
//		t.Errorf("WebSocket write failed: %v", err)
//	}
//	err = conn2.WriteJSON(messageReady)
//	if err != nil {
//		t.Errorf("WebSocket write failed: %v", err)
//	}
//
//	//接收并验证准备消息
//	//_, resp1, err := conn1.ReadMessage()
//	//if err != nil {
//	//	t.Errorf("WebSocket read failed: %v", err)
//	//}
//	//_, resp2, err := conn2.ReadMessage()
//	//if err != nil {
//	//	t.Errorf("WebSocket read failed: %v", err)
//	//}
//	//验证接收到的消息是否符合预期
//	//
//	//发送移动消息等
//	//
//	//接收并验证响应消息等
//	//
//	//继续执行其他测试逻辑
//}
