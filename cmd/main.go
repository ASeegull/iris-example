package main

import (
	"context"
	"crypto/tls"
	"flag"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ASeegull/iris-example/config"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Message struct {
	Message string `bson:"message"`
}

func main() {
	cfgPath := flag.String("config", "config/config.yaml", "Location of config File")
	dbPassword := flag.String("password", "", "Password to access database")

	flag.Parse()
	cfg, err := config.Parse(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	app := iris.Default()
	app.RegisterView(iris.HTML("./dist", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	dialInfo, err := mgo.ParseURL(cfg.DBurl(*dbPassword))
	if err != nil {
		log.Error(err)
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		return tls.Dial("tcp", addr.String(), tlsConfig)
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Error(err)
	}
	c := session.DB("test").C("table")

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		defer session.Close()
		app.Shutdown(ctx)
	})
	app.Get("/hello", func(ctx iris.Context) {
		result := &Message{}
		err := c.Find(bson.M{}).One(result)
		if err != nil {
			log.Error(err)
		}
		ctx.JSON(result)
	})

	app.Run(iris.Addr(cfg.Addr()), iris.WithoutInterruptHandler)
}
