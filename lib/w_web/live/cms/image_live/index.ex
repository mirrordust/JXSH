defmodule WWeb.CMS.ImageLive.Index do
  use WWeb, :live_view

  alias W.CMS
  alias W.CMS.Image

  @impl true
  def mount(_params, _session, socket) do
    {:ok, assign(socket, :images, list_images())}
  end

  @impl true
  def handle_params(params, _url, socket) do
    {:noreply, apply_action(socket, socket.assigns.live_action, params)}
  end

  defp apply_action(socket, :edit, %{"id" => id}) do
    socket
    |> assign(:page_title, "Edit Image")
    |> assign(:image, CMS.get_image!(id))
  end

  defp apply_action(socket, :new, _params) do
    socket
    |> assign(:page_title, "New Image")
    |> assign(:image, %Image{})
  end

  defp apply_action(socket, :index, _params) do
    socket
    |> assign(:page_title, "Listing Images")
    |> assign(:image, nil)
  end

  @impl true
  def handle_event("delete", %{"id" => id}, socket) do
    image = CMS.get_image!(id)
    {:ok, _} = CMS.delete_image(image)

    {:noreply, assign(socket, :images, list_images())}
  end

  defp list_images do
    CMS.list_images()
  end
end
