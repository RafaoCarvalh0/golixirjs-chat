defmodule ChatElixirWeb.RoomChannelTest do
  use ChatElixir.DataCase

  import Phoenix.ChannelTest

  @endpoint ChatElixirWeb.Endpoint

  @mock_room_id "room:lobby"

  setup do
    unsubscribed_socket =
      socket(ChatElixirWeb.UserSocket, "user:id", %{some_assigns: 1})

    {_, _, subscribed_socket1} =
      ChatElixirWeb.UserSocket
      |> socket("user:joined_user1", %{some_assigns: 1})
      |> subscribe_and_join(@mock_room_id)

    {_, _, subscribed_socket2} =
      ChatElixirWeb.UserSocket
      |> socket("user:joined_user2", %{some_assigns: 1})
      |> subscribe_and_join(@mock_room_id)

    %{
      unsubscribed_socket: unsubscribed_socket,
      subscribed_socket1: subscribed_socket1,
      subscribed_socket2: subscribed_socket2
    }
  end

  describe "join/3" do
    test "returns subcribed socket joining lobby with a valid topic", %{
      unsubscribed_socket: unsubscribed_socket
    } do
      {:ok, %{}, unsubscribed_socket} = subscribe_and_join(unsubscribed_socket, "room:123")

      assert unsubscribed_socket.joined == true
      assert unsubscribed_socket.topic == "room:123"
    end

    test "reises for invalid topic", %{
      unsubscribed_socket: unsubscribed_socket
    } do
      assert_raise(RuntimeError, fn ->
        subscribe_and_join(unsubscribed_socket, "foobarbaz:nonexistent123")
      end)
    end
  end

  describe "handle_in/3" do
    test "subscribed sockets can send messages to the room", %{
      subscribed_socket1: subscribed_socket1,
      subscribed_socket2: subscribed_socket2
    } do
      expected_message = %{data: "Greetings! from #{subscribed_socket1.id}"}

      push(subscribed_socket1, "new_msg", expected_message)

      assert_push "new_msg", ^expected_message

      expected_message2 = %{data: "Hi There! from #{subscribed_socket2.id}"}

      push(subscribed_socket2, "new_msg", expected_message2)

      assert_push "new_msg", ^expected_message2
    end
  end
end
