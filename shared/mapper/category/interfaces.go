package categoryapimapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/response"
)

type CategoryBaseResponseMapper interface {
	ToResponseCategory(category *pb.CategoryResponse) *response.CategoryResponse
	ToResponsesCategory(categories []*pb.CategoryResponse) []*response.CategoryResponse
}

type CategoryQueryResponseMapper interface {
	CategoryBaseResponseMapper
	ToApiResponseCategory(pbResponse *pb.ApiResponseCategory) *response.ApiResponseCategory
	ToApiResponsesCategory(pbResponse *pb.ApiResponsesCategory) *response.ApiResponsesCategory
	ToApiResponsePaginationCategory(pbResponse *pb.ApiResponsePaginationCategory) *response.ApiResponsePaginationCategory
	ToApiResponsePaginationCategoryDeleteAt(pbResponse *pb.ApiResponsePaginationCategoryDeleteAt) *response.ApiResponsePaginationCategoryDeleteAt
}

type CategoryCommandResponseMapper interface {
	CategoryBaseResponseMapper
	ToResponseCategoryDelete(category *pb.CategoryResponseDeleteAt) *response.CategoryResponseDeleteAt
	ToResponsesCategoryDeleteAt(categories []*pb.CategoryResponseDeleteAt) []*response.CategoryResponseDeleteAt
	ToApiResponseCategoryDeleteAt(pbResponse *pb.ApiResponseCategoryDeleteAt) *response.ApiResponseCategoryDeleteAt
	ToApiResponseCategory(pbResponse *pb.ApiResponseCategory) *response.ApiResponseCategory
	ToApiResponseCategoryDelete(pbResponse *pb.ApiResponseCategoryDelete) *response.ApiResponseCategoryDelete
	ToApiResponseCategoryAll(pbResponse *pb.ApiResponseCategoryAll) *response.ApiResponseCategoryAll
	ToApiResponsePaginationCategoryDeleteAt(pbResponse *pb.ApiResponsePaginationCategoryDeleteAt) *response.ApiResponsePaginationCategoryDeleteAt
}
