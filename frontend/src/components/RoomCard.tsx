import useRoomActions from "@/hooks/useRoomActions";
import { ROOMS_QUERY } from "@/lib/queries";

export interface Room {
  id: string;
  name: string;
  description: string;
  hasDesks: boolean;
  desksCount: number;
  isBooked: boolean;
}

interface ExistingRooms {
  rooms: Room[];
}

export const RoomCard = ({ room }: { room: Room }) => {
  const { deleteRoom, toggleBooking } = useRoomActions();

  return (
    <div
      key={room.id}
      className="bg-gray-700 my-4 p-4 rounded shadow-lg transition-colors duration-200 ease-in-out hover:bg-gray-600"
    >
      <div className="flex items-center">
        <span
          role="img"
          aria-label="room"
          className="text-4xl mr-4"
          style={{ opacity: room.isBooked ? 0.3 : 1 }}
        >
          🏠
        </span>
        <h3 className="text-2xl font-semibold">{room.name}</h3>
      </div>
      <p className="text-gray-300">{room.description}</p>
      <p className="mt-2">
        {room.hasDesks
          ? `Desks Count: ${room.desksCount}`
          : "No desks in this room"}
      </p>
      <p
        className={`mt-2 font-semibold ${
          room.isBooked ? "text-red-500" : "text-green-500"
        }`}
      >
        {room.isBooked ? "Booked" : "Available"}
      </p>
      <div className="flex mt-4">
        <button className="bg-yellow-500 text-white py-2 px-4 rounded">
          Edit
        </button>
        <button
          className="ml-4 bg-red-500 text-white py-2 px-4 rounded"
          onClick={() => deleteRoom[0]({ variables: { id: room.id } })}
        >
          Delete
        </button>
        {!room.hasDesks && (
          <button
            className="ml-4 bg-blue-500 text-white py-2 px-4 rounded"
            onClick={() =>
              toggleBooking[0]({
                variables: { id: room.id },
                update: (cache, { data }) => {
                  if (!data) return;

                  const toggledRoom: Partial<Room> = data.toggleBooking;
                  const existingRooms: ExistingRooms | null = cache.readQuery({
                    query: ROOMS_QUERY,
                  });

                  if (existingRooms) {
                    const newRooms = existingRooms.rooms.map((r: Room) =>
                      r.id === toggledRoom.id ? toggledRoom : r
                    );

                    cache.writeQuery({
                      query: ROOMS_QUERY,
                      data: { rooms: newRooms },
                    });
                  }
                },
              })
            }
          >
            {room.isBooked ? "Unbook" : "Book"}
          </button>
        )}
      </div>
    </div>
  );
};
