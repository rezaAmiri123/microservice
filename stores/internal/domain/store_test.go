package domain

import (
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const (
	testStoreName     = "store-name"
	testStoreLocation = "store-location"
)

func TestNewStore(t *testing.T) {
	type args struct {
		id string
	}
	tests := map[string]struct {
		args args
		want *Store
	}{
		"Store": {
			args: args{id: "store-id"},
			want: &Store{
				Aggregate: es.NewAggregate("store-id", StoreAggregate),
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewStore(tt.args.id)

			assert.Equal(t, got.ID(), tt.want.ID())
			assert.Equal(t, got.AggregateName(), tt.want.AggregateName())
		})
	}
}
func TestStore_ApplySnapshot(t *testing.T) {
	type fields struct {
		Name          string
		Location      string
		Participating bool
	}
	type args struct {
		snapshot es.Snapshot
	}

	tests := map[string]struct {
		fields  fields
		args    args
		want    fields
		wantErr bool
	}{
		"V1": {
			fields: fields{},
			args: args{
				snapshot: &StoreV1{
					Name:          testStoreName,
					Location:      testStoreLocation,
					Participating: true,
				},
			},
			want: fields{
				Name:          testStoreName,
				Location:      testStoreLocation,
				Participating: true,
			},
		},
		"Unknown": {
			fields: fields{},
			args: args{
				snapshot: nil,
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := &Store{
				Aggregate:     es.NewMockAggregate(t),
				Name:          tt.fields.Name,
				Location:      tt.fields.Location,
				Participating: tt.fields.Participating,
			}
			if err := s.ApplySnapshot(tt.args.snapshot); (err != nil) != tt.wantErr {
				t.Errorf("ApplySnapshot() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, s.Name, tt.want.Name)
			assert.Equal(t, s.Location, tt.want.Location)
			assert.Equal(t, s.Participating, tt.want.Participating)
		})
	}
}

func TestStore_ToSnapshot(t *testing.T) {
	type fields struct {
		Name          string
		Location      string
		Participating bool
	}

	tests := map[string]struct {
		fields  fields
		want    es.Snapshot
		wantErr bool
	}{
		"V1": {
			fields: fields{
				Name:          testStoreName,
				Location:      testStoreLocation,
				Participating: true,
			},
			want: &StoreV1{
				Name:          testStoreName,
				Location:      testStoreLocation,
				Participating: true,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := &Store{
				Aggregate:     es.NewMockAggregate(t),
				Name:          tt.fields.Name,
				Location:      tt.fields.Location,
				Participating: tt.fields.Participating,
			}
			if got := s.ToSnapshot(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSnapshot() = %v, want = %v", got, tt.want)
			}
		})
	}
}
func TestStore_ApplyEvent(t *testing.T) {
	type fields struct {
		Name          string
		Location      string
		Participating bool
	}
	type args struct {
		event ddd.Event
	}

	tests := map[string]struct {
		fields  fields
		args    args
		want    fields
		wantErr bool
	}{
		"StoreCreated": {
			fields: fields{},
			args: args{
				event: ddd.NewEvent(StoreCreatedEvent, &StoreCreated{
					Name:     testStoreName,
					Location: testStoreLocation,
				}),
			},
			want: fields{
				Name:     testStoreName,
				Location: testStoreLocation,
			},
		},
		"StoreParticipationToggled.Enabled": {
			fields: fields{
				Name:          testStoreName,
				Location:      testStoreLocation,
				Participating: false,
			},
			args: args{
				event: ddd.NewEvent(StoreParticipationEnabledEvent, &StoreParticipationToggled{
					Participating: true,
				}),
			},
			want: fields{
				Name:          testStoreName,
				Location:      testStoreLocation,
				Participating: true,
			},
		},
		"StoreParticipationToggled.Disabled": {
			fields: fields{
				Name:          testStoreName,
				Location:      testStoreLocation,
				Participating: true,
			},
			args: args{
				event: ddd.NewEvent(StoreParticipationEnabledEvent, &StoreParticipationToggled{
					Participating: false,
				}),
			},
			want: fields{
				Name:          testStoreName,
				Location:      testStoreLocation,
				Participating: false,
			},
		},
		"StoreRebranded": {
			fields: fields{
				Name:     "another-store-name",
				Location: testStoreLocation,
			},
			args: args{
				event: ddd.NewEvent(StoreRebrandedEvent, &StoreRebranded{
					Name: testStoreName,
				}),
			},
			want: fields{
				Name:     testStoreName,
				Location: testStoreLocation,
			},
		},
		"Unknown": {
			fields: fields{},
			args: args{
				event: ddd.NewEvent("Unknown", nil),
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := Store{
				Aggregate:     es.NewMockAggregate(t),
				Name:          tt.fields.Name,
				Location:      tt.fields.Location,
				Participating: tt.fields.Participating,
			}

			if err := s.ApplyEvent(tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ApplyEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, s.Name, tt.want.Name)
			assert.Equal(t, s.Location, tt.want.Location)
			assert.Equal(t, s.Participating, tt.want.Participating)
		})
	}
}
func TestStore_EnableParticipation(t *testing.T) {
	type fields struct {
		Name          string
		Location      string
		Participating bool
	}
	type args struct {
		Name     string
		Location string
	}
	tests := map[string]struct {
		fields  fields
		on      func(a *es.MockAggregate)
		want    ddd.Event
		wantErr bool
	}{
		"OK": {
			fields: fields{
				//Name:          testStoreName,
				//Location:      testStoreLocation,
				Participating: false,
			},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", StoreParticipationEnabledEvent, &StoreParticipationToggled{
					Participating: true,
				})
			},
			want: ddd.NewEvent(StoreParticipationEnabledEvent, &Store{
				//Name:          testStoreName,
				//Location:      testStoreLocation,
				Participating: true,
			}),
		},
		"Participating.enabled": {
			fields: fields{
				//Name:          testStoreName,
				//Location:      testStoreLocation,
				Participating: true,
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			s := &Store{
				Aggregate:     aggregate,
				Name:          tt.fields.Name,
				Location:      tt.fields.Location,
				Participating: tt.fields.Participating,
			}

			if tt.on != nil {
				tt.on(aggregate)
			}
			got, err := s.EnableParticipation()

			if (err != nil) != tt.wantErr {
				t.Errorf("EnableParticipation() error = %v, wantErr %v", err, tt.wantErr)
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
func TestStore_DisableParticipation(t *testing.T) {
	type fields struct {
		Name          string
		Location      string
		Participating bool
	}
	type args struct {
		Name     string
		Location string
	}
	tests := map[string]struct {
		fields  fields
		on      func(a *es.MockAggregate)
		want    ddd.Event
		wantErr bool
	}{
		"OK": {
			fields: fields{
				//Name:          testStoreName,
				//Location:      testStoreLocation,
				Participating: true,
			},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", StoreParticipationDisabledEvent, &StoreParticipationToggled{
					Participating: false,
				})
			},
			want: ddd.NewEvent(StoreParticipationDisabledEvent, &Store{
				//Name:          testStoreName,
				//Location:      testStoreLocation,
				Participating: false,
			}),
		},
		"Participating.disabled": {
			fields: fields{
				//Name:          testStoreName,
				//Location:      testStoreLocation,
				Participating: false,
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			s := &Store{
				Aggregate:     aggregate,
				Name:          tt.fields.Name,
				Location:      tt.fields.Location,
				Participating: tt.fields.Participating,
			}

			if tt.on != nil {
				tt.on(aggregate)
			}
			got, err := s.DisableParticipation()

			if (err != nil) != tt.wantErr {
				t.Errorf("DisableParticipation() error = %v, wantErr %v", err, tt.wantErr)
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
func TestStore_Rebrand(t *testing.T) {
	type fields struct {
		Name          string
		Location      string
		Participating bool
	}
	type args struct {
		Name string
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
			args:   args{Name: testStoreName},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", StoreRebrandedEvent, &StoreRebranded{
					Name: testStoreName,
				})
			},
			want: ddd.NewEvent(StoreRebrandedEvent, &Store{
				Name: testStoreName,
				//Location:      testStoreLocation,
				//Participating: false,
			}),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			s := &Store{
				Aggregate: aggregate,
				//Name:          tt.fields.Name,
				//Location:      tt.fields.Location,
				//Participating: tt.fields.Participating,
			}

			if tt.on != nil {
				tt.on(aggregate)
			}
			got, err := s.Rebrand(tt.args.Name)

			if (err != nil) != tt.wantErr {
				t.Errorf("Rebrand() error = %v, wantErr %v", err, tt.wantErr)
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

func TestStore_InitStore(t *testing.T) {
	type fields struct {
		Name     string
		Location string
		//Participating bool
	}
	type args struct {
		Name     string
		Location string
	}

	tests := map[string]struct {
		fields  fields
		args    args
		on      func(a *es.MockAggregate)
		want    ddd.Event
		wantErr bool
	}{
		"OK": {
			fields: fields{
				Name:     "store-name",
				Location: "store-location",
			},
			args: args{
				Name:     "store-name",
				Location: "store-location",
			},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", StoreCreatedEvent, &StoreCreated{
					Name:     "store-name",
					Location: "store-location",
				})
			},
			want: ddd.NewEvent(StoreCreatedEvent, &Store{
				Name:     "store-name",
				Location: "store-location",
			}),
		},
		"NoName": {
			fields: fields{
				Name:     "store-name",
				Location: "store-location",
			},
			args: args{
				//Name:     "store-name",
				Location: "store-location",
			},
			wantErr: true,
		},
		"NoLocation": {
			fields: fields{
				Name:     "store-name",
				Location: "store-location",
			},
			args: args{
				Name: "store-name",
				//Location: "store-location",
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			s := Store{
				Aggregate: aggregate,
				Name:      tt.fields.Name,
				Location:  tt.fields.Location,
				//Participating: tt.fields.Participating,
			}
			if tt.on != nil {
				tt.on(aggregate)
			}
			got, err := s.InitStore(tt.args.Name, tt.args.Location)

			if (err != nil) != tt.wantErr {
				t.Errorf("InitStore error = %v, wantErr %v", err, tt.wantErr)
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
