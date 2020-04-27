package e6

var (
  Filter = false
  FilterScore = "2"
  Sample = false
)

type e621 struct {
	Posts []Post `json:"posts"`
}
//Post from e621
type Post struct {
	ID            int         `json:"id"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	File          File          `json:"file"`
	Preview       Preview       `json:"preview"`
	Sample        SampleSt        `json:"sample"`
	Score         Score         `json:"score"`
	Tags          Tags          `json:"tags,omitempty"`
	LockedTags    []interface{} `json:"locked_tags,omitempty"`
	ChangeSeq     int         `json:"change_seq,omitempty"`
	Flags         Flags         `json:"flags,omitempty"`
	Rating        string        `json:"rating"`
	FavCount      int         `json:"fav_count"`
	Sources       []string      `json:"sources,omitempty"`
	Pools         []interface{} `json:"pools,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
	ApproverID    int         `json:"approver_id"`
	UploaderID    int         `json:"uploader_id"`
	Description   string        `json:"description,omitempty"`
	CommentCount  int         `json:"comment_count"`
	IsFavorited   bool          `json:"is_favorited,omitempty"`
}
//File from e621.Post
type File struct {
	Width  int  `json:"width"`
	Height int  `json:"height"`
	EXT    string `json:"ext"`
	Size   int  `json:"size"`
	Md5    string `json:"md5"`
	URL    string `json:"url"`
}
//Flags from e621.Post
type Flags struct {
	Pending      bool `json:"pending,omitempty"`
	Flagged      bool `json:"flagged,omitempty"`
	NoteLocked   bool `json:"note_locked,omitempty"`
	StatusLocked bool `json:"status_locked,omitempty"`
	RatingLocked bool `json:"rating_locked,omitempty"`
	Deleted      bool `json:"deleted,omitempty"`
}
//Preview from e621.Post
type Preview struct {
	Width  int  `json:"width,omitempty"`
	Height int  `json:"height,omitempty"`
	URL    string `json:"url,omitempty"`
}
//Relationships from e621.Post
type Relationships struct {
	ParentID          interface{}   `json:"parent_id,omitempty"`
	HasChildren       bool          `json:"has_children,omitempty"`
	HasActiveChildren bool          `json:"has_active_children,omitempty"`
	Children          []interface{} `json:"children,omitempty"`
}
//Sample from e621.Post
type SampleSt struct {
	Has    bool   `json:"has,omitempty"`
	Height int  `json:"height,omitempty"`
	Width  int  `json:"width,omitempty"`
	URL    string `json:"url,omitempty"`
}
//Score from e621.Post
type Score struct {
	Up    int `json:"up"`
	Down  int `json:"down"`
	Total int `json:"total"`
}
//Tags from e621.Post
type Tags struct {
	General   []string      `json:"general,omitempty"`
	Species   []string      `json:"species,omitempty"`
	Character []string      `json:"character,omitempty"`
	Copyright []string      `json:"copyright,omitempty"`
	Artist    []string      `json:"artist,omitempty"`
	Invalid   []interface{} `json:"invalid,omitempty"`
	Lore      []interface{} `json:"lore,omitempty"`
	Meta      []interface{} `json:"meta,omitempty"`
}

// EImage - yes
type EImage struct {
	URL       string
	Page      string
	Artist    string
	Source    string
	Score     int
	Tags struct{
		General   []string
		Species   []string
		Character []string
	}
	Rating    string
	TimeStamp string
	ID 				int
	EXT				string
	SoundWarning bool
}
