import { useMutation, MutationTuple } from "@apollo/client";
import apolloClient from "@/lib/apolloClient";
import {
  DELETE_ROOM_MUTATION,
  TOGGLE_BOOKING_MUTATION,
  CREATE_ROOM_MUTATION,
  ROOMS_QUERY,
} from "@/lib/queries";
import { Room } from "@/components/RoomCard";

interface DeleteRoomResponse {
  deleteRoom: boolean;
}

export interface ToggleBookingResponse {
  toggleBooking: { id: string; isBooked: boolean };
}

interface CreateRoomResponse {
  createRoom: Room;
}

interface RoomInput {
  name: string;
  description: string;
  hasDesks: boolean;
  desksCount: number;
}

interface UseRoomActionsHook {
  deleteRoom: MutationTuple<DeleteRoomResponse, { id: string }>;
  toggleBooking: MutationTuple<ToggleBookingResponse, { id: string }>;
  createRoom: MutationTuple<CreateRoomResponse, { roomInput: RoomInput }>;
  handleCreateRoom: () => void;
}

export default function useRoomActions(): UseRoomActionsHook {
  const deleteRoom = useMutation<DeleteRoomResponse, { id: string }>(
    DELETE_ROOM_MUTATION,
    {
      client: apolloClient,
      refetchQueries: [{ query: ROOMS_QUERY }],
    }
  );

  const toggleBooking = useMutation<ToggleBookingResponse, { id: string }>(
    TOGGLE_BOOKING_MUTATION,
    {
      client: apolloClient,
    }
  );

  const createRoom = useMutation<CreateRoomResponse, { roomInput: RoomInput }>(
    CREATE_ROOM_MUTATION,
    {
      client: apolloClient,
      refetchQueries: [{ query: ROOMS_QUERY }],
    }
  );

  const handleCreateRoom = () => {
    const hasDesks = Math.random() > 0.5;
    const desksCount = hasDesks ? Math.floor(Math.random() * 10) + 1 : 0;

    createRoom[0]({
      variables: {
        roomInput: {
          name: "Test Room",
          description: "Frontend generated room",
          hasDesks: hasDesks,
          desksCount: desksCount,
        },
      },
    });
  };

  return { deleteRoom, toggleBooking, createRoom, handleCreateRoom };
}
