package minecraft

import (
	"discraft/collection"
	"fmt"
	"github.com/alteamc/minequery/ping"
	"log"
	"net/url"
	"strconv"
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
	res, err := ping.Ping(s.Hostname, s.Port)
	if err != nil {
		panic(err)
	}
	log.Printf("%s:%d ping OK: %v", s.Hostname, s.Port, res)

	var names collection.Strings
	for _, p := range res.Players.Sample {
		names = append(names, p.Name)
	}

	return names
}
