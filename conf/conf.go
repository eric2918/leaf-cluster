package conf

var (
	LenStackBuf = 4096

	// log
	LogLevel string
	LogPath  string
	LogFlag  int

	// console
	ConsolePort   int
	ConsolePrompt string = "Leaf# "
	ProfilePath   string

	// cluster
	ListenAddr        string
	ConnAddrs         map[string]string
	PendingWriteNum   int
	HeartBeatInterval int

	ServerName string
)
