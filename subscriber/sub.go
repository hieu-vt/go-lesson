package subscriber

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/component/asyncjob"
	"lesson-5-goland/pubsub"
	"lesson-5-goland/skio"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx   component.AppContext
	rtEngine skio.RealtimeEngine
}

func NewEngine(appCtx component.AppContext, rtEngine skio.RealtimeEngine) *consumerEngine {
	return &consumerEngine{appCtx: appCtx, rtEngine: rtEngine}
}

func (engine *consumerEngine) Start() error {
	//ps := engine.appCtx.GetPubsub()

	//engine.startSubTopic(common.ChanNoteCreated, asyncjob.NewGroup(
	//	false,
	//	asyncjob.NewJob(SendNotificationAfterCreateNote(engine.appCtx, context.Background(), nil))),
	//)
	//

	//engine.startSubTopic(
	//	common.TopicNoteCreated,
	//	true,
	//	DeleteImageRecordAfterCreateNote(engine.appCtx),
	//	SendEmailAfterCreateNote(engine.appCtx),
	//	EmitRealtimeAfterCreateNote(engine.appCtx, rtEngine),
	//)

	//engine.startSubTopic(
	//	common.TopicNoteCreated,
	//	false,
	//	DeleteImageRecordAfterCreateNote(engine.appCtx),
	//	SendEmailAfterCreateNote(engine.appCtx),
	//)
	// Many sub on a topic

	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		RunIncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		DeCreateLikeCountAfterUserDislikeRestaurant(engine.appCtx),
	)

	engine.startSubTopic(common.TopicHandleOrderWhenUserOrderFood,
		true,
		HandleOrderAfterClientOrder(engine.appCtx),
	)

	engine.startSubTopic(common.TopicEmitEvenWhenUserCreateOrderSuccess,
		false,
		EmitRealtimeAfterOrderEnd(engine.appCtx, engine.rtEngine),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup consumer for:", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}