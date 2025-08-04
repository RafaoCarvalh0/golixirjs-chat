defmodule ChatElixirWeb.RoomChannel do
  use Phoenix.Channel

  def join("room:" <> private_room_id, _params, socket) do
    # TODO: implement auth
    if authorized?(socket, private_room_id) do
      {:error, %{reason: "unauthorized"}}
    end
  end

  def handle_in("new_msg", %{"data" => body}, socket) do
    broadcast!(socket, "new_msg", %{data: body})
    {:noreply, socket}
  end

  defp authorized?(_, _), do: false
end
