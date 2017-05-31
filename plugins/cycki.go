package plugins

import (
	"context"
	"regexp"

	"github.com/arachnist/repost-plugin-server"
	"github.com/arachnist/repost-plugin-server/util"
)

var stripCycki *regexp.Regexp

func cycki(ctx context.Context, request rps.Request) (response rps.Response) {
	img, err := util.HTTPGetXpath("http://oboobs.ru/random/", "//img/@src")
	if err != nil {
		response.Ok = false
		response.Err = err.Error()
	} else {
		response.Ok = true
		response.Message = []string{"cycki", "(nsfw):", string(stripCycki.ReplaceAll([]byte(img), []byte("boobs")))}
	}

	return
}

func init() {
	stripCycki = regexp.MustCompile("/boobs_preview")
	rps.Register("cycki", cycki)
}
