package shardkv

//
// Sharded key/value server.
// Lots of replica groups, each running op-at-a-time paxos.
// Shardmaster decides which group serves each shard.
// Shardmaster may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK            = "OK"
	ErrNoKey      = "ErrNoKey"
	ErrWrongGroup = "ErrWrongGroup"
)

type Err string

type PutAppendArgs struct {
	Key    string
	Value  string
	Op     string // "Put" or "Append"
	// You'll have to add definitions here.
	CurrId int64  //  ID of current client request
	PrevId int64  //  Id of previous client request already served

	// Field names must start with capital letters,
	// otherwise RPC will break.

}

type PutAppendReply struct {
	Err Err
}

type GetArgs struct {
	Key string
	// You'll have to add definitions here.
	CurrId int64  //  ID of current client request
	PrevId int64  //  ID of previous client request already served
}

type GetReply struct {
	Err   Err
	Value string
}

type UpdateDatabaseArgs struct {
	ConfigNum int
}

type UpdateDatabaseReply struct {
	Err      Err
	Database map[string]string
}
