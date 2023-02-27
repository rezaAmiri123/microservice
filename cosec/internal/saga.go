package internal

import (
	"context"
	"github.com/rezaAmiri123/microservice/cosec/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/am"
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
		Action(saga.createShoppingList)
	//Action(saga.createShoppingList).
	//OnActionReply(depotpb.CreatedShoppingListReply, saga.onCreatedShoppingListReply).
	//Compensation(saga.cancelShoppingList)

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

func (s createOrderSaga) rejectOrder(ctx context.Context, data *domain.CreateOrderData) am.Command {
	//return am.NewCommand(orderingpb.RejectOrderCommand, orderingpb.CommandChannel, &orderingpb.RejectOrder{Id: data.OrderID})
	return nil
}

func (s createOrderSaga) authorizeUser(ctx context.Context, data *domain.CreateOrderData) am.Command {
	return am.NewCommand(userspb.AuthorizeUserCommand, userspb.CommandChannel, &userspb.AuthorizeUser{Id: data.UserID})
}

func (s createOrderSaga) createShoppingList(ctx context.Context, data *domain.CreateOrderData) am.Command {
	//items := make([]*depotpb.CreateShoppingList_Item, len(data.Items))
	//for i, item := range data.Items {
	//	items[i] = &depotpb.CreateShoppingList_Item{
	//		ProductId: item.ProductID,
	//		StoreId:   item.StoreID,
	//		Quantity:  int32(item.Quantity),
	//	}
	//}
	//
	//return am.NewCommand(depotpb.CreateShoppingListCommand, depotpb.CommandChannel, &depotpb.CreateShoppingList{
	//	OrderId: data.OrderID,
	//	Items:   items,
	//})
	return nil
}

func (s createOrderSaga) onCreatedShoppingListReply(ctx context.Context, data *domain.CreateOrderData, reply ddd.Reply) error {
	//payload := reply.Payload().(*depotpb.CreatedShoppingList)
	//
	//data.ShoppingID = payload.GetId()

	return nil
}

func (s createOrderSaga) cancelShoppingList(ctx context.Context, data *domain.CreateOrderData) am.Command {
	//return am.NewCommand(depotpb.CancelShoppingListCommand, depotpb.CommandChannel, &depotpb.CancelShoppingList{Id: data.ShoppingID})
	return nil
}

func (s createOrderSaga) confirmPayment(ctx context.Context, data *domain.CreateOrderData) am.Command {
	//return am.NewCommand(paymentspb.ConfirmPaymentCommand, paymentspb.CommandChannel, &paymentspb.ConfirmPayment{
	//	Id:     data.PaymentID,
	//	Amount: data.Total,
	//})
	return nil
}

func (s createOrderSaga) initiateShopping(ctx context.Context, data *domain.CreateOrderData) am.Command {
	//return am.NewCommand(depotpb.InitiateShoppingCommand, depotpb.CommandChannel, &depotpb.InitiateShopping{Id: data.ShoppingID})
	return nil
}

func (s createOrderSaga) approveOrder(ctx context.Context, data *domain.CreateOrderData) am.Command {
	//return am.NewCommand(orderingpb.ApproveOrderCommand, orderingpb.CommandChannel, &orderingpb.ApproveOrder{
	//	Id:         data.OrderID,
	//	ShoppingId: data.ShoppingID,
	//})
	return nil
}
