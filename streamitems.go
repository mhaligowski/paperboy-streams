package streams

type StreamItem struct {
	StreamItemId  string `datastore:"id"`;
	UserId        string;
	TargetId      string;
	Title         string;
	OrderSequence int64;
}
