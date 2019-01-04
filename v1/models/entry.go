package models

type Entry struct {
	ID                int64          `json:"id"`
	Author            string         `json:"author"`
	AuthorAvatar      string         `json:"author_avatar"`
	AuthorAvatarBig   string         `json:"author_avatar_big"`
	AuthorAvatarMed   string         `json:"author_avatar_med"`
	AuthorAvatarLo    string         `json:"author_avatar_lo"`
	AuthorGroup       uint16         `json:"author_group"`
	AuthorSex         string         `json:"author_sex"`
	Date              WykopTime      `json:"date"`
	Body              string         `json:"body"`
	Source            string         `json:"source"`
	URL               string         `json:"url"`
	Receiver          string         `json:"receiver"`
	ReceiverAvatar    string         `json:"receiver_avatar"`
	ReceiverAvatarBig string         `json:"receiver_avatar_big"`
	ReceiverAvatarMed string         `json:"receiver_avatar_med"`
	ReceiverAvatarLo  string         `json:"receiver_avatar_lo"`
	ReceiverGroup     string         `json:"receiver_group"`
	ReceiverSex       string         `json:"receiver_sex"`
	Comments          []EntryComment `json:"comments"`
	Blocked           bool           `json:"blocked"`
	VoteCount         int            `json:"vote_count"`
	UserVote          int            `json:"user_vote"`
	UserFavorite      bool           `json:"user_favorite"`
	Voters            []Voter        `json:"voters"`
	Type              string         `json:"type"`
	Embed             string         `json:"embed"`
	Deleted           bool           `json:"deleted"`
	ViolationURL      string         `json:"violation_url"`
	CanComment        bool           `json:"can_comment"`
	App               string         `json:"app"`
	CommentCount      int            `json:"comment_count"`
}

type EntryComment struct {
	ID              int       `json:"id"`
	Author          string    `json:"author"`
	AuthorAvatar    string    `json:"author_avatar"`
	AuthorAvatarBig string    `json:"author_avatar_big"`
	AuthorAvatarMed string    `json:"author_avatar_med"`
	AuthorAvatarLo  string    `json:"author_avatar_lo"`
	AuthorGroup     uint16    `json:"author_group"`
	AuthorSex       string    `json:"author_sex"`
	Date            WykopTime `json:"date"`
	Body            string    `json:"body"`
	Source          string    `json:"source"`
	EntryID         int       `json:"entry_id"`
	Blocked         bool      `json:"blocked"`
	Deleted         bool      `json:"deleted"`
	VoteCount       int       `json:"vote_count"`
	UserVote        int       `json:"user_vote"`
	Voters          []Voter   `json:"voters"`
	Embed           string    `json:"embed"`
	Type            string    `json:"type"`
	App             string    `json:"app"`
	ViolationURL    string    `json:"violation_url"`
}

type Voter struct {
	Author          string    `json:"author"`
	AuthorGroup     int       `json:"author_group"`
	AuthorAvatar    string    `json:"author_avatar"`
	AuthorAvatarBig string    `json:"author_avatar_big"`
	AuthorAvatarMed string    `json:"author_avatar_med"`
	AuthorAvatarLo  string    `json:"author_avatar_lo"`
	AuthorSex       string    `json:"author_sex"`
	Date            WykopTime `json:"date"`
}
