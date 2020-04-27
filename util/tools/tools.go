package tools
import (
  "os"
  "bufio"
  "io"
  "net/http"
	"io/ioutil"
  "encoding/base64"

  "github.com/dokvis/goatbrotesquared/util/gvars"
)
//DirExists - Checks if dir exists
func DirExists(path string) (exists bool, dirErr error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

//FileToBase64 - Converts file to base64
func FileToBase64(file string) (base64file string) {
	f, _ := os.Open(file)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	base64file = base64.StdEncoding.EncodeToString(content)
	return base64file
}

//FileGetter Downloads files
func FileGetter(url string, file string) (err error) {
	mkfile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer mkfile.Close()

	/* Old http get
	data, err := http.Get(url)
	if err != nil {
		return err
	}
	defer data.Body.Close()
	*/
	client := &http.Client{}
	fileGet, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	fileGet.Header.Set("User-Agent", "GoatBroteSquared_DiscordGo_Bot/"+gvars.Version)
	fileResp, err := client.Do(fileGet)
	if err != nil {
		return err
	}
	defer fileResp.Body.Close()
	io.Copy(mkfile, fileResp.Body)
	return nil
}

//UniqueSilce - Removes dupes
func UniqueSilce(strSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range strSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}
