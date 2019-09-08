package eventNames
//
//import (
//	"../../db"
//	"context"
//)
//
//func GetAllEventName() (*[]eventNames, error) {
//	fireStoreClient := db.GetFirestore()
//
//	categoriesIterator, err := fireStoreClient.Collection("events").Doc("eventsName").Collections(context.Background()).GetAll()
//	if err != nil {
//		return nil, err
//	}
//
//	var arrayList []eventNames
//	for index :=0; index < len(categoriesIterator) ; index++ {
//		arr, err := GetEventByCategory(categoriesIterator[index].ID)
//
//		if err != nil {
//			return nil, err
//		} else {
//			arrayList = append(arrayList, *arr)
//		}
//	}
//
//	return &arrayList, nil
//}
