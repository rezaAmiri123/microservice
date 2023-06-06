package domain

import (
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testUserID      = "user-id"
	testPaymentID   = "payment-id"
	testInvoiceID   = "invoice-id"
	testShoppingId  = "shopping-id"
	testProductID   = "product-id"
	testProductName = "product-name"
	testStoreID     = "store-id"
	testStoreName   = "store-name"
)

func TestOrder_Reject(t *testing.T) {

	type fields struct {
		UserID     string
		PaymentID  string
		InvoiceID  string
		ShoppingID string
		Items      []Item
		Status     OrderStatus
	}
	tests := map[string]struct {
		fields  fields
		on      func(o *es.MockAggregate)
		want    ddd.Event
		wantErr bool
	}{
		"OK": {
			fields: fields{},
			on: func(o *es.MockAggregate) {
				o.On("AddEvent", OrderRejectedEvent, &OrderRejected{})
			},
			want: ddd.NewEvent(OrderRejectedEvent, &Order{}),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			o := &Order{
				Aggregate:  aggregate,
				UserID:     tt.fields.UserID,
				PaymentID:  tt.fields.PaymentID,
				ShoppingID: tt.fields.ShoppingID,
				InvoiceID:  tt.fields.InvoiceID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}

			if tt.on != nil {
				tt.on(aggregate)
			}
			got, err := o.Reject()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				assert.Equal(t, tt.want.EventName(), got.EventName())
				//assert.IsType(t, tt.want.Payload(), got.Payload())
				assert.Equal(t, tt.want.Metadata(), got.Metadata())
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestOrder_CreateOrder(t *testing.T) {
	items := []Item{{
		ProductID:   testProductID,
		ProductName: testProductName,
		StoreID:     testStoreID,
		StoreName:   testStoreName,
		Quantity:    5,
		Price:       100,
	}}
	type fields struct {
		UserID     string
		PaymentID  string
		InvoiceID  string
		ShoppingID string
		Items      []Item
		Status     OrderStatus
	}
	type args struct {
		userID    string
		paymentID string
		items     []Item
	}

	tests := map[string]struct {
		fields  fields
		args    args
		on      func(a *es.MockAggregate)
		want    ddd.Event
		wantErr bool
	}{
		"OK": {
			fields: fields{},
			args: args{
				userID:    testUserID,
				paymentID: testPaymentID,
				items:     items,
			},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", OrderCreatedEvent, &OrderCreated{
					UserID:    testUserID,
					PaymentID: testPaymentID,
					Items:     items,
				})
			},
			want: ddd.NewEvent(OrderCreatedEvent, &Order{
				UserID:    testUserID,
				PaymentID: testPaymentID,
				Items:     items,
			}),
		},
		"NoUserID": {
			fields: fields{},
			args: args{
				//userID:    testUserID,
				paymentID: testPaymentID,
				items:     items,
			},
			wantErr: true,
		},
		"NoPaymentID": {
			fields: fields{},
			args: args{
				userID: testUserID,
				//paymentID: testPaymentID,
				items: items,
			},
			wantErr: true,
		},
		"NoItem": {
			fields: fields{},
			args: args{
				userID:    testUserID,
				paymentID: testPaymentID,
				items:     make([]Item, 0),
			},
			wantErr: true,
		},
		"OrderExists": {
			fields: fields{
				Status: OrderIsPending,
			},
			args: args{
				userID:    testUserID,
				paymentID: testPaymentID,
				items:     items,
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			o := &Order{
				Aggregate:  aggregate,
				UserID:     tt.fields.UserID,
				PaymentID:  tt.fields.PaymentID,
				InvoiceID:  tt.fields.InvoiceID,
				ShoppingID: tt.fields.ShoppingID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}

			if tt.on != nil {
				tt.on(aggregate)
			}

			got, err := o.CreateOrder(tt.args.userID, tt.args.paymentID, tt.args.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				assert.Equal(t, tt.want.EventName(), got.EventName())
				assert.IsType(t, tt.want.Payload(), got.Payload())
				assert.Equal(t, tt.want.Metadata(), got.Metadata())
			} else {
				assert.Nil(t, got)
			}
		})
	}
}
