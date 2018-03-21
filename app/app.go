package app

import (
	"context"
	"time"

	"github.com/ASeegull/iris-example/config"

	"github.com/kataras/iris"

	mgo "gopkg.in/mgo.v2"
)

// Hotel holds MongoDB session and is object reciever for router handlers
type Hotel struct {
	session *mgo.Session
}

// Start sets up iris application with default configuration,
// registers routes and custom handler on iterrupt (to close db connection)
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
	app.Get("/room/{id}", hotel.roomInfo)
	app.Post("/room/new", hotel.addRoom)
	app.Put("/room/{id}", hotel.editRoom)
	app.Delete("/room/{id}", hotel.deleteRoom)

	app.Get("/guest/{id}", hotel.guestInfo)
	app.Post("/guest/new", hotel.newGuest)
	app.Put("/guest/{id}", hotel.updateGuestInfo)
	app.Delete("/guest/{id}", hotel.removeGuestInfo)

	// Removes default interrupt handler as we declared a custom one above
	app.Run(iris.Addr(cfg.Addr()), iris.WithoutInterruptHandler)
}
