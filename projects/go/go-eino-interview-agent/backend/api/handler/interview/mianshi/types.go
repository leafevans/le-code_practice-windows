package mianshi

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SessionManager struct {
	sessions map[string]*InterviewSession
	mu       sync.RWMutex
}

// InterviewSession 面试会话
type InterviewSession struct {
	SessionID     string
	UserID        uint
	RecordID      uint64
	ResumeID      int64
	HasResume     bool
	Query         string
	Type          string //综合面试，专项面试
	Domain        string //go，java，中间件
	Difficulty    string //校招，社招（专项面试没有区分）
	Status        string // 会话状态：active, paused, completed, failed
	CompanyName   string // 公司名称
	PositionName  string // 岗位名称
	AnswerChan    chan string
	CancelFunc    context.CancelFunc // 用于取消面试循环
	StartTime     time.Time
	LastActivity  time.Time
	QuestionCount int32 // 面试问题总数（用于统计）
}

var (
	globalSessionManager *SessionManager
	sessionManagerOnce   sync.Once
)

// GetGlobalSessionManager 获取全局会话管理器
func GetGlobalSessionManager() *SessionManager {
	sessionManagerOnce.Do(func() {
		globalSessionManager = &SessionManager{
			sessions: make(map[string]*InterviewSession),
		}
	})
	return globalSessionManager
}

// CreateSession 创建会话
func (sm *SessionManager) CreateSession(userID uint, recordID uint64, resumeID int64, hasResume bool, query string, interviewType string, domain string, difficulty string) *InterviewSession {
	return sm.CreateSessionWithDetails(userID, recordID, resumeID, hasResume, query, interviewType, domain, difficulty, "", "")
}

// CreateSessionWithDetails 创建会话（带公司和岗位信息）
func (sm *SessionManager) CreateSessionWithDetails(userID uint, recordID uint64, resumeID int64, hasResume bool, query string, interviewType string, domain string, difficulty string, companyName string, positionName string) *InterviewSession {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sessionID := generateSessionID()
	session := &InterviewSession{
		SessionID:    sessionID,
		UserID:       userID,
		RecordID:     recordID,
		ResumeID:     resumeID,
		HasResume:    hasResume,
		Query:        query,
		Type:         interviewType,
		Domain:       domain,
		Difficulty:   difficulty,
		Status:       "active",
		CompanyName:  companyName,
		PositionName: positionName,
		AnswerChan:   make(chan string, 1), // 缓冲区大小为1
		StartTime:    time.Now(),
		LastActivity: time.Now(),
	}

	sm.sessions[sessionID] = session
	return session
}

// GetSession 获取会话
func (sm *SessionManager) GetSession(sessionID string) *InterviewSession {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.sessions[sessionID]
}

// SubmitAnswer 提交答案
func (sm *SessionManager) SubmitAnswer(sessionID string, answer string) error {
	sm.mu.RLock()
	session := sm.sessions[sessionID]
	sm.mu.RUnlock()

	if session == nil {
		return fmt.Errorf("session not found")
	}

	session.LastActivity = time.Now()
	select {
	case session.AnswerChan <- answer:
		return nil
	default:
	}

	return nil
}

// GetAnswer 获取答案
func (sm *SessionManager) GetAnswer(sessionID string, timeout time.Duration) (string, error) {
	sm.mu.RLock()
	session := sm.sessions[sessionID]
	sm.mu.RUnlock()

	if session == nil {
		return "", fmt.Errorf("session not found")
	}

	select {
	case answer := <-session.AnswerChan:
		session.LastActivity = time.Now()
		return answer, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout waiting for answer")
	}
}

// ClearAnswer 清除答案
func (sm *SessionManager) ClearAnswer(sessionID string) {
	sm.mu.RLock()
	session := sm.sessions[sessionID]
	sm.mu.RUnlock()

	if session == nil {
		return
	}

	select {
	case <-session.AnswerChan:
	default:
	}
}

// DeleteSession 删除会话
func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, sessionID)
}

// generateSessionID 生成唯一的会话ID
func generateSessionID() string {
	return "session_" + time.Now().Format("20060102150405") + "_" + randomString(8)
}

// randomString 生成随机字符串
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
