package person

import (
	"context"
	"sort"

	. "common/structs"
	"errors"

	"google.golang.org/api/iterator"

	"db"
)

func GetSubTeams(team string, role []string) (*[]Team, error) {
	var teams []Team
	for _, value := range role {
		team, err := GetSubTeam(team, value)
		if err != nil {
			return nil, err
		}
		teams = append(teams, *team)
	}
	return &teams, nil
}

func GetSubTeam(teamType string, role string) (*Team, error) {
	persons, err := GetAllPersonsForARole(teamType, role)
	if err != nil {
		return nil, err
	}

	team := Team{
		Name:    role,
		Members: *persons,
	}

	return &team, nil
}

func GetAllPersonsForARole(team string, role string) (*[]Person, error) {
	firestoreClient := db.GetFirestore()

	personIterator := firestoreClient.Collection("contacts").Doc(team).Collection(role).Documents(context.Background())

	if personIterator == nil {
		return nil, errors.New("Internal Server Error / No Members")
	}

	var persons []Person

	for {
		doc, err := personIterator.Next()

		if err == iterator.Done {
			sort.SliceStable(persons, func(i, j int) bool {
				return persons[i].Priority < persons[j].Priority
			})
			return &persons, nil
		}

		if err != nil {
			return nil, err
		}

		var person Person
		errInConversion := ConvertToPersonObject(doc, &person)

		if errInConversion != nil {
			return nil, errInConversion
		}
		persons = append(persons, person)
	}
}

func GetAllSubTeams(team string) (*[]Team, error) {
	firestoreClient := db.GetFirestore()

	collectionIterator := firestoreClient.Collection("contacts").Doc(team).Collections(context.Background())
	var teams []Team
	for {
		collRef, err := collectionIterator.Next()
		if err == iterator.Done {
			return &teams, nil
		}
		if err != nil {
			return nil, err
		}
		team, err := GetSubTeam(team, collRef.ID)
		if err != nil {
			return nil, err
		}
		teams = append(teams, *team)
	}
}

func GetNamesOfSubTeams(team string) (*[]string, error) {
	firestoreClient := db.GetFirestore()

	collectionIterator := firestoreClient.Collection("contacts").Doc(team).Collections(context.Background())
	var subTeams []string
	for {
		collRef, err := collectionIterator.Next()
		if err == iterator.Done {
			return &subTeams, nil
		}
		if err != nil {
			return nil, err
		}
		subTeams = append(subTeams, collRef.ID)
	}
}
