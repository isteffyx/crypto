package skc_session

type SessionDetailsStore struct {
	SessionId           []byte
	AttestationEvidence []byte
}

var SessionDetails *SessionDetailsStore
