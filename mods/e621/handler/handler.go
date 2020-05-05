package e6

import (
  "encoding/json"
  "io/ioutil"
  "math/rand"
  "net/http"
  "strconv"
  "strings"
  "time"

  log "github.com/Sirupsen/logrus"

  "github.com/dokvis/goatbrotesquared/util/gvars"
)
//E621Handler - Handles e621 requests
func E621Handler(search string, forceID bool, forcesearch string, nsfw bool, nolewd bool, blacklist string) (eStuff EImage, err error) {
	var e621s e621
	search = strings.Replace(search, " ", "+", -1)
	//cub begone!
	search = strings.Replace(search, ";", "", -1)
	filter := ""
	eLink := ""
	if !forceID {
		if Filter {
			filter = "score:>="+FilterScore
		}
		eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+rating:s&limit=320&page="
		if nsfw {
			if !nolewd {
				 //log.Println("did a thing 1")
					eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+-cub+-young+-rating:s&limit=320&page="
			} else {
				  //log.Println("did a thing 1")
				  eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+rating:s&limit=320&page="
			}

		}
	} else {
			eLink = "https://e621.net/posts.json?tags="+search+forcesearch+"+rating:s&limit=320&page="
			if nsfw {
				if !nolewd {
					  //log.Println("did thing 1")
						eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+-cub+-young+-rating:s&limit=320&page="
				} else {
					 	//log.Println("did thing 2")
					  eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+rating:s&limit=320&page="
				}
			}
	}
	//log.Println("Json: "+eLink)
	rand.Seed(time.Now().UnixNano())
	//fileGetter(eLink, "temp/e621.json")
	client := &http.Client{}
	req, err := http.NewRequest("GET", eLink, nil)
	if err != nil {
		log.Println("Borked: Request Error: " + err.Error())
		return eStuff, err
	}
	req.Header.Set("User-Agent", "GoatBroteSquared_DiscordGo_Bot/"+gvars.Version)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Borked: Resp Error: " + err.Error())
		return eStuff, err
	}
	//log.Println(resp )
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Borked: Read Error:" + err.Error())
		return eStuff, err
	}
	if string(body) == "" {
		log.Println("Borked: body is blank")
		return eStuff, err
	}
	fuckMeErr := json.Unmarshal(body, &e621s)
	if fuckMeErr != nil {
			log.Println("Borked: unmarshaling failed" + fuckMeErr.Error())
			return eStuff, err
	}
	var maxE621 int
	maxE621 = len(e621s.Posts)
	if maxE621 == 0 {
		return eStuff, err
		}
	/*
		Fixes shit since rand can't be 0 and getting a result of 1, like with sans,
		getting rid of `-1` for rand still crashes, ITS FUCKING MAGIC I TELL YOU
	*/
	if maxE621 == 1 {
		maxE621 = 2
	}
	numE621 := rand.Intn(maxE621 - 1)

	eStuff.Rating = e621s.Posts[numE621].Rating
	//Having some issues with loading of images, using sample instead
	if !Sample {
		eStuff.URL = e621s.Posts[numE621].File.URL
	} else if e621s.Posts[numE621].File.EXT == "webm" {
		eStuff.URL = e621s.Posts[numE621].File.URL
	} else {
		eStuff.URL = e621s.Posts[numE621].Sample.URL
	}
	eStuff.URL = strings.Replace(eStuff.URL, " ", "%20", -1)
	eStuff.Tags.Character = e621s.Posts[numE621].Tags.Character
	eStuff.Tags.General = e621s.Posts[numE621].Tags.General
	eStuff.Tags.Species = e621s.Posts[numE621].Tags.Species
	eStuff.Score = e621s.Posts[numE621].Score.Total
	if nsfw {
		eStuff.Page = "https://e621.net/post/show/" + strconv.Itoa(e621s.Posts[numE621].ID)
	} else {
		eStuff.Page = "https://e926.net/post/show/" + strconv.Itoa(e621s.Posts[numE621].ID)
	}
	eStuff.SoundWarning = false
	for a := 0; a < len(e621s.Posts[numE621].Tags.Artist); a++ {
		if e621s.Posts[numE621].Tags.Artist[a] == "sound_warning" {
      e621s.Posts[numE621].Tags.Artist = append(e621s.Posts[numE621].Tags.Artist[:a], e621s.Posts[numE621].Tags.Artist[a+1:]...)
			eStuff.SoundWarning = true
		}
	}
  switch {
    case len(e621s.Posts[numE621].Tags.Artist) == 1: eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0]
    case len(e621s.Posts[numE621].Tags.Artist) == 2: eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0] + ",\n" + e621s.Posts[numE621].Tags.Artist[1]
    case len(e621s.Posts[numE621].Tags.Artist) > 2 : eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0] + ",\n**"+strconv.Itoa(len(e621s.Posts[numE621].Tags.Artist)-1) + "** ***more***"
    default: eStuff.Artist = "unknown artist"
  }
	/* if len(e621s.Posts[numE621].Tags.Artist) == 1 {
		eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0]
	} else if len(e621s.Posts[numE621].Tags.Artist) == 2 {
		if e621s.Posts[numE621].Tags.Artist[1] == "sound_warning" {
			eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0]
		} else {
			eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0] + ",\n" + e621s.Posts[numE621].Tags.Artist[1]
		}
	} else if len(e621s.Posts[numE621].Tags.Artist) > 2 {
		eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0] + ",\n**"+strconv.Itoa(len(e621s.Posts[numE621].Tags.Artist)-1) + "** ***more***"
	} else {
		eStuff.Artist = "unknown artist"
	}
  */
  eStuff.Artist = strings.Replace(eStuff.Artist, "_", " ", -1)
	if len(e621s.Posts[numE621].Sources) > 0 {
			eStuff.Source = e621s.Posts[numE621].Sources[0]
	}
	eStuff.TimeStamp = e621s.Posts[numE621].CreatedAt

	eStuff.EXT = e621s.Posts[numE621].File.EXT
	eStuff.ID = e621s.Posts[numE621].ID

	return eStuff, err
}
