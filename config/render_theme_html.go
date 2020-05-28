package config

import (
	"bytes"
	"github.com/ChatFalcon/ChatFalcon/redis"
	"github.com/aws/aws-sdk-go/service/s3"
	"html/template"
	"io/ioutil"
	"strings"
	"time"
)

// Gets the theme HTML.
func getThemeHtml(c *ServerConfig) (string, error) {
	// Try and get from redis.
	cache := false
	if redis.Client != nil {
		s, err := redis.Client.Get("theme_html").Result()
		if err == nil {
			return s, err
		} else {
			if err == redis.Nil {
				cache = true
			} else {
				return s, err
			}
		}
	}

	// Try and get from S3.
	ses := c.S3Config.CreateS3Client()
	if c.CurrentTheme == "" {
		c.CurrentTheme = "default"
	}
	path := "themes/" + c.CurrentTheme + "/base.html"
	result, err := ses.GetObject(&s3.GetObjectInput{
		Bucket: &c.S3Config.Bucket,
		Key:    &path,
	})
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}
	if cache {
		err = redis.Client.Set("theme_html", b, time.Duration(0)).Err()
		if err != nil {
			return "", err
		}
	}
	return string(b), nil
}

// Gets the title.
func (c *ServerConfig) getTitle(PageTitle string) string {
	TitleFormat := c.TitleFormat
	if TitleFormat == "" {
		TitleFormat = "{name} - {page_title}"
	}
	return strings.ReplaceAll(
		strings.ReplaceAll(TitleFormat, "{name}", c.Name),
		"{page_title}", PageTitle)
}

// RenderThemeHTML is used to render the current themes HTML.
func (c *ServerConfig) RenderThemeHTML(URL, Title string) (string, error) {
	theme, err := getThemeHtml(c)
	if err != nil {
		return "", err
	}
	t, err := template.New("base").Parse(theme)
	if err != nil {
		return "", err
	}
	b := &bytes.Buffer{}
	u, err := c.S3Config.GenerateURL("themes/" + c.CurrentTheme)
	if err != nil {
		return "", err
	}
	err = t.Execute(b, map[string]interface{}{
		"Title":        c.getTitle(Title),
		"Keywords":     c.Keywords,
		"Description":  c.Description,
		"URL":          URL,
		"ThemePath":    u,
		"CustomHead":   template.HTML(c.CustomHead),
		"CustomBody":   template.HTML(c.CustomBody),
		"ServerConfig": c,
	})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
