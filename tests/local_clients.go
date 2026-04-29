package tests

import (
	"context"

	role_handler "github.com/MamangRust/microservice-ecommerce-grpc-role/handler"
	user_handler "github.com/MamangRust/microservice-ecommerce-grpc-user/handler"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LocalUserClient struct {
	Handler *user_handler.Handler
}

// UserCommandServiceClient implementation
func (c *LocalUserClient) Create(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	return c.Handler.UserCommand.Create(ctx, in)
}

func (c *LocalUserClient) Update(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	return c.Handler.UserCommand.Update(ctx, in)
}

func (c *LocalUserClient) UpdateIsVerified(ctx context.Context, in *pb.UpdateUserIsVerifiedRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	return c.Handler.UserCommand.UpdateIsVerified(ctx, in)
}

func (c *LocalUserClient) UpdatePassword(ctx context.Context, in *pb.UpdateUserPasswordRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	return c.Handler.UserCommand.UpdatePassword(ctx, in)
}

func (c *LocalUserClient) TrashedUser(ctx context.Context, in *pb.FindByIdUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUserDeleteAt, error) {
	return c.Handler.UserCommand.TrashedUser(ctx, in)
}

func (c *LocalUserClient) RestoreUser(ctx context.Context, in *pb.FindByIdUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUserDeleteAt, error) {
	return c.Handler.UserCommand.RestoreUser(ctx, in)
}

func (c *LocalUserClient) DeleteUserPermanent(ctx context.Context, in *pb.FindByIdUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUserDelete, error) {
	return c.Handler.UserCommand.DeleteUserPermanent(ctx, in)
}

func (c *LocalUserClient) RestoreAllUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.ApiResponseUserAll, error) {
	return c.Handler.UserCommand.RestoreAllUser(ctx, in)
}

func (c *LocalUserClient) DeleteAllUserPermanent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.ApiResponseUserAll, error) {
	return c.Handler.UserCommand.DeleteAllUserPermanent(ctx, in)
}

// UserQueryServiceClient implementation
func (c *LocalUserClient) FindAll(ctx context.Context, in *pb.FindAllUserRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationUser, error) {
	return c.Handler.UserQuery.FindAll(ctx, in)
}

func (c *LocalUserClient) FindById(ctx context.Context, in *pb.FindByIdUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	return c.Handler.UserQuery.FindById(ctx, in)
}

func (c *LocalUserClient) FindByEmail(ctx context.Context, in *pb.FindByEmailRequest, opts ...grpc.CallOption) (*pb.ApiResponseUserWithPassword, error) {
	return c.Handler.UserQuery.FindByEmail(ctx, in)
}

func (c *LocalUserClient) FindByVerificationCode(ctx context.Context, in *pb.FindByVerificationCodeRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	return c.Handler.UserQuery.FindByVerificationCode(ctx, in)
}

func (c *LocalUserClient) FindByActive(ctx context.Context, in *pb.FindAllUserRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationUserDeleteAt, error) {
	return c.Handler.UserQuery.FindByActive(ctx, in)
}

func (c *LocalUserClient) FindByTrashed(ctx context.Context, in *pb.FindAllUserRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationUserDeleteAt, error) {
	return c.Handler.UserQuery.FindByTrashed(ctx, in)
}

type LocalRoleClient struct {
	Handler *role_handler.Handler
}

// RoleCommandServiceClient implementation
func (c *LocalRoleClient) CreateRole(ctx context.Context, in *pb.CreateRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponseRole, error) {
	return c.Handler.RoleCommand.CreateRole(ctx, in)
}

func (c *LocalRoleClient) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponseRole, error) {
	return c.Handler.RoleCommand.UpdateRole(ctx, in)
}

func (c *LocalRoleClient) TrashedRole(ctx context.Context, in *pb.FindByIdRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponseRole, error) {
	return c.Handler.RoleCommand.TrashedRole(ctx, in)
}

func (c *LocalRoleClient) RestoreRole(ctx context.Context, in *pb.FindByIdRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponseRole, error) {
	return c.Handler.RoleCommand.RestoreRole(ctx, in)
}

func (c *LocalRoleClient) DeleteRolePermanent(ctx context.Context, in *pb.FindByIdRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponseRoleDelete, error) {
	return c.Handler.RoleCommand.DeleteRolePermanent(ctx, in)
}

func (c *LocalRoleClient) RestoreAllRole(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.ApiResponseRoleAll, error) {
	return c.Handler.RoleCommand.RestoreAllRole(ctx, in)
}

func (c *LocalRoleClient) DeleteAllRolePermanent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.ApiResponseRoleAll, error) {
	return c.Handler.RoleCommand.DeleteAllRolePermanent(ctx, in)
}

func (c *LocalRoleClient) AssignRoleToUser(ctx context.Context, in *pb.AssignRoleToUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUserRole, error) {
	return c.Handler.RoleCommand.AssignRoleToUser(ctx, in)
}

func (c *LocalRoleClient) RemoveRoleFromUser(ctx context.Context, in *pb.RemoveRoleFromUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.Handler.RoleCommand.RemoveRoleFromUser(ctx, in)
}

// RoleQueryServiceClient implementation
func (c *LocalRoleClient) FindAllRole(ctx context.Context, in *pb.FindAllRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationRole, error) {
	return c.Handler.RoleQuery.FindAllRole(ctx, in)
}

func (c *LocalRoleClient) FindByIdRole(ctx context.Context, in *pb.FindByIdRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponseRole, error) {
	return c.Handler.RoleQuery.FindByIdRole(ctx, in)
}

func (c *LocalRoleClient) FindByNameRole(ctx context.Context, in *pb.FindByNameRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponseRole, error) {
	return c.Handler.RoleQuery.FindByNameRole(ctx, in)
}

func (c *LocalRoleClient) FindByActive(ctx context.Context, in *pb.FindAllRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationRoleDeleteAt, error) {
	return c.Handler.RoleQuery.FindByActive(ctx, in)
}

func (c *LocalRoleClient) FindByTrashed(ctx context.Context, in *pb.FindAllRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationRoleDeleteAt, error) {
	return c.Handler.RoleQuery.FindByTrashed(ctx, in)
}

func (c *LocalRoleClient) FindByUserId(ctx context.Context, in *pb.FindByIdUserRoleRequest, opts ...grpc.CallOption) (*pb.ApiResponsesRole, error) {
	return c.Handler.RoleQuery.FindByUserId(ctx, in)
}
