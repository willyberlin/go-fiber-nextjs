package rooms

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var roomType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Room",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.String},
			"name":        &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"hasDesks":    &graphql.Field{Type: graphql.Boolean},
			"desksCount":  &graphql.Field{Type: graphql.Int},
			"isBooked":    &graphql.Field{Type: graphql.Boolean},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"rooms": &graphql.Field{
				Type: graphql.NewList(roomType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetAllRooms(), nil
				},
			},
			"room": &graphql.Field{
				Type: roomType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idStr, _ := p.Args["id"].(string)
					id, err := uuid.Parse(idStr)
					if err != nil {
						return nil, err
					}
					return GetRoom(id), nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createRoom": &graphql.Field{
				Type: roomType,
				Args: graphql.FieldConfigArgument{
					"roomInput": &graphql.ArgumentConfig{
						Type: graphql.NewInputObject(
							graphql.InputObjectConfig{
								Name: "RoomInput",
								Fields: graphql.InputObjectConfigFieldMap{
									"name":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
									"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
									"hasDesks":    &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
									"desksCount":  &graphql.InputObjectFieldConfig{Type: graphql.Int},
								},
							},
						),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					roomInput := p.Args["roomInput"].(map[string]interface{})
					name, _ := roomInput["name"].(string)
					description, _ := roomInput["description"].(string)
					hasDesks, _ := roomInput["hasDesks"].(bool)
					desksCount, _ := roomInput["desksCount"].(int)
					return CreateRoom(name, description, hasDesks, desksCount), nil
				},
			},
			"updateRoom": &graphql.Field{
				Type: roomType,
				Args: graphql.FieldConfigArgument{
					"id":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"name":        &graphql.ArgumentConfig{Type: graphql.String},
					"description": &graphql.ArgumentConfig{Type: graphql.String},
					"hasDesks":    &graphql.ArgumentConfig{Type: graphql.Boolean},
					"desksCount":  &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idStr, _ := p.Args["id"].(string)
					id, err := uuid.Parse(idStr)
					if err != nil {
						return nil, err
					}
					name, _ := p.Args["name"].(string)
					description, _ := p.Args["description"].(string)
					hasDesks, _ := p.Args["hasDesks"].(bool)
					desksCount, _ := p.Args["desksCount"].(int)
					return UpdateRoom(id, name, description, hasDesks, desksCount), nil
				},
			},
			"deleteRoom": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idStr, _ := p.Args["id"].(string)
					id, err := uuid.Parse(idStr)
					if err != nil {
						return false, err
					}
					return DeleteRoom(id), nil
				},
			},
			"toggleBooking": &graphql.Field{
				Type: roomType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idStr, _ := p.Args["id"].(string)
					id, err := uuid.Parse(idStr)
					if err != nil {
						return nil, err
					}
					return ToggleBooking(id)
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
