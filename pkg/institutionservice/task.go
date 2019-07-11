package institutionservice

import (
	"strconv"
	"userService/pkg/camunda"
	"userService/pkg/camunda/pb"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var topics = []string{
	"add_ins",
	"add_mtch",
}

func RunServiceTask(format string, workerNum int) {
	ch := make(chan int, workerNum)
	ctx := context.TODO()
	for i := 0; i < workerNum; i++ {
		go finishRegister(ctx, i, ch)
	}
	c := cron.New()
	_ = c.AddFunc(format, func() {
		ch <- 1
	})
	c.Start()
}

func finishRegister(ctx context.Context, workerId int, ch <-chan int) {
	id := strconv.Itoa(workerId)
	for {
		<-ch
		client := camunda.Get()
		for _, topic := range topics {
			resp, err := client.ExternalTask.FetchAndLock(ctx, &pb.FetchAndLockExternalTaskReq{
				WorkerId: id,
				MaxTasks: 1,
				Topics: []*pb.Topic{
					{
						TopicName:    topic,
						LockDuration: 10000,
					},
				},
			})
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			if camunda.CheckError(resp) {
				logrus.Errorln(camunda.TransError(resp))
				continue
			}
			if len(resp.Item) == 0 {
				continue
			}
			switch topic {
			case "add_ins":
				// todo 机构注册
				err = institutionRegister(resp)
			case "add_mtch":
				// todo 商户注册
				err = merchantRegister(resp)
			}
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			completeResp, err := client.ExternalTask.Complete(ctx, &pb.CompleteExternalTaskReq{
				Id: resp.Item[0].Id,
				Body: &pb.CompleteExternalTaskReqBody{
					WorkerId: id,
				},
			})
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			if camunda.CheckError(completeResp) {
				logrus.Errorln(camunda.TransError(completeResp))
				continue
			}
			logrus.Infoln("任务完成: ", resp.Item[0].Id, ", topic:", topic)
		}
	}
}

func institutionRegister(in *pb.FetchAndLockExternalTaskResp) error {
	return nil
}

func merchantRegister(in *pb.FetchAndLockExternalTaskResp) error {
	return nil
}
