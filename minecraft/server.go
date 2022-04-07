package minecraft

import (
	"discraft/collection"
	"fmt"
	"github.com/alteamc/minequery/ping"
	"log"
	"net/url"
	"strconv"
	"time"
)

type Server struct {
	Hostname string
	Port     uint16
}

func NewServerFromURL(u url.URL) Server {
	parseInt, err := strconv.ParseUint(u.Port(), 10, 16)
	if err != nil {
		panic(fmt.Sprintf("'%s' is not a valid Port: %v", u.Port(), err))
	}

	return Server{
		Hostname: u.Hostname(),
		Port:     uint16(parseInt),
	}
}

func (s Server) CurrentPlayerList() collection.Strings {
	names := collection.Strings{}

	res, err := ping.PingWithTimeout(s.Hostname, s.Port, 10*time.Second)
	if err != nil {
		log.Printf("%s:%d ping FAILED: %v", s.Hostname, s.Port, err)

		return names
	}
	log.Printf("%s:%d ping ok: %v", s.Hostname, s.Port, res)

	for _, p := range res.Players.Sample {
		names = append(names, p.Name)
	}

	return names
}
