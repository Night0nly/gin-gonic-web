package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/service"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/domain"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const ChiasenhacBaseURL = "http://search.chiasenhac.vn/search.php?s="

type ChiasenhacCollector struct {
	service.Collector
	baseURL string
}

func NewChiasenhacCollector() *ChiasenhacCollector {
	return &ChiasenhacCollector{
		baseURL: ChiasenhacBaseURL,
	}
}

func (c *ChiasenhacCollector) GetSongList(query string) (domain.SongList, error) {
	if query == ""{
		return nil, nil
	}
	doc := c.fetchFromInternet(c.baseURL + strings.Replace(query, " ", "%20", -1))
	songList := c.parseSearchResultDoc(*doc)
	if len(songList) < 1 {
		return nil, errors.New("Fail to parse data or no song found")
	}
	return songList, nil
}

func (c *ChiasenhacCollector) GetCollectorType() domain.CollectorType {
	return domain.CHIASENHAC
}

func (c *ChiasenhacCollector) fetchFromInternet(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		panic("Error making request.")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Status code error: " + strconv.Itoa(res.StatusCode) + res.Status)
	}

	doc, erro := goquery.NewDocumentFromReader(res.Body)
	if erro != nil {
		panic("Error fetching song")
	}
	return doc
}

func (c *ChiasenhacCollector) parseSearchResultDoc(document goquery.Document) [] *domain.Song {
	songsInfo := document.Find("table").First().Find("tr")
	songList := make([]*domain.Song, songsInfo.Length()-1, songsInfo.Length()-1)

	wg := sync.WaitGroup{}
	wg.Add(songsInfo.Length())

	songsInfo.Each(
		func(i int, selection *goquery.Selection) {
			go func(i int, selection *goquery.Selection) {
				defer wg.Done()
				if i != 0 {
					songInfo := selection.Find(".tenbh").First().Find("p")
					songQualityLength := selection.Find("span.gen")
					songQuality := songQualityLength.Find("span").Text()
					songLength := strings.Replace(songQualityLength.Text(), songQuality, "", -1)
					url, _ := selection.Find("a").First().Attr("href")

					song := domain.NewSong(
						songInfo.First().Text(),
						songInfo.Next().Text(),
						songLength,
						url,
						c.getDownloadList(url),
					)
					songList[i-1] = song
				}
			}(i, selection)
		})

	wg.Wait()
	return songList
}

func (c *ChiasenhacCollector) getDownloadList(url string) domain.DownloadURLList {
	downloadList := domain.EmptyDownloadURLList()

	doc := c.fetchFromInternet(strings.Replace(url, ".html", "_download.html", -1))

	doc.Find("#downloadlink2").
		Find("b").Eq(1).
		Find("a").Each(func(i int, selection *goquery.Selection) {
		switch downloadUrl, _ := selection.Attr("href"); selection.Find("span").Text() {
		case domain.LOSSLESS.String():
			downloadList.Add(domain.LOSSLESS, downloadUrl)
		case domain.KBPS320.String():
			downloadList.Add(domain.KBPS320, downloadUrl)
		case domain.KBPS192.String():
			downloadList.Add(domain.KBPS192, downloadUrl)
		case domain.KBPS128.String():
			downloadList.Add(domain.KBPS128, downloadUrl)
		case domain.HD1080P.String():
			downloadList.Add(domain.HD1080P, downloadUrl)
		case domain.HD720P.String():
			downloadList.Add(domain.HD720P, downloadUrl)
		}
	})

	return *downloadList
}

