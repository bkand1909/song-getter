package utils

import (
	"errors"
	"github.com/bkand1909/song-getter/models"
	"log"
	"strconv"
	"strings"
)

type ZingParser struct {
	Parser
}

type RawZingXml struct {
	Item []struct {
		Title        string `xml:"title"`
		Performer    string `xml:"performer"`
		Link         string `xml:"link"`
		Source       string `xml:"source"`
		Hq           string `xml:"hq"`
		Duration     int    `xml:"duration"`
		Lyric        string `xml:"lyric"`
		MvLink       string `xml:"mvlink"`
		BackImage    string `xml:"backimage"`
		AdParam      string `xml:"adparam"`
		ErrorCode    int    `xml:"errorcode"`
		ErrorMessage string `xml:"errormessage"`
		Type         string `xml:"type,attr"`
	} `xml:"item"`
	Duration string `xml:"duration"`
}

const (
	ERR_CANNOT_EXTRACT_DATA_XML_URL = "Cannot extract data-xml url"
	ERR_CANNOT_FETCH_DATA_XML       = "Cannot fetch data-xml"
	ERR_CANNOT_UNMARSHAL_DATA_XML   = "Cannot unmarshal data-xml"
)

func (p *ZingParser) ToSong(url string, html string) (error, *models.ZingSong) {
	// extract data-xml url
	xmlUrl := FindOneStringSubmatch(`data-xml="([^"]+)"`, html)
	if xmlUrl == "" {
		log.Println(ERR_CANNOT_EXTRACT_DATA_XML_URL)
		return errors.New(ERR_CANNOT_EXTRACT_DATA_XML_URL), nil
	}

	// get data-xml
	resp := GetUrl(xmlUrl, nil)
	if resp.Error != nil {
		log.Println(resp.Error.Error())
		return errors.New(ERR_CANNOT_FETCH_DATA_XML), nil
	}

	// decode song from xml
	rawXml := RawZingXml{}
	err := resp.XML(&rawXml, nil)
	if err != nil || len(rawXml.Item) != 1 {
		log.Println(ERR_CANNOT_UNMARSHAL_DATA_XML)
		return errors.New(ERR_CANNOT_UNMARSHAL_DATA_XML), nil
	}
	// log.Println(JsonIndentStruct(rawXml))

	item := rawXml.Item[0]
	song := models.ZingSong{
		Song: models.Song{
			Url:       url,
			Lyric:     strings.TrimSpace(item.Lyric),
			Performer: strings.TrimSpace(item.Performer),
			Source:    strings.TrimSpace(item.Source),
			Title:     strings.TrimSpace(item.Title),
			Type:      strings.TrimSpace(item.Type),
		},
		Hq:        strings.TrimSpace(item.Hq),
		BackImage: strings.TrimSpace(item.BackImage),
	}

	return err, &song
}

func (p *ZingParser) ToAlbum(url string, html string) (error, *models.ZingAlbum) {
	// extract data-xml url
	xmlUrl := FindOneStringSubmatch(`data-xml="([^"]+)"`, html)
	if xmlUrl == "" {
		log.Println(ERR_CANNOT_EXTRACT_DATA_XML_URL)
		return errors.New(ERR_CANNOT_EXTRACT_DATA_XML_URL), nil
	}

	// get data-xml
	resp := GetUrl(xmlUrl, nil)
	if resp.Error != nil {
		log.Println(resp.Error.Error())
		return errors.New(ERR_CANNOT_FETCH_DATA_XML), nil
	}

	// decode song from xml
	rawXml := RawZingXml{}
	err := resp.XML(&rawXml, nil)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ERR_CANNOT_UNMARSHAL_DATA_XML), nil
	}
	// log.Println(JsonIndentStruct(rawXml))

	album := models.ZingAlbum{}
	album.Song = make([]models.ZingSong, len(rawXml.Item))
	album.Duration, _ = strconv.Atoi(strings.TrimSpace(rawXml.Duration))
	album.Url = url
	album.Title = FindOneStringSubmatch(`<h1 class="txt-primary">([^<]+)</h1>`, html)
	for idx, item := range rawXml.Item {
		album.Song[idx] = models.ZingSong{
			Song: models.Song{
				Url:       url,
				Lyric:     strings.TrimSpace(item.Lyric),
				Performer: strings.TrimSpace(item.Performer),
				Source:    strings.TrimSpace(item.Source),
				Title:     strings.TrimSpace(item.Title),
				Type:      strings.TrimSpace(item.Type),
			},
			Hq:        strings.TrimSpace(item.Hq),
			BackImage: strings.TrimSpace(item.BackImage),
		}
	}

	return nil, &album
}
