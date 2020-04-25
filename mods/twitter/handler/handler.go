package twit

import (
  "net/url"
  "time"

  "github.com/ChimeraCoder/anaconda"
  log "github.com/Sirupsen/logrus"
)

var (
	DefaultConfig string
	CurrentConfg string
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
	Delay time.Duration
	All bool
	Lock bool
	PerChanInfo string
	ChanInfo []PerChanInfoStruct
)

func Twitter(global bool, chanID string) *anaconda.TwitterApi {
  if global {
    api := anaconda.NewTwitterApiWithCredentials(AccessToken, AccessTokenSecret, ConsumerKey, ConsumerSecret)
    return api
  }

  api := anaconda.NewTwitterApiWithCredentials(AccessToken, AccessTokenSecret, ConsumerKey, ConsumerSecret)
  return api
}

func Tweet(status string, v url.Values) (url string, err error) {
  tweet, err := Twitter(true, "").PostTweet(status, v)
  url = "https://twitter.com/"+tweet.User.ScreenName+"/status/"+tweet.IdStr
  Twitter(true, "").Close()
  return url, err
}


func Follow(twitterUser string) (string, error){
  user, err := Twitter(true, "").FollowUser(twitterUser)
  log.Println(err)
  Twitter(true, "").Close()
  return user.Name, err
}

func ChanPaser(chanID string) (){

}


type PerChanInfoStruct struct {
	ChanID string
	TwitterAccount string
	GetReplies bool //When I work in a checks
	GetLikes bool // ditto
	GetRetweets bool //ditto
	OwnerOnly bool
}
