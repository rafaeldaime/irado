package main

import (
	"time"
)

type User struct {
	UserId     string    `db:"userid" json:"userid"`       // *PK max: 20
	UserName   string    `db:"username" json:"username"`   // *UQ max: 20
	PicId      string    `db:"picid" json:"picid"`         // *PK max: 20
	FullName   string    `db:"fullname" json:"fullname"`   // *UQ max: 20
	LikeCount  int       `db:"likecount" json:"likecount"` // default: 0
	Creation   time.Time `db:"creation" json:"-"`          // *NN
	LastUpdate time.Time `db:"lastupdate" json:"-"`        // *NN
	Deleted    bool      `db:"deleted" json:"-"`           // default: 0
	Admin      bool      `db:"admin" json:"-"`             // default: 0
}

type Pic struct {
	PicId    string    `db:"picid" json:"picid"`       // *PK max: 20
	Creation time.Time `db:"creation" json:"creation"` // *NN
	Deleted  bool      `db:"deleted" json:"deleted"`   // default: 0
}

type Profile struct {
	ProfileId    string    `db:"profileid" json:"profileid"`       // *PK max: 40
	UserId       string    `db:"userid" json:"userid"`             // *FK max: 20
	UserName     string    `db:"username" json:"username"`         // *NN max: 20
	Email        string    `db:"email" json:"email"`               // *NN max: 40
	FullName     string    `db:"fullname" json:"fullname"`         // *NN max: 40
	Gender       string    `db:"gender" json:"gender"`             // *NN enum: male, female
	ProfileUrl   string    `db:"profileurl" json:"profileurl"`     // *NN max: 100
	Language     string    `db:"language" json:"language"`         //  max: 10, default: pt_BR
	Verified     bool      `db:"verified" json:"verified"`         // default: 0
	FirstName    string    `db:"firstname" json:"firstname"`       // *NN max: 20
	LastName     string    `db:"lastname" json:"lastname"`         // *NN max: 20
	SourceUpdate time.Time `db:"sourceupdate" json:"sourceupdate"` // *NN
	AccessToken  string    `db:"accesstoken" json:"accesstoken"`   // *NN max: 100
	RefreshToken string    `db:"refreshtoken" json:"refreshtoken"` // *NN max: 100
	Scope        string    `db:"scope" json:"scope"`               // *NN max: 40
	TokenExpiry  time.Time `db:"tokenexpiry" json:"tokenexpiry"`   // *NN Time token spiries
	Creation     time.Time `db:"creation" json:"creation"`         // *NN
	LastUpdate   time.Time `db:"lastupdate" json:"lastupdate"`     // *NN
}

type Token struct {
	TokenId  string    `db:"tokenid" json:"tokenid"` // *PK max: 20
	UserId   string    `db:"userid" json:"userid"`   // *FK max: 20
	Creation time.Time `db:"creation" json:"-"`      // *NN
}

type Channel struct {
	ChannelId   string `db:"channelid" json:"channelid"`     // *PK max: 20
	ChannelName string `db:"channelname" json:"channelname"` //  max: 20
	ChannelSlug string `db:"channelslug" json:"channelslug"` // *UQ max: 20 Index
	LikeCount   int    `db:"likecount" json:"likecount"`     // default: 0
}

// Save in public/img/ImageId-Size.png {"small", "medium", "large"}
type Image struct {
	ImageId  string    `db:"imageid" json:"imageid"`   // *PK max: 20
	MaxSize  string    `db:"maxsize" json:"maxsize"`   // *NN enum: {small, medium, large}
	Creation time.Time `db:"creation" json:"creation"` // *NN
	Deleted  bool      `db:"deleted" json:"deleted"`   // default: 0
}

type Url struct {
	UrlId     string    `db:"urlid" json:"urlid"`         // *PK max: 5
	FullUrl   string    `db:"fullurl" json:"fullurl"`     //  max: 20
	UserId    string    `db:"userid" json:"userid"`       // *FK max: 20
	Creation  time.Time `db:"creation" json:"creation"`   // *NN
	ViewCount int       `db:"viewcount" json:"viewcount"` // default: 0
}

type Content struct {
	ContentId   string    `db:"contentid" json:"contentid"`     // *PK max: 20
	UrlId       string    `db:"urlid" json:"ulid"`              // *FK max: 5
	ChannelId   string    `db:"channelid" json:"channelid"`     // *FK max: 20
	Title       string    `db:"title" json:"title"`             //  max: 255 (250)
	Slug        string    `db:"slug" json:"slug"`               //  max: 255 (250)
	Description string    `db:"description" json:"description"` //  max: 255
	UserId      string    `db:"userid" json:"userid"`           // *FK max: 20
	ImageId     string    `db:"imageid" json:"imageid"`         // *FK max: 20
	MaxSize     string    `db:"maxsize" json:"maxsize"`         // *NN enum: {small, medium, large}
	LikeCount   int       `db:"likecount" json:"likecount"`     // default: 0
	Creation    time.Time `db:"creation" json:"creation"`       // *NN
	LastUpdate  time.Time `db:"lastupdate" json:"lastupdate"`   // *NN
	Deleted     bool      `db:"deleted" json:"deleted"`         // default: 0
	FullUrl     string    `db:"-" json:"fullurl"`               // Not in Data Base
}