package streams

type StreamItem struct {
	StreamItemId  string `datastore:"id"`;
	UserId        string `datastore:"user_id"`;
	TargetId      string `datastore:"target_id"`;
	Title         string `datastore:"title"`;
	OrderSequence int64 `datastore:"order_sequence"`;
}
