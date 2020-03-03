package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/appengine"

	"strings"

	"github.com/google/go-github/github"
	"github.com/hugozhu/godingtalk"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

//Handle 处理Github的Webhook Events
func Handle(w http.ResponseWriter, r *http.Request) {
	var err error

	action := r.Header.Get("x-github-event")
	signature := r.Header.Get("x-hub-signature")
	// id := r.Header.Get("x-github-delivery")
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if !verifySignature([]byte(os.Getenv("GITHUB_WEBHOOK_SECRET")), signature, body) {
		http.Error(w, fmt.Sprintf("%v", "Invalid or empty signature"), http.StatusBadRequest)
		return
	}

	var event github.WebHookPayload
	json.Unmarshal(body, &event)

	context := appengine.NewContext(r)
	c := godingtalk.NewDingTalkClient(os.Getenv("CORP_ID"), os.Getenv("CORP_SECRET"))
	// c.HTTPClient = urlfetch.Client(context)
	// c.HTTPClient.Transport = &urlfetch.Transport{
	// 	Context: context,
	// 	AllowInvalidServerCertificate: true,
	// }
	c.Cache = NewZeroCache(context, ".access_token")
	refresh_token_error := c.RefreshAccessToken()

	msg := godingtalk.OAMessage{}
	msg.Head.Text = "Github"
	msg.Head.BgColor = "FF00AABB"
	switch action {
	case "push":
		msg.Body.Title = "[" + *event.Repo.Name + "] Push"
		msg.URL = *event.Compare
		for _, commit := range event.Commits {
			value := *commit.Message
			if len(commit.Added) > 0 {
				value = value + "\n Added: " + strings.Join(commit.Added, ", ")
			}
			if len(commit.Modified) > 0 {
				value = value + "\n Modified: " + strings.Join(commit.Modified, ", ")
			}
			if len(commit.Removed) > 0 {
				value = value + "\n Removed: " + strings.Join(commit.Removed, ", ")
			}
			msg.Body.Form = append(msg.Body.Form, godingtalk.OAMessageForm{
				Key:   "Commits: ",
				Value: value,
			})
		}
	case "watch":
		msg.Body.Title = "[" + *event.Repo.Name + "] Watch Updated"
		msg.URL = *event.Repo.HTMLURL
		msg.Body.Form = []godingtalk.OAMessageForm{
			{
				Key:   "Watchers: ",
				Value: fmt.Sprintf("%d", *event.Repo.WatchersCount),
			},
			{
				Key:   "Forks: ",
				Value: fmt.Sprintf("%d", *event.Repo.ForksCount),
			},
		}
	default:
		msg.Body.Title = "[" + *event.Repo.Name + "] " + action
		msg.URL = *event.Repo.HTMLURL
	}
	msg.Body.Author = *event.Sender.Login

	data, err := c.SendOAMessage(os.Getenv("SENDER_ID"), os.Getenv("CHAT_ID"), msg)
	if err != nil {
		http.Error(w, fmt.Sprintf("send error: %v %v", err, data), http.StatusInternalServerError)
	} else if refresh_token_error != nil {
		http.Error(w, fmt.Sprintf("refresh error: %v", refresh_token_error), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", data)
	}
}

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

func verifySignature(secret []byte, signature string, body []byte) bool {

	const signaturePrefix = "sha1="
	const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))

	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(signBody(secret, body), actual)
}

type ZeroCache struct {
}

func NewZeroCache(ctx context.Context, key string) *ZeroCache {
	return &ZeroCache{}
}

func (c *ZeroCache) Set(data godingtalk.Expirable) error {
	return nil
}

func (c *ZeroCache) Get(data godingtalk.Expirable) error {
	return errors.New("Not found")
}
