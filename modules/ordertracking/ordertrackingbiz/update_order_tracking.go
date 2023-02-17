package ordertrackingbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/ordertracking/ordertrackingmodel"
)

type updateOrderTrackingStore interface {
	Update(ctx context.Context, data *ordertrackingmodel.UpdateOrderTracking) error
}

type updateOrderBiz struct {
	store updateOrderTrackingStore
}

func NewUpdateOrderBiz(store updateOrderTrackingStore) *updateOrderBiz {
	return &updateOrderBiz{store: store}
}

func (biz *updateOrderBiz) UpdateOrderTracking(ctx context.Context, data *ordertrackingmodel.UpdateOrderTracking) error {

	if err := biz.store.Update(ctx, data); err != nil {
		return common.ErrNoPermission(err)
	}

	return nil
}
