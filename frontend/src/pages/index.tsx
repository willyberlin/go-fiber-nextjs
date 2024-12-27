import { Room, RoomCard } from "@/components/RoomCard";
import useRooms from "@/hooks/useRooms";
import useRoomActions from "@/hooks/useRoomActions";

export default function Home() {
  const { loading, error, filteredRooms, search, setSearch } = useRooms();
  const { handleCreateRoom } = useRoomActions();

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error :(</p>;

  return (
    <div className="dark bg-gray-800 text-white min-h-screen p-8">
      <h1 className="text-4xl font-bold mb-8">Room Planner</h1>
      <h2 className="text-3xl mb-4">Rooms</h2>

      <input
        type="text"
        placeholder="Search rooms..."
        className="px-4 py-2 bg-gray-700 rounded"
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />

      <button
        className="mt-4 bg-green-500 text-white py-2 px-4 rounded"
        onClick={handleCreateRoom}
      >
        Create Room
      </button>

      {filteredRooms.map((room: Room) => (
        <RoomCard key={room.id} room={room} />
      ))}
    </div>
  );
}
