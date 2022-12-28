package subcriber

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notifyModel"
	"managerstudent/modules/notifedProvider/notifyStorage"
)

func SendNotifyAfterAddStudentToClass(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "AddStudentToClassNotify")

	store := notifyStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {
			msg := <-c
			notify := msg.Data().(notifyModel.Notification)
			_ = store.CreateNewNotification(ctx, &notify)
		}
	}()
}
