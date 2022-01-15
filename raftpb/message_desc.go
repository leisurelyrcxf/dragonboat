package raftpb

import "fmt"

func (m *Message) String() string {
    return fmt.Sprintf("MessageType: %s, %d->%d, Term: %d, LogTerm: %d, LogIndex: %d", m.Type.String(), m.From, m.To, m.Term, m.LogTerm, m.LogIndex)
}
