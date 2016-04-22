package fetcher

import (
  "errors"
  "regexp"
)

const (
  USERNAME_REGEXP = "www.douban.com/people/([^/.]+)/?"
)

func FetchUsername(url string) (string, error) {
  re := regexp.MustCompile(USERNAME_REGEXP)
  res := re.FindStringSubmatch(url)

  if len(res) > 0 {
    return res[1], nil
  }

  return "", errors.New("not found")
}
