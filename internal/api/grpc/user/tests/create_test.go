package tests

//func TestCreate(t *testing.T) {
//	t.Parallel()
//	type userServiceMockFunc func(mc *minimock.Controller) service.UserService
//
//	type args struct {
//		ctx context.Context
//		req *proto.CreateRequest
//	}
//
//	var (
//		ctx = context.Background()
//		mc  = minimock.NewController(t)
//
//		id              = int64(7)
//		name            = "Bob"
//		email           = "bob@gmail.ru"
//		password        = "123"
//		passwordConfirm = "123"
//		role            = "user"
//
//		serviceErr = fmt.Errorf("service error")
//
//		req = &proto.CreateRequest{
//			Name:            name,
//			Email:           email,
//			Password:        password,
//			PasswordConfirm: passwordConfirm,
//			Role:            role,
//		}
//
//		mappedReq = &dtoGrpcUser.CreateRequest{
//			Name:            name,
//			Email:           email,
//			Password:        password,
//			PasswordConfirm: passwordConfirm,
//			Role:            role,
//		}
//
//		res = &proto.CreateResponse{
//			Id: id,
//		}
//	)
//
//	t.Cleanup(mc.Finish)
//
//	tests := []struct {
//		name            string
//		args            args
//		want            *proto.CreateResponse
//		err             error
//		userServiceMock userServiceMockFunc
//	}{
//		{
//			name: "success case",
//			args: args{
//				ctx: ctx,
//				req: req,
//			},
//			want: res,
//			err:  nil,
//			userServiceMock: func(mc *minimock.Controller) service.UserService {
//				mock := serviceMocks.NewUserServiceMock(mc)
//				mock.CreateMock.Expect(ctx, mappedReq).Return(id, nil)
//				return mock
//			},
//		},
//		{
//			name: "service error case",
//			args: args{
//				ctx: ctx,
//				req: req,
//			},
//			want: nil,
//			err:  serviceErr,
//			userServiceMock: func(mc *minimock.Controller) service.UserService {
//				mock := serviceMocks.NewUserServiceMock(mc)
//				mock.CreateMock.Expect(ctx, mappedReq).Return(0, serviceErr)
//				return mock
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//
//			noteServiceMock := tt.userServiceMock(mc)
//			api := user.NewImplementation(noteServiceMock)
//
//			newID, err := api.Create(tt.args.ctx, tt.args.req)
//			require.Equal(t, tt.err, err)
//			require.Equal(t, tt.want, newID)
//		})
//	}
//}
