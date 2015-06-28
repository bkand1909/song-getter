package utils

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/levigross/grequests"
	"github.com/rainycape/unidecode"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type M bson.M

func S(s interface{}) (res string) {
	defer func() {
		if r := recover(); r != nil {
			res = ""
		}
	}()
	res = s.(string)
	return res
}

func GetDocument(html string) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil
	} else {
		return doc
	}
}

func GetNodeHtml(node *goquery.Selection) string {
	s, err := node.Html()
	if err != nil {
		return ""
	} else {
		return s
	}
}

func FindAllStringMatch(pattern string, text string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindStringSubmatch(text)
}

func FindAllStringSubmatch(pattern string, text string) []string {
	re := regexp.MustCompile(pattern)
	ar := re.FindAllStringSubmatch(text, -1)
	s := make([]string, len(ar))
	for i, val := range ar {
		s[i] = val[1]
	}
	return s
}

func FindOneStringSubmatch(pattern string, text string) string {
	s := FindAllStringMatch(pattern, text)
	if len(s) >= 2 {
		return s[1]
	} else {
		return ""
	}
}

func SplitRegexp(pattern string, text string) []string {
	re := regexp.MustCompile(pattern)
	return re.Split(text, -1)
}

func ReplaceRegex(pattern string, text string, replace string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, replace)
}

func GetSiteName(link string) string {
	values, err := url.Parse(link)
	if err != nil {
		return ""
	} else {
		if strings.HasPrefix(values.Host, "www.") {
			return values.Host[4:]
		} else {
			return values.Host
		}
	}
}

func GetUrl(url string, ro *grequests.RequestOptions) *grequests.Response {
	var resp = grequests.Get(url, ro)
	return resp
}

func HttpGetHtml(url string) (html string) {
	defer func() {
		if r := recover(); r != nil {
			html = ""
		}
	}()
	res, _ := http.Get(url)
	data, _ := ioutil.ReadAll(res.Body)
	html = string(data)
	return
}

func JsonIndent(data []byte) string {
	var v M
	json.Unmarshal(data, &v)
	buf, _ := json.MarshalIndent(v, "", "    ")
	return string(buf)
}

func JsonIndentStruct(v interface{}) string {
	buf, _ := json.MarshalIndent(v, "", "    ")
	return string(buf)
}

func GetFileContent(filename string) (content string) {
	defer func() {
		if r := recover(); r != nil {
			content = ""
		}
	}()
	buf, _ := ioutil.ReadFile(filename)
	content = string(buf)
	return
}

func HttpPostForm(url string, params url.Values) ([]byte, error) {
	res, err := http.PostForm(url, params)
	if err != nil {
		return []byte{}, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func ToTitle(s string) string {
	s = unidecode.Unidecode(s)
	s = strings.ToLower(strings.Replace(s, " ", "_", -1))
	return s
}

func ToDDMMYYYY(timestamp int64) string {
	d := time.Unix(timestamp, 0)
	return d.Format("02/01/2006")
}
