package main

import (
	"crypto/tls"
	"flag"
	"net"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	"github.com/ASeegull/iris-example/app"
	"github.com/ASeegull/iris-example/config"
)

func main() {
	cfgPath := flag.String("config", "config/config.yaml", "Location of config File")
	dbPassword := flag.String("password", "", "Password to access database")

	flag.Parse()
	cfg, err := config.Parse(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	dialInfo, err := mgo.ParseURL(cfg.DBurl(*dbPassword))
	if err != nil {
		log.Error(err)
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		return tls.Dial("tcp", addr.String(), &tls.Config{InsecureSkipVerify: true})
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Error(err)
	}

	app.Start(cfg, session)
}
