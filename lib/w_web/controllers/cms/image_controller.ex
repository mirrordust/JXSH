defmodule WWeb.CMS.ImageController do
  use WWeb, :controller

  alias W.CMS
  alias W.CMS.Image

  action_fallback WWeb.FallbackController

  def index(conn, _params) do
    images = CMS.list_images()
    render(conn, "index.json", images: images)
  end

  def create(conn, %{"image" => image_params}) do
    with {:ok, %Image{} = image} <- CMS.create_image(image_params) do
      conn
      |> put_status(:created)
      |> put_resp_header("location", Routes.cms_image_path(conn, :show, image))
      |> render("show.json", image: image)
    end
  end

  def show(conn, %{"id" => id}) do
    image = CMS.get_image!(id)
    render(conn, "show.json", image: image)
  end

  def update(conn, %{"id" => id, "image" => image_params}) do
    image = CMS.get_image!(id)

    with {:ok, %Image{} = image} <- CMS.update_image(image, image_params) do
      render(conn, "show.json", image: image)
    end
  end

  def delete(conn, %{"id" => id}) do
    image = CMS.get_image!(id)

    with {:ok, %Image{}} <- CMS.delete_image(image) do
      send_resp(conn, :no_content, "")
    end
  end
end
