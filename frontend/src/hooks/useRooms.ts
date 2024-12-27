import { useState } from "react";
import { useQuery } from "@apollo/client";
import apolloClient from "@/lib/apolloClient";
import { ROOMS_QUERY } from "@/lib/queries";
import { Room } from "@/components/RoomCard";

export default function useRooms() {
  const [search, setSearch] = useState("");
  const { loading, error, data } = useQuery(ROOMS_QUERY, {
    client: apolloClient,
  });

  const filteredRooms = data?.rooms.filter((room: Room) =>
    room.name.toLowerCase().includes(search.toLowerCase())
  );

  return { loading, error, filteredRooms, search, setSearch };
}
