package user

import (
	"context"
	"testing"

	"admin-server/api/internal/config"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserListLogic_GetUserList(t *testing.T) {
	c := config.Config{}
	mockSvcCtx := svc.NewServiceContext(c)
	// init mock service context here

	tests := []struct {
		name       string
		ctx        context.Context
		setupMocks func()
		req        *types.GetUserListReq
		wantErr    bool
		checkResp  func(resp *types.GetUserListResp, err error)
	}{
		{
			name: "response error",
			ctx:  context.Background(),
			setupMocks: func() {
				// mock data for this test case
			},
			req: &types.GetUserListReq{
				// TODO: init your request here
			},
			wantErr: true,
			checkResp: func(resp *types.GetUserListResp, err error) {
				// TODO: Add your check logic here
			},
		},
		{
			name: "successful",
			ctx:  context.Background(),
			setupMocks: func() {
				// Mock data for this test case
			},
			req: &types.GetUserListReq{
				// TODO: init your request here
			},
			wantErr: false,
			checkResp: func(resp *types.GetUserListResp, err error) {
				// TODO: Add your check logic here
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			l := NewGetUserListLogic(tt.ctx, mockSvcCtx)
			resp, err := l.GetUserList(tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}
			tt.checkResp(resp, err)
		})
	}
}
