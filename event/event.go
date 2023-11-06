package event

func (e *OpenebsEvent) CategoryStr() string {
	return e.Category
}

func (e *OpenebsEvent) ActionStr() string {
	return e.Action
}
