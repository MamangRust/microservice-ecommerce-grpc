package repository

import (
	db "github.com/MamangRust/microservice-ecommerce-pkg/database/schema"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type Repositories struct {
	MerchantQuery           MerchantQueryRepository
	MerchantBusinessQuery   MerchantBusinessQueryRepository
	MerchantBusinessCommand MerchantBusinessCommandRepository
}

func NewRepositories(DB *db.Queries, merchantQuery pb.MerchantQueryServiceClient) *Repositories {
	return &Repositories{
		MerchantQuery:           NewMerchantQueryRepository(merchantQuery),
		MerchantBusinessQuery:   NewMerchantBusinessQueryRepository(DB),
		MerchantBusinessCommand: NewMerchantBusinessCommandRepository(DB),
	}
}
