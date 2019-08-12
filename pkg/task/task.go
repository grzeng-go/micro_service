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
	addIns               = "add_ins"
	addMcht              = "add_mcht"
	deleteIns            = "delete_ins"
	deleteMcht           = "delete_mcht"
	updateIns            = "update_ins"
	cancelUpdateIns      = "cancel_update_ins"
	updateMcht           = "update_mcht"
	cancelUpdateMcht     = "cancel_update_mcht"
	insUnregister        = "ins_unregister"
	cancelInsUnregister  = "cancel_ins_unregister"
	freezeMcht           = "freeze_mcht"
	cancelFreezeMcht     = "cancel_freeze_mcht"
	unfreezeMcht         = "unfreeze_mcht"
	cancelUnfreezeMcht   = "cancel_unfreeze_mcht"
	mchtUnregister       = "mcht_unregister"
	cancelMchtUnregister = "cancel_mcht_unregister"
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
	insUnregister,
	cancelInsUnregister,
	freezeMcht,
	cancelFreezeMcht,
	unfreezeMcht,
	cancelUnfreezeMcht,
	mchtUnregister,
	cancelMchtUnregister,
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
					// 机构删除
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
					// 取消更新商户
					err = merchantUpdateCancel(db, resp.Item[0])
				case insUnregister:
					// 机构注销
					err = institutionUnRegister(db, resp.Item[0])
				case cancelInsUnregister:
					// 取消机构注销
					err = institutionCancelUnRegister(db, resp.Item[0])
				case freezeMcht:
					// todo 商户冻结
					err = merchantFreeze(db, resp.Item[0])
				case cancelFreezeMcht:
					// todo 取消商户冻结
					err = cancelMerchantFreeze(db, resp.Item[0])
				case unfreezeMcht:
					// todo 商户解冻
					err = merchantUnFreeze(db, resp.Item[0])
				case cancelUnfreezeMcht:
					// todo 取消商户解冻
					err = cancelMerchantUnFreeze(db, resp.Item[0])
				case mchtUnregister:
					// todo 商户注销
					err = merchantUnregister(db, resp.Item[0])
				case cancelMchtUnregister:
					// todo 取消商户注销
					err = cancelMerchantUnregister(db, resp.Item[0])
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
