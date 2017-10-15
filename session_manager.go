package main

type SessionManager struct {
	subscribers *Subscribers
}

func (sm *SessionManager) StartSession(userId string) *Session {
	subscriber, exists := sm.subscribers.Get(userId)

	if !exists {
		subscriber = NewSubscriber(userId)
		session := subscriber.RegisterSession()
		sm.subscribers.Add(subscriber)
		subscriber.HandleBroadcast()
		return session
	}

	return subscriber.RegisterSession()
}

func (sm *SessionManager) EndSession(session *Session) {
	sessionId := session.Id
	subscriber := session.Subscriber
	subscriberId := subscriber.Id
	session.SafeClose()
	subscriber.SessionsCount--
	session.Subscriber.Sessions.Delete(sessionId)

	if subscriber.SessionsCount < 1 {
		subscriber.Subscription.Close()
		subscriber.BusConn.Close()
		sm.subscribers.Delete(subscriberId)
	}
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		subscribers: new(Subscribers),
	}
}
