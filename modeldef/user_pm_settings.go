package modeldef

type UserPmSettings struct {
	Id                int
	Notify            bool  //whether notify the user when new message comes
	NewTorrent        bool  //whether show user new torrent when available
	MaxMessage        int   //how many messages are shown in mailbox per page(maximum:100)
	RecieveType       int   //bit1:admin bit2:friend bit3:others
	DeleteAfterReply  bool  //whether delete the message after replying
	SaveSendedMessage bool  //whether save the sended message into "sended" mailbox
	NewComment        bool  //Notify when user's torrent has new comment
	ReferedNotice     bool  //Notify when user's thread is being reffered
	User              *User `orm:"reverse(one)"`
}
