package rooms

import (
	"sync"

	"github.com/google/uuid"
)

type Room struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	HasDesks    bool      `json:"has_desks"`
	DesksCount  int       `json:"desks_count"`
	IsBooked    bool      `json:"is_booked"`
}

var Rooms = make(map[uuid.UUID]*Room)
var mu sync.Mutex

func GetAllRooms() []*Room {
	mu.Lock()
	defer mu.Unlock()
	var roomsList []*Room
	for _, room := range Rooms {
		roomsList = append(roomsList, room)
	}
	return roomsList
}

func GetRoom(id uuid.UUID) *Room {
	mu.Lock()
	defer mu.Unlock()
	return Rooms[id]
}

func CreateRoom(name, description string, hasDesks bool, desksCount int) *Room {
	mu.Lock()
	defer mu.Unlock()
	room := &Room{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		HasDesks:    hasDesks,
		DesksCount:  desksCount,
		IsBooked:    false,
	}
	Rooms[room.ID] = room
	return room
}

func UpdateRoom(id uuid.UUID, name, description string, hasDesks bool, desksCount int) *Room {
	mu.Lock()
	defer mu.Unlock()
	if room, exists := Rooms[id]; exists {
		room.Name = name
		room.Description = description
		room.HasDesks = hasDesks
		room.DesksCount = desksCount
		return room
	}
	return nil
}

func DeleteRoom(id uuid.UUID) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := Rooms[id]; exists {
		delete(Rooms, id)
		return true
	}
	return false
}

func ToggleBooking(id uuid.UUID) (*Room, error) {
	mu.Lock()
	defer mu.Unlock()
	if room, exists := Rooms[id]; exists {
		if room.HasDesks {
			return nil, ErrRoomWithDesks
		}
		room.IsBooked = !room.IsBooked
		return room, nil
	}
	return nil, ErrRoomNotFound
}

var ErrRoomNotFound = &RoomError{"Room not found"}
var ErrRoomWithDesks = &RoomError{"Cannot book room with desks"}

type RoomError struct {
	Message string
}

func (e *RoomError) Error() string {
	return e.Message
}
