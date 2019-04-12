package service
//
//import (
//	"github.com/PuerkitoBio/goquery"
//	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/service"
//	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/domain"
//	"net/http"
//	"strconv"
//	"strings"
//	"sync"
//)
//
//const ZingBaseURL = ""
//
//type ZingCollector struct {
//	service.Collector
//	baseURL string
//}
//
//func NewZingCollector() *ZingCollector {
//	return &ZingCollector{
//		baseURL: ZingBaseURL,
//	}
//}
//
//func (z *ZingCollector) GetSongList(query string) (domain.SongList, error) {
//	if query == "" {
//		return nil, nil
//	}
//	doc := z.fetchFromInternet(z.baseURL + strings.Replace(query, " ", "%20", -1))
//	songList := z.parseSearchResultDoc(*doc)
//
//
//}
//
//func (z *ZingCollector) GetCollectorType() domain.CollectorType {
//	return domain.MP3ZING
//}
//
//func (z *ZingCollector) fetchFromInternet(url string) *goquery.Document {
//	res, err := http.Get(url)
//	if err != nil {
//		panic("Error making request")
//	}
//	defer res.Body.Close()
//	if res.StatusCode != 200 {
//		panic("Status code error: " + strconv.Itoa(res.StatusCode) + res.Status)
//	}
//
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		panic("Error fetching song")
//	}
//	return doc
//}
//
//func (z *ZingCollector) parseSearchResultDoc(document goquery.Document) [] *domain.Song {
//	songInfo := document.Find("table").First().Find("tr")
//	songList := make([]*domain.Song, songInfo.Length()-1, songInfo.Length()-1)
//
//	wg := sync.WaitGroup{}
//	wg.Add(songInfo.Length())
//
//	songInfo.Each(
//		func(i int, selection *goquery.Selection) {
//			go func(i int, selection *goquery.Selection) {
//				defer wg.Done()
//				if i != 0 {
//					songInfo:= selection.Find(".tenbh").First().Find("p")
//					songQuality := selection.Find("hom nay dau nam ong ay khong co viec lam ah. Hom nay minh lam gi bh nhi, kiem cai con viet bao cao")
//					songLength := strings.Replace("o day la file gi nhi")
//					url := selection.Find("meo hieu may anh zai nay dang lam gi")
//
//					song := domain.NewSong(
//						songInfo.Text(),
//						songInfo.Next().Text(),
//						songLength,
//						url,
//						ZingBaseURL)
//				}
//			}(i, selection)
//		})
//}