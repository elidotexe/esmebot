package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	c "github.com/elidotexe/esme/internal/bot/common"
	"github.com/elidotexe/esme/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var goabaseAPI string

func (h *Handlers) OnInfoCommand(m *tgbotapi.Message) {
	username := c.GetUsername(m.From)
	chatID := m.Chat.ID

	goabaseAPI = "https://www.goabase.net/api/party/json/?country=united%20kingdom"

	cmdArgs := strings.ToLower(m.CommandArguments())
	if cmdArgs != "" {
		escapedUrl := "https://www.goabase.net/api/party/json/?country=" +
			"united%20kingdom&search="
		goabaseAPI = fmt.Sprintf("%s%s", escapedUrl, cmdArgs)
	}

	resp, err := h.GetAPIResponse(goabaseAPI)
	if err != nil {
		h.logger.Errorf("Error getting response: %v", err)
		return
	}

	h.logger.Infof("Info command received from %v", username)

	var pr models.PartyList
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&pr)
	if err != nil {
		h.logger.Errorf("Error decoding response: %v", err)
		return
	}

	counter := 0

	var msgText string
	msgText += fmt.Sprintf("4u %s 👽❤️\n\n", username)

	if len(pr.PartyList) == 0 {
		msgText = fmt.Sprintf("Sorry %s, "+
			"I couldn't find any events for %s 😿", username, cmdArgs)
	}

	for _, e := range pr.PartyList {
		msgText += getMsgText(e)

		counter++
		if counter == 5 {
			break
		}
	}

	msg := tgbotapi.NewMessage(chatID, msgText)
	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = true

	sentMsg, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Errorf("Error sending message: %v", err)
		return
	}

	go h.DeleteMessage(chatID, sentMsg.MessageID, DeleteMsgDelayFiveMin)

	go h.DeleteMessage(chatID, m.MessageID, DeleteMsgDelayFiveMin)
}

func (h *Handlers) GetAPIResponse(goabaseAPI string) (*http.Response, error) {
	req, err := http.NewRequest("GET", goabaseAPI, nil)
	if err != nil {
		h.logger.Errorf("Error creating request: %v", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.logger.Errorf("Error sending request: %v", err)
		return nil, err
	}

	return resp, nil
}

func getMsgText(e models.Event) string {
	return fmt.Sprintf("*Name*: %s\n"+
		"*Date*: %s\n"+
		"*Type*: %s\n"+
		"*Genre*: Psytrance\n"+
		"*City*: %s\n"+
		"*Country*: %s\n"+
		"*Organiser*: %s\n"+
		"*Link*: [More Info](%s) 🔗\n\n",
		e.Name,
		stripoutDate(e.Date),
		e.Type,
		e.City,
		e.Country,
		stripoutOrganiser(e.Organiser),
		getEventLink(e))
}

func stripoutDate(date string) string {
	dateStr := strings.Split(date, "T")[0]
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "Unknown"
	}

	return t.Format("02.01.2006")
}

func stripoutOrganiser(organiser string) string {
	re := regexp.MustCompile(`[_\.]`)
	return re.ReplaceAllString(organiser, " ")
}

func getEventLink(e models.Event) string {
	const SkiddleURL = "https://www.skiddle.com/"
	const Tag = "?sktag=15145"

	// Will be removed when the event is over
	if e.ID == 110366 {
		return fmt.Sprintf("%s%d%s", SkiddleURL+"e/", 36327598, Tag)
	}
	// Will be removed when the event is over
	if e.ID == 109820 {
		return fmt.Sprintf("%s%d%s", SkiddleURL+"e/", 36224871, Tag)
	}

	if e.URL != "" {
		if strings.Contains(e.URL, "skiddle") {
			re := regexp.MustCompile(`\d{8}`)
			matchID := re.FindString(e.URL)

			return fmt.Sprintf("%s%s%s", SkiddleURL+"e/", matchID, Tag)
		}

		url := removeAfterR(e.URL)
		if !strings.Contains(url, "skiddle") && strings.Contains(e.Organiser, "Boogie Woogie") {
			return fmt.Sprintf("%s%s",
				SkiddleURL+"whats-on/events/all/?keyword=boogie+woogie&hidecancelled=1", Tag)
		} else if !strings.Contains(url, "skiddle") && strings.Contains(e.Organiser, "Lu") {
			return fmt.Sprintf("%s%s",
				SkiddleURL+"whats-on/events/all/?keyword=psy+gypsies&hidecancelled=1", Tag)
		}

		return url
	}

	return removeAfterR(e.GoabaseURL)
}

func removeAfterR(input string) string {
	idx := strings.Index(input, "\r")
	if idx == -1 {
		return input
	}

	return input[:idx]
}
