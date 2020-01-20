package types

import (
	"time"
)

// PostID represents a numeric identifier that is used to uniquely identify a danbooru post
type PostID int

// TagString contains space-separated danbooru tags
type TagString string

// Post represents a single danbooru post
type Post struct {
	PostID       PostID `json:"id"`
	ParentPostID PostID `json:"parent_id"`
	UploaderID   UserID `json:"uploader_id"`
	ApproverID   UserID `json:"approver_id"`
	UploaderName string `json:"uploader_name"`

	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	LastNotedAt         time.Time `json:"last_noted_at"`
	LastCommentedAt     time.Time `json:"last_commented_at"`
	LastCommentBumbedAt time.Time `json:"last_comment_bumped_at"`

	Score         int    `json:"score"`
	Rating        Rating `json:"rating"`
	FavoriteCount int    `json:"fav_count"`
	UpScore       int    `json:"up_score"`
	DownScore     int    `json:"down_score"`

	TagString          TagString `json:"tag_string"`
	TagCount           int       `json:"tag_count"`
	TagStringGeneral   TagString `json:"tag_string_general"`
	TagCountGeneral    int       `json:"tag_count_general"`
	TagStringArtist    TagString `json:"tag_string_artist"`
	TagCountArtist     int       `json:"tag_count_artist"`
	TagStringCharacrer TagString `json:"tag_string_character"`
	TagCountCharacter  int       `json:"tag_count_character"`
	TagStringCopyright TagString `json:"tag_string_copyright"`
	TagCountCopyright  int       `json:"tag_count_copyright"`
	TagStringMeta      TagString `json:"tag_string_meta"`
	TagCountMeta       int       `json:"tag_count_meta"`

	PoolString string `json:"pool_string"`

	ImageHeight    int    `json:"image_height"`
	ImageWidth     int    `json:"image_width"`
	FileExtension  string `json:"file_ext"`
	FileSize       int    `json:"file_size"`
	FileURL        string `json:"file_url"`
	LargeFileURL   string `json:"large_file_url"`
	PreviewFileURL string `json:"preview_file_url"`
	Source         string `json:"source"`
	MD5            string `json:"md5"`

	HasChildren        bool `json:"has_children"`
	HasActiveChildren  bool `json:"has_active_children"`
	HasVisibleChildren bool `json:"has_visible_children"`
	IsFavorited        bool `json:"is_favorited"`
	IsPending          bool `json:"is_pending"`
	IsFlagged          bool `json:"is_flagged"`
	IsDeleted          bool `json:"is_deleted"`
	IsBanned           bool `json:"is_banned"`
	IsNoteLocked       bool `json:"is_note_locked"`
	IsRatingLocked     bool `json:"is_rating_locked"`
	IsStatusLocked     bool `json:"is_status_locked"`
}

// Posts is an ordered collection of Post instances
type Posts []Post
