package app

import (
	"context"
	"time"

	"github.com/ASeegull/iris-example/config"
	"github.com/ASeegull/iris-example/schema"
	"github.com/kataras/iris"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Hotel holds MongoDB session and is object reciever for router handlers
type Hotel struct {
	session *mgo.Session
}

func Start(cfg *config.Settings, s *mgo.Session) {
	hotel := &Hotel{session: s}
	app := iris.Default()
	app.RegisterView(iris.HTML("./dist", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		defer hotel.session.Close()
		app.Shutdown(ctx)
	})
	app.Get("/{hotel:string}/listrooms", hotel.listrooms)
	app.Post("/addroom", hotel.addRoom)

	app.Run(iris.Addr(cfg.Addr()), iris.WithoutInterruptHandler)
}
