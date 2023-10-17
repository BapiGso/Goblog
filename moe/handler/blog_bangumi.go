package handler

import (
	"SMOE/moe/database"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"time"
)

// todo complete
// SubjectType
// CollectionType
const (
	// Book 书籍
	Book = 1
	// Animation 动画
	Animation = 2
	// Music 音乐
	Music = 3
	// Game 游戏
	Game = 4
	// ThreeD 三次元
	ThreeD = 5

	// WantToWatch 想看
	WantToWatch = 1
	// Watched 看过
	Watched = 2
	// Watching 在看
	Watching = 3
	// OnHold 搁置
	OnHold = 4
	// Dropped 抛弃
	Dropped = 5
)

type AutoGenerated struct {
	Data   []Data
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Data struct {
	UpdatedAt time.Time     `json:"updated_at"`
	Comment   interface{}   `json:"comment"`
	Tags      []interface{} `json:"tags"`
	Subject   struct {
		Date   string `json:"date"`
		Images struct {
			Small  string `json:"small"`
			Grid   string `json:"grid"`
			Large  string `json:"large"`
			Medium string `json:"medium"`
			Common string `json:"common"`
		} `json:"images"`
		Name         string `json:"name"`
		NameCn       string `json:"name_cn"`
		ShortSummary string `json:"short_summary"`
		Tags         []struct {
			Name  string `json:"name"`
			Count int    `json:"count"`
		} `json:"tags"`
		Score           float64 `json:"score"`
		Type            int     `json:"type"`
		ID              int     `json:"id"`
		Eps             int     `json:"eps"`
		Volumes         int     `json:"volumes"`
		CollectionTotal int     `json:"collection_total"`
		Rank            int     `json:"rank"`
	} `json:"subject"`
	SubjectID   int  `json:"subject_id"`
	VolStatus   int  `json:"vol_status"`
	EpStatus    int  `json:"ep_status"`
	SubjectType int  `json:"subject_type"`
	Type        int  `json:"type"`
	Rate        int  `json:"rate"`
	Private     bool `json:"private"`
}

func curlBGM(url string) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "trim21/bangumi-episode-ics (https://github.com/Trim21/bangumi-episode-calendar)")
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	m := AutoGenerated{}
	if err := json.Unmarshal(body, &m); err != nil {
		return err
	}
	finalData := [6][]Data{}
	for _, v := range m.Data {
		finalData[v.Type] = append(finalData[v.Type], v)
	}
	bgm = bgmCache{finalData, time.Now().Unix()}

	return err
}

type bgmCache struct {
	Data [6][]Data
	TTL  int64
}

var bgm = bgmCache{}

// Bangumi todo https://freefrontend.com/css-cards/
func Bangumi(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	ops, err := qpu.GetOption("Goplugin:BangumiList")
	if err != nil {
		return err
	}
	m := struct {
		UserID string
		AppID  string
	}{}
	if err := json.Unmarshal([]byte(ops), &m); err != nil {
		return err
	}
	//oldurl := "https://api.bgm.tv/user/" + m.UserID + "/collections/anime?app_id=" + m.AppID + "&max_results=99"
	newapi := "https://api.bgm.tv/v0/users/" + m.UserID + "/collections?subject_type=2&limit=100&offset=0"
	//每七天更新一下
	if time.Now().Unix()-bgm.TTL > 604800 {
		if err := curlBGM(newapi); err != nil {
			return err
		}
	}
	//return c.Render(200, "404.template", "维护中...")
	return c.Render(200, "page-bangumi.template", bgm)
}
