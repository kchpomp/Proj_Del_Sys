package session

import "github.com/kchpomp/Proj_Del_Sys/utils"

//Structure for storing username
type sessionData struct {
	Username string
	password string
}

//Structure to store a data in map with string key
type Session struct {
	data map[string]*sessionData
}

//Session object
func NewSession() *Session {
	// Object for new session
	s := new(Session)

	//Create not nil object
	s.data = make(map[string]*sessionData)

	//return object
	return s
}

//Session initialization method that receives username and password
func (s *Session) Init(username string, password string) string {
	//Generate session ID using utils
	sessionId := utils.GenerateId()

	//create data that will be stored in session
	data := &sessionData{Username: username, password: password}

	//write data in session ID
	s.data[sessionId] = data

	//return session ID
	return sessionId
}

func (s *Session) Get(sessionId string) (string, string) {
	data := s.data[sessionId]

	if data == nil {
		return "", ""
	}

	return data.Username, data.password
}
