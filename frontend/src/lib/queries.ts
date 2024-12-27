import { gql } from "@apollo/client";

export const ROOMS_QUERY = gql`
  query Rooms {
    rooms {
      id
      name
      description
      hasDesks
      desksCount
      isBooked
    }
  }
`;

export const CREATE_ROOM_MUTATION = gql`
  mutation CreateRoom($roomInput: RoomInput!) {
    createRoom(roomInput: $roomInput) {
      id
      name
      description
      hasDesks
      desksCount
      isBooked
    }
  }
`;

export const DELETE_ROOM_MUTATION = gql`
  mutation DeleteRoom($id: String!) {
    deleteRoom(id: $id)
  }
`;

export const TOGGLE_BOOKING_MUTATION = gql`
  mutation ToggleBooking($id: String!) {
    toggleBooking(id: $id) {
      id
      isBooked
      name
      description
      hasDesks
      desksCount
    }
  }
`;
