package main

import (
	"context"
	"errors"
	"github.com/pbthang/go_messaging_app/rpc-server/model"
	"testing"

	"github.com/pbthang/go_messaging_app/rpc-server/kitex_gen/rpc"
	"github.com/stretchr/testify/assert"
)

func TestIMServiceImpl_Send(t *testing.T) {
	chatName := "TestIMServiceImpl_Send:test1:test2"
	type args struct {
		ctx context.Context
		req *rpc.SendRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &rpc.SendRequest{
					Message: &rpc.Message{
						Chat:   chatName,
						Text:   "test",
						Sender: "test1",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &rpc.SendRequest{
					Message: &rpc.Message{
						Chat:   chatName,
						Text:   "!@#$%^&*(())_+{}|:<>?",
						Sender: "",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IMServiceImpl{}
			got, err := s.Send(tt.args.ctx, tt.args.req)
			assert.True(t, errors.Is(err, tt.wantErr))
			assert.NotNil(t, got)
		})
	}
	// Clean up
	model.DeleteChat(chatName)
}

func TestIMServiceImpl_Pull(t *testing.T) {
	// flush db
	model.FlushAll()

	const chatName = "TestIMServiceImpl_Pull:test1:test2"

	// setup
	for i := 0; i < 10; i++ {
		err := model.InsertMessage(&rpc.Message{
			Chat:     chatName,
			Text:     "test",
			Sender:   "test1",
			SendTime: int64(i),
		})
		assert.Nil(t, err)
	}

	type args struct {
		ctx context.Context
		req *rpc.PullRequest
	}
	tests := []struct {
		name    string
		args    args
		len     int
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &rpc.PullRequest{
					Chat:    chatName,
					Cursor:  0,
					Limit:   20,
					Reverse: new(bool),
				},
			},
			len:     10,
			wantErr: nil,
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &rpc.PullRequest{
					Chat:    chatName,
					Cursor:  1,
					Limit:   8,
					Reverse: func() *bool { b := true; return &b }(),
				},
			},
			len:     8,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IMServiceImpl{}
			got, err := s.Pull(tt.args.ctx, tt.args.req)
			assert.True(t, errors.Is(err, tt.wantErr))
			assert.Equal(t, tt.len, len(got.Messages))
			assert.NotNil(t, got)
		})
	}
	// Clean up
	model.DeleteChat(chatName)
}
