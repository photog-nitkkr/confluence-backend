package event

func GetEventsName(categoryName string, eventNames []string) (*[]Event, error) {
	var events []Event

	for _, eventName := range eventNames {
		event, err := GetEventName(categoryName, eventName)

		if err != nil {
			return nil, err
		}

		events = append(events, *event)
	}
	return &events, nil
}
