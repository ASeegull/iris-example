package app

import (
	"github.com/ASeegull/iris-example/schema"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

func (h *Hotel) roomInfo(ctx iris.Context) {
	id := ctx.Params().Get("id")
	room := &schema.Room{}
	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	if err := c.Find(bson.M{"_id": id}).One(room); err != nil {
		ctx.Application().Logger().Errorf("failed to find room info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to update room info")
	}
	ctx.JSON(&room)
}

func (h *Hotel) listrooms(ctx iris.Context) {
	hotel := ctx.Params().Get("hotel")
	var rooms []schema.Room
	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	if err := c.Find(bson.M{"hotel": hotel}).All(rooms); err != nil {
		ctx.Application().Logger().Errorf("failed to retrive rooms info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to retrieve rooms info")
	}
	ctx.JSON(rooms)
}

func (h *Hotel) addRoom(ctx iris.Context) {
	room := &schema.Room{}
	ctx.ReadJSON(room)
	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	if err := c.Insert(room); err != nil {
		ctx.Application().Logger().Errorf("failed to add room info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to add room info")
	}
	ctx.JSON(room)
}

func (h *Hotel) editRoom(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var update schema.Room
	ctx.ReadJSON(&update)

	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	if err := c.Update(bson.M{"_id": id}, &update); err != nil {
		ctx.Application().Logger().Errorf("failed to update guest info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to update guest info")
	}

	var room schema.Room
	if err := c.Find(bson.M{"_id": id}).One(room); err != nil {
		ctx.Application().Logger().Errorf("failed to find guest info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to update guest info")
	}
	ctx.JSON(&room)
}

func (h *Hotel) deleteRoom(ctx iris.Context) {
	id := ctx.Params().Get("id")
	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	if err := c.Remove(bson.M{"_id": id}); err != nil {
		ctx.Application().Logger().Errorf("failed to remove room info: %+v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef("failed to remove room info")
	}
	ctx.Writef("Room info removed successfully")
}

func (h *Hotel) guestInfo(ctx iris.Context) {
	id := ctx.Params().Get("id")
	guest := &schema.Guest{}
	c := h.session.DB(schema.DBName).C(schema.GuestsCollection)
	if err := c.Find(bson.M{"_id": id}).One(guest); err != nil {
		ctx.Application().Logger().Errorf("failed to find guest info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to find guest info")
	}
	ctx.JSON(guest)
}

func (h *Hotel) newGuest(ctx iris.Context) {
	var guest schema.Guest
	ctx.ReadJSON(&guest)
	c := h.session.DB(schema.DBName).C(schema.GuestsCollection)
	if err := c.Insert(guest); err != nil {
		ctx.Application().Logger().Errorf("failed to add guest info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to add guest info")
	}
}

func (h *Hotel) updateGuestInfo(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var update schema.Guest
	ctx.ReadJSON(&update)

	c := h.session.DB(schema.DBName).C(schema.GuestsCollection)
	if err := c.Update(bson.M{"_id": id}, &update); err != nil {
		ctx.Application().Logger().Errorf("failed to update guest info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to update guest info")
	}

	var guest schema.Guest
	if err := c.Find(bson.M{"_id": id}).One(guest); err != nil {
		ctx.Application().Logger().Errorf("failed to find guest info: %+v", err)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Writef("Failed to update guest info")
	}
	ctx.JSON(&guest)
}

func (h *Hotel) removeGuestInfo(ctx iris.Context) {
	id := ctx.Params().Get("id")
	c := h.session.DB(schema.DBName).C(schema.GuestsCollection)
	if err := c.Remove(bson.M{"_id": id}); err != nil {
		ctx.Application().Logger().Errorf("failed to remove guest info: %+v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef("failed to remove guest info")
	}
	ctx.Writef("Guest info removed successfully")
}
