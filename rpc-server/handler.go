package main

import (
	"context"
	"github.com/pbthang/go_messaging_app/rpc-server/kitex_gen/rpc"
	"github.com/pbthang/go_messaging_app/rpc-server/model"
	"log"
	"time"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(_ context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	req.Message.SendTime = time.Now().Unix()
	log.Println("SEND: ", req.Message.Chat, req.Message.Text, req.Message.Sender, req.Message.SendTime)
	resp := rpc.NewSendResponse()

	err := model.InsertMessage(req.Message)
	if err != nil {
		resp.Code = -1
		resp.Msg = "Failed"
		log.Fatal(err)
		return resp, err
	}

	resp.Code = 0
	resp.Msg = "Success"
	return resp, nil
}

func (s *IMServiceImpl) Pull(_ context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	log.Println("PULL: ", req.Chat, req.Cursor, req.Limit, req.Reverse)
	resp := rpc.NewPullResponse()
	if req.Reverse == nil {
		req.Reverse = new(bool)
		*req.Reverse = false
	}

	messages, hasMore, nextCursor, err := model.GetMessages(req.Chat, req.Cursor, int64(req.Limit), *req.Reverse)
	if err != nil {
		resp.Code = -1
		resp.Msg = "Failed"
		log.Fatal(err)
		return resp, err
	}

	resp.Code = 0
	resp.Msg = "Success"
	resp.Messages = messages
	resp.HasMore = &hasMore
	resp.NextCursor = &nextCursor
	return resp, nil
}
