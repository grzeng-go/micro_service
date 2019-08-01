package task

import (
	"strconv"
	"userService/pkg/camunda"
	"userService/pkg/camunda/pb"
	"userService/pkg/common"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	addIns           = "add_ins"
	addMcht          = "add_mcht"
	deleteIns        = "delete_ins"
	deleteMcht       = "delete_mcht"
	updateIns        = "update_ins"
	cancelUpdateIns  = "cancel_update_ins"
	updateMcht       = "update_mcht"
	cancelUpdateMcht = "cancel_update_mcht"
)

var topics = []string{
	addIns,
	addMcht,
	deleteIns,
	deleteMcht,
	updateIns,
	cancelUpdateIns,
	updateMcht,
	cancelUpdateMcht,
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
		for _, t := range topics {
			func(topic string) {
				db := common.DB.Begin()
				defer db.Rollback()
				logrus.Debugln("拉取外部任务")
				resp, err := client.ExternalTask.FetchAndLock(ctx, &pb.FetchAndLockExternalTaskReq{
					WorkerId: id,
					MaxTasks: 1,
					Topics: []*pb.Topic{
						{
							TopicName:    topic,
							LockDuration: 100000,
						},
					},
				})
				if err != nil {
					logrus.Errorln(err)
					return
				}
				if camunda.CheckError(resp) {
					logrus.Errorln(camunda.TransError(resp))
					return
				}
				if len(resp.Item) == 0 {
					return
				}
				switch topic {
				case addIns:
					// 机构注册
					err = institutionRegister(db, resp.Item[0])
				case addMcht:
					// 商户注册
					err = merchantRegister(db, resp.Item[0])
				case deleteIns:
					err = deleteInstitution(db, resp.Item[0])
				case deleteMcht:
					// 删除商户
					err = deleteMerchant(db, resp.Item[0])
				case updateIns:
					// 更新机构
					err = institutionUpdate(db, resp.Item[0])
				case cancelUpdateIns:
					// 取消更新机构
					err = institutionUpdateCancel(db, resp.Item[0])
				case updateMcht:
					// 更新商户
					err = merchantUpdate(db, resp.Item[0])
				case cancelUpdateMcht:
					// 消更新商户
					err = merchantUpdateCancel(db, resp.Item[0])
				}
				if err != nil {
					logrus.Errorln(err)
					return
				}
				completeResp, err := client.ExternalTask.Complete(ctx, &pb.CompleteExternalTaskReq{
					Id: resp.Item[0].Id,
					Body: &pb.CompleteExternalTaskReqBody{
						WorkerId: id,
					},
				})
				if err != nil {
					logrus.Errorln(err)
					return
				}
				if camunda.CheckError(completeResp) {
					logrus.Errorln(camunda.TransError(completeResp))
					return
				}
				db.Commit()
				logrus.Infoln("任务完成: ", resp.Item[0].Id, ", topic:", topic)

			}(t)
		}
	}
}
