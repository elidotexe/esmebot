package handlers

import (
	"fmt"
	"math/rand"
	"time"

	c "github.com/elidotexe/esme/internal/bot/common"
	"github.com/elidotexe/esme/internal/logger"
	"github.com/elidotexe/esme/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const RavenexusID = -1001628672322 // Test chat
// const RavenexusID = -1001626631509 // Ravenexus

// Button constants
const VerifyUserButton = "verify"

type Handlers struct {
	bot     *tgbotapi.BotAPI
	logger  *logger.Logger
	storage storage.Storage
}

// Initialize initializes a new instance of the Handlers struct with the provided BotAPI
// and Logger, and returns it along with a nil error.
func Initialize(
	b *tgbotapi.BotAPI,
	logger *logger.Logger,
	s storage.Storage) (*Handlers, error) {
	return &Handlers{
		bot:     b,
		logger:  logger,
		storage: s,
	}, nil
}

// Button handlers

func (h *Handlers) ButtonQueryHandler(query *tgbotapi.CallbackQuery) {
	if query.Data == VerifyUserButton {
		h.VerifyButtonQueryHandler(query)
	}
}

func (h *Handlers) VerifyButtonQueryHandler(query *tgbotapi.CallbackQuery) {
	_, ok := h.storage.Exist(query.Message.Chat.ID, query.From.ID)
	if !ok {
		return
	}

	isUserHuman = true

	username := c.GetUsername(query.From)
	verifiedUserMsgText := getVerifiedUserMsgText(username)

	sentMsg, err := h.bot.Send(tgbotapi.NewMessage(RavenexusID, verifiedUserMsgText))
	if err != nil {
		h.logger.Error("Error sending message", err)
		return
	}

	h.storage.Remove(query.Message.Chat.ID, query.From.ID)

	go h.DeleteMessage(
		query.Message.Chat.ID,
		query.Message.MessageID,
		DeleteMsgDelayZeroMin)

	go h.DeleteMessage(RavenexusID, sentMsg.MessageID, DeleteMsgDelayThirty)
}

func getVerifiedUserMsgText(username string) string {
	var messages = []string{
		"Oh, %s! So you're not an alien after all! Welcome to the ravenexus, fellow earthling!👾",

		"Finally, a real-life human! %s, Welcome to the ravenexus, where humans and aliens unite " +
			"for epic adventures!👽",

		"Well, well, well, if it isn't a fellow human! %s! Welcome to the ravenexus, where we " +
			"gather to plot our domination of the avian world!🕉",

		"Look who's here, %s! Another human in the midst of our raven-centric universe! " +
			"Welcome to the ravenexus, where we unravel the mysteries of cawing and feathered " +
			"mischief together.👻",

		"Wait a second, %s! We've got a real human among us! Welcome to the ravenexus, where we " +
			"celebrate our bird-like nature and have a great time!🤖",

		"Well, well, well... %s! Look who's got the stamp of approval! Welcome to the " +
			"ravenexus, where we party like there's no tomorrow and leave all sanity at the " +
			"door. Brace yourself for a wild ride!🎉",

		"Congrats! You've officially passed the secret initiation test and made it into the " +
			"ravenexus! %s. Get ready for non-stop dancing, glitter showers, and glow stick " +
			"battles!✨",

		"You've unlocked the VIP entrance to the ravenexus! Prepare yourself for mind-blowing " +
			"beats, neon lights, and a dance floor that defies gravity. Welcome, party pro, %s!🙏",

		"Hold on tight, because you've just stepped into the wildest realm of electronic " +
			"music and endless euphoria – the ravenexus! %s, prepare for bass drops, laser beams, " +
			"and a community of party animals that'll make your head spin!🤯",

		"Buckle up, my friend, %s! You've ventured into the realm of infinite beats and " +
			"shimmering madness known as the ravenexus. Embrace the madness, let loose your " +
			"inner raver, and get ready to create legendary memories!🤪",

		"Greetings, Earthling %s! Prepare for an extraterrestrial experience as you step " +
			"foot into the ravenexus. Don't worry, we won't probe you (much). Join us in " +
			"celebrating the fusion of intergalactic rhythms and earthly revelry. Welcome to " +
			"our cosmic party hub, where humans and aliens shake their tentacles together!🛸",

		"Ahoy, Earthling %s! Brace yourself for a close encounter of the ravenexus kind. " +
			"We promise not to abduct you (for now). Join our celestial gathering where humans " +
			"mingle with aliens, all in the name of cosmic fun and mind-bending beats!👐",

		"Greetings, puny human %s! Welcome to the ravenexus, where the vibrations of " +
			"otherworldly music mingle with your feeble mortal senses. Prepare to have your " +
			"neural circuits rewired as we dance the cosmic waltz together!👋",

		"Attention, Earth specimen %s! You have arrived at the ravenexus, a dimension where " +
			"humans and extraterrestrials unite in the pursuit of epic revelry. Get ready for " +
			"an out-of-this-world experience that will abduct your senses and leave you begging " +
			"for more!🙌",

		"Welcome to the ravenexus, fragile human %s! Behold the nexus of ravens and " +
			"intergalactic beings, where we groove to tunes that resonate across the cosmos. " +
			"Prepare for an alien-infused party that will make your terrestrial shenanigans pale " +
			"in comparison!🤌",

		"Step right up to the ravenexus, your one-stop shop for all the juicy deets on upcoming " +
			"events! Forget secret handshakes %s, just grab a ticket and join the party parade!🎊",

		"Welcome to the ravenexus, %s! The magical realm of event intel! We've got the inside " +
			"scoop on all the wild happenings coming your way. Join us, or forever be doomed " +
			"to FOMO!🫶",

		"Behold the ravenexus, your portal to event enlightenment! %s prepare to be showered " +
			"with knowledge about the coolest parties in town. Join our club and never miss a " +
			"chance to get your groove on!🎶",

		"Calling all event enthusiasts! Enter the ravenexus, your ultimate destination for " +
			"event extravaganza. %s, we've got the lowdown on the hottest soirées, shindigs, and " +
			"shenanigans. Get ready to party like there's no tomorrow!👏",

		"Welcome to the ravenexus, your VIP access to the top-secret vault of event wizardry! " +
			"%s we've got the magical map that reveals all the epic gatherings. Join us, and " +
			"together we'll dance our way into party legend!🕺",
	}

	rand.Seed(time.Now().UnixNano())
	randMsg := messages[rand.Intn(len(messages))]

	return fmt.Sprintf(randMsg, username)
}
