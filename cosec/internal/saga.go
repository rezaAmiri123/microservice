package internal

import (
	"context"
	"github.com/rezaAmiri123/microservice/cosec/internal/domain"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/sec"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

const CreateOrderSagaName = "cosec.CreateOrder"
const CreateOrderReplyChannel = "mallbots.cosec.replies.CreateOrder"

type createOrderSaga struct {
	sec.Saga[*domain.CreateOrderData]
}

func NewCreateOrderSaga() sec.Saga[*domain.CreateOrderData] {
	saga := createOrderSaga{
		Saga: sec.NewSaga[*domain.CreateOrderData](CreateOrderSagaName, CreateOrderReplyChannel),
	}

	// 0. -RejectOrder
	saga.AddStep().
		Compensation(saga.rejectOrder)

	// 1. AuthorizeUser
	saga.AddStep().
		Action(saga.authorizeUser)

	// 2. CreateShoppingList, -CancelShoppingList
	saga.AddStep().
		Action(saga.createShoppingList).
		OnActionReply(depotpb.CreatedShoppingListReply, saga.onCreatedShoppingListReply).
		Compensation(saga.cancelShoppingList)

	// 3. ConfirmPayment
	saga.AddStep().
		Action(saga.confirmPayment)

	// 4. InitiateShopping
	saga.AddStep().
		Action(saga.initiateShopping)

	// 5. ApproveOrder
	saga.AddStep().
		Action(saga.approveOrder)

	return saga
}

func (s createOrderSaga) rejectOrder(ctx context.Context, data *domain.CreateOrderData) (string, ddd.Command, error) {
	return orderingpb.CommandChannel, ddd.NewCommand(orderingpb.RejectOrderCommand, &orderingpb.RejectOrder{Id: data.OrderID}), nil
}

func (s createOrderSaga) authorizeUser(ctx context.Context, data *domain.CreateOrderData) (string, ddd.Command, error) {
	//return am.NewCommand(userspb.AuthorizeUserCommand, userspb.CommandChannel, &userspb.AuthorizeUser{Id: data.UserID})
	return userspb.CommandChannel, ddd.NewCommand(userspb.AuthorizeUserCommand, &userspb.AuthorizeUser{Id: data.UserID}), nil
}

func (s createOrderSaga) createShoppingList(ctx context.Context, data *domain.CreateOrderData) (string, ddd.Command, error) {
	items := make([]*depotpb.CreateShoppingList_Item, len(data.Items))
	for i, item := range data.Items {
		items[i] = &depotpb.CreateShoppingList_Item{
			ProductId: item.ProductID,
			StoreId:   item.StoreID,
			Quantity:  int32(item.Quantity),
		}
	}

	return depotpb.CommandChannel, ddd.NewCommand(depotpb.CreateShoppingListCommand, &depotpb.CreateShoppingList{
		OrderId: data.OrderID,
		Items:   items,
	}), nil
}

func (s createOrderSaga) onCreatedShoppingListReply(ctx context.Context, data *domain.CreateOrderData, reply ddd.Reply) error {
	payload := reply.Payload().(*depotpb.CreatedShoppingList)

	data.ShoppingID = payload.GetId()

	return nil
}

func (s createOrderSaga) cancelShoppingList(ctx context.Context, data *domain.CreateOrderData) (string, ddd.Command, error) {
	//return am.NewCommand(depotpb.CancelShoppingListCommand, depotpb.CommandChannel, &depotpb.CancelShoppingList{Id: data.ShoppingID})
	return depotpb.CommandChannel, ddd.NewCommand(depotpb.CancelShoppingListCommand, &depotpb.CancelShoppingList{Id: data.ShoppingID}), nil
}

func (s createOrderSaga) confirmPayment(ctx context.Context, data *domain.CreateOrderData) (string, ddd.Command, error) {
	return paymentspb.CommandChannel, ddd.NewCommand(paymentspb.ConfirmPaymentCommand, &paymentspb.ConfirmPayment{
		Id:     data.PaymentID,
		Amount: data.Total,
	}), nil
}

func (s createOrderSaga) initiateShopping(ctx context.Context, data *domain.CreateOrderData) (string, ddd.Command, error) {
	return depotpb.CommandChannel, ddd.NewCommand(depotpb.InitiateShoppingCommand, &depotpb.InitiateShopping{Id: data.ShoppingID}), nil
}

func (s createOrderSaga) approveOrder(ctx context.Context, data *domain.CreateOrderData) (string, ddd.Command, error) {
	return orderingpb.CommandChannel, ddd.NewCommand(orderingpb.ApproveOrderCommand, &orderingpb.ApproveOrder{
		Id:         data.OrderID,
		ShoppingId: data.ShoppingID,
	}), nil
}
