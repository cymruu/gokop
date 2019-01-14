package models

type Entry struct {
	ID            int       `json:"id"`
	Date          string    `json:"date"`
	Body          string    `json:"body"`
	Author        Author    `json:"author"`
	Blocked       bool      `json:"blocked"`
	Favorite      bool      `json:"favorite"`
	VoteCount     int       `json:"vote_count"`
	CommentsCount int       `json:"comments_count"`
	Comments      []Comment `json:"comments"`
	Status        string    `json:"status"`
	Embed         Embed     `json:"embed"`
	Survey        Survey    `json:"survey"`
	CanComment    bool      `json:"can_comment"`
	UserVote      int       `json:"user_vote"`
}

type Author struct {
	Login  string `json:"login"`
	Color  int64  `json:"color"`
	Sex    string `json:"sex"`
	Avatar string `json:"avatar"`
}

type Comment struct {
	ID        int    `json:"id"`
	Author    Author `json:"author"`
	Date      string `json:"date"`
	Body      string `json:"body"`
	Blocked   bool   `json:"blocked"`
	Favorite  bool   `json:"favorite"`
	VoteCount int    `json:"vote_count"`
	Status    string `json:"status"`
	UserVote  int    `json:"user_vote"`
	Embed     *Embed `json:"embed,omitempty"`
}

type Embed struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Source   string `json:"source"`
	Preview  string `json:"preview"`
	Plus18   bool   `json:"plus18"`
	Size     string `json:"size"`
	Animated bool   `json:"animated"`
	Ratio    int    `json:"ratio"`
}

type Survey struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	ID         float64 `json:"id"`
	Answer     string  `json:"answer"`
	Count      float64 `json:"count"`
	Percentage float64 `json:"percentage"`
}
