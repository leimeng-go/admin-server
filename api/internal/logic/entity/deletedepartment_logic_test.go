package entity

import (
	"context"
	"testing"

	"admin-server/api/internal/config"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeletedepartmentLogic_Deletedepartment(t *testing.T) {
	c := config.Config{}
	mockSvcCtx := svc.NewServiceContext(c)
	// init mock service context here

	tests := []struct {
		name       string
		ctx        context.Context
		setupMocks func()
		req        *types.DeleteDepartmentReq
		wantErr    bool
		checkResp  func(err error)
	}{
		{
			name: "response error",
			ctx:  context.Background(),
			setupMocks: func() {
				// mock data for this test case
			},
			req: &types.DeleteDepartmentReq{
				// TODO: init your request here
			},
			wantErr: true,
			checkResp: func(err error) {
				// TODO: Add your check logic here
			},
		},
		{
			name: "successful",
			ctx:  context.Background(),
			setupMocks: func() {
				// Mock data for this test case
			},
			req: &types.DeleteDepartmentReq{
				// TODO: init your request here
			},
			wantErr: false,
			checkResp: func(err error) {
				// TODO: Add your check logic here
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			l := NewDeletedepartmentLogic(tt.ctx, mockSvcCtx)
			err := l.Deletedepartment(tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)

			}
			tt.checkResp(err)
		})
	}
}
