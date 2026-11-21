package service

import (
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	keywordRequest "inventory-service/app/models/request/keyword"
)

type KeywordService struct {
	datastore dao.IKeywordDataStore
}

func NewKeywordService(datastore dao.IKeywordDataStore) IKeywordService {
	return &KeywordService{datastore: datastore}
}

func (k *KeywordService) Add(request *keywordRequest.AddRequest) (int64, error) {
	normalizeName(&request.Keyword.Name)
	return k.datastore.Add(request.Keyword)
}

func (k *KeywordService) GetById(request *keywordRequest.GetByIdRequest) (*model.Keyword, error) {
	return k.datastore.GetById(request.Id)
}

func (k *KeywordService) GetAll(request *keywordRequest.GetAllRequest) (map[string]*model.Keyword, error) {
	return k.datastore.GetAll()
}
