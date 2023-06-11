package raftpb

import (
    "fmt"
    "strings"
)

func (m *Message) String() string {
    return fmt.Sprintf("MessageType: %s, %d->%d, Term: %d, LogTerm: %d, LogIndex: %d, Entries: %s",
        m.Type.String(), m.From, m.To, m.Term, m.LogTerm, m.LogIndex, m.EntriesDesc())
}

func (m *Message) EntriesDesc() string {
    if len(m.Entries) == 0 {
        return "[]"
    }
    descs := make([]string, len(m.Entries))
    for idx, entry := range m.Entries {
        descs[idx] = entry.String()
    }
    return fmt.Sprintf("%d[%s]", len(m.Entries), strings.Join(descs, ", "))
}

func (e *Entry) String() string {
    return fmt.Sprintf("Entry{Term: %d, Index: %d, Type: %s, Key: %d, Cmd: '%s'}", e.Term, e.Index, e.Type.String(), e.Key, string(e.Cmd))
}
