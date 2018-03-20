package app

func (h *Hotel) listrooms(ctx iris.Context) {
	hotel := ctx.Params().Get("hotel")
	room := &schema.Room{}
	c := h.session.DB("elementary").C("rooms")
	err := c.Find(bson.M{"hotel": hotel}).One(room)
	if err != nil {
		log.Error(err)
	}
	ctx.JSON(room)
}

func (h *Hotel) addRoom(ctx iris.Context) {
	hotel := ctx.Params().Get("hotel")
	room := &schema.Room{}
	c := h.session.DB("elementary").C("rooms")
	err := c.Find(bson.M{"hotel": hotel}).One(room)
	if err != nil {
		log.Error(err)
	}
	ctx.JSON(room)
}