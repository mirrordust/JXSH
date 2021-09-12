defmodule WWeb.CMS.TagLive.Index do
  use WWeb, :live_view

  alias W.CMS
  alias W.CMS.Tag

  @impl true
  def mount(_params, session, socket) do
    socket =
      authenticate_user(session, socket)
      |> assign(:tag_groups, list_tag_groups())

    {:ok, socket}
  end

  @impl true
  def handle_params(params, _url, socket) do
    {:noreply, apply_action(socket, socket.assigns.live_action, params)}
  end

  defp apply_action(socket, :index, _params) do
    socket
    |> assign(:page_title, "Listing Tags")
    |> assign(:tag, nil)
  end

  @impl true
  def handle_event("create", param, socket) do
    IO.puts("🐸")
    IO.puts(inspect(param))
    IO.puts("🐸")
    {:noreply, socket}
  end

  @impl true
  def handle_event("validate", param, socket) do
    IO.puts("🐸")
    IO.puts(inspect(param))
    IO.puts("🐸")
    {:noreply, socket}
  end

  @impl true
  def handle_event("update", %{"id" => id, "value" => new_value}, socket) do
    tag = CMS.get_tag!(id)
    {:ok, _} = CMS.update_tag(tag, %{name: new_value})

    socket =
      assign(socket, :tag_groups, list_tag_groups())
      |> put_flash(:info, "updated!")

    {:noreply, socket}
  end

  @impl true
  def handle_event("delete", %{"id" => id}, socket) do
    tag = CMS.get_tag!(id)
    {:ok, _} = CMS.delete_tag(tag)

    {:noreply, assign(socket, :tag_groups, list_tag_groups())}
  end

  defp list_tag_groups do
    CMS.list_tags()
    |> group_tags_by_initial_letter()
  end

  defp group_tags_by_initial_letter(tags) do
    Enum.group_by(tags, fn t -> initial_letter(t.name) end)
    |> Enum.map(fn {k, v} -> {k, Enum.sort_by(v, & &1.name)} end)
    |> Enum.sort_by(&elem(&1, 0))
  end

  defp initial_letter(word) do
    if String.first(word) =~ ~r/^[a-z]$/i do
      String.first(word)
      |> String.upcase()
    else
      "#"
    end
  end
end
