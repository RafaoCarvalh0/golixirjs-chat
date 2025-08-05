defmodule ChatElixirWeb.RoomChannel do
  use Phoenix.Channel

  def join("room:" <> _private_room_id, _params, socket) do
    {:ok, socket}
  end

  def handle_in("new_msg", %{"data" => body}, socket) do
    broadcast!(socket, "new_msg", %{data: body})
    {:noreply, socket}
  end
end
