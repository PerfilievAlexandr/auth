package tests

//func TestUpdate(t *testing.T) {
//	t.Parallel()
//	type userServiceMockFunc func(mc *minimock.Controller) service.UserService
//
//	type args struct {
//		ctx context.Context
//		req *proto.UpdateRequest
//	}
//
//	var (
//		ctx = context.Background()
//		mc  = minimock.NewController(t)
//
//		id    = int64(7)
//		name  = "Bob"
//		email = "bob@gmail.ru"
//
//		serviceErr = fmt.Errorf("service error")
//
//		req = &proto.UpdateRequest{
//			Id:    id,
//			Name:  wrapperspb.String(name),
//			Email: wrapperspb.String(email),
//		}
//
//		mappedReq = &dtoGrpcUser.UpdateRequest{
//			Id:    id,
//			Name:  name,
//			Email: email,
//		}
//	)
//
//	t.Cleanup(mc.Finish)
//
//	tests := []struct {
//		name            string
//		args            args
//		want            *emptypb.Empty
//		err             error
//		userServiceMock userServiceMockFunc
//	}{
//		{
//			name: "success case",
//			args: args{
//				ctx: ctx,
//				req: req,
//			},
//			want: &emptypb.Empty{},
//			err:  nil,
//			userServiceMock: func(mc *minimock.Controller) service.UserService {
//				mock := serviceMocks.NewUserServiceMock(mc)
//				mock.UpdateMock.Expect(ctx, mappedReq).Return(&emptypb.Empty{}, nil)
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
//				mock.UpdateMock.Expect(ctx, mappedReq).Return(nil, serviceErr)
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
//			newID, err := api.Update(tt.args.ctx, tt.args.req)
//			require.Equal(t, tt.err, err)
//			require.Equal(t, tt.want, newID)
//		})
//	}
//}
