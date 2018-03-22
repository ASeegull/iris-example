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
	if err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(room); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to find room %v info: %+v")
	}
	ctx.JSON(&room)
}

func (h *Hotel) listrooms(ctx iris.Context) {
	hotel := ctx.Params().Get("hotel")
	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	n, err := c.Find(bson.M{"hotel": hotel}).Count()
	if err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to retrive rooms count of %s: %+v")
	}

	rooms := make([]*schema.Room, 0, n)
	if err := c.Find(bson.M{"hotel": hotel}).All(&rooms); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to retrive rooms info: %+v")
	}
	ctx.JSON(rooms)
}

func (h *Hotel) addRoom(ctx iris.Context) {
	room := &schema.Room{}
	ctx.ReadJSON(room)
	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	if err := c.Insert(room); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to add room info: %+v")
	}
	ctx.JSON(room)
}

func (h *Hotel) editRoom(ctx iris.Context) {
	id := ctx.Params().Get("id")

	update := &schema.Room{}
	ctx.ReadJSON(update)

	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	if err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, update); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to update guest info: %+v")
	}

	var room schema.Room
	if err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(room); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to find guest info: %+v")
	}
	ctx.JSON(&room)
}

func (h *Hotel) deleteRoom(ctx iris.Context) {
	id := ctx.Params().Get("id")
	c := h.session.DB(schema.DBName).C(schema.RoomsCollection)
	if err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to remove room info: %+v")
	}
	ctx.Writef("Room info removed successfully")
}

func (h *Hotel) guestInfo(ctx iris.Context) {
	id := ctx.Params().Get("id")
	guest := &schema.Guest{}
	c := h.session.DB(schema.DBName).C(schema.GuestsCollection)
	if err := c.Find(bson.M{"_id": id}).One(guest); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to add guest info with %s: %+v", id)
	}
	ctx.JSON(guest)
}

func (h *Hotel) newGuest(ctx iris.Context) {
	var guest schema.Guest
	ctx.ReadJSON(&guest)
	c := h.session.DB(schema.DBName).C(schema.GuestsCollection)
	if err := c.Insert(guest); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to add guest info: %+v")
	}
}

func (h *Hotel) updateGuestInfo(ctx iris.Context) {
	id := ctx.Params().Get("id")
	req := &schema.Guest{}
	ctx.ReadJSON(req)

	c := h.session.DB(schema.DBName).C(schema.GuestsCollection)
	if err := c.Update(bson.M{"_id": id}, bson.M{"$set": req}); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to update guest info: %+v")
	}

	var guest schema.Guest
	if err := c.Find(bson.M{"_id": id}).One(guest); err != nil {
		sendError(ctx, err, iris.StatusNotFound, "failed to find guest info: %+v")
	}
	ctx.JSON(&guest)
}

func (h *Hotel) removeGuestInfo(ctx iris.Context) {
	id := ctx.Params().Get("id")
	c := h.session.DB(schema.DBName).C(schema.GuestsCollection)
	if err := c.Remove(bson.M{"_id": id}); err != nil {
		sendError(ctx, err, iris.StatusInternalServerError, "failed to remove guest info: %+v")
	}
	ctx.Writef("Guest info removed successfully")
}

func sendError(ctx iris.Context, err error, code int, msg string, args ...interface{}) {
	ctx.Application().Logger().Errorf(msg, err)
	ctx.StatusCode(code)
	ctx.Writef(msg)
}
