defmodule WWeb.CMS.TagController do
  use WWeb, :controller

  alias W.CMS
  alias W.CMS.Tag

  action_fallback WWeb.FallbackController

  def index(conn, _params) do
    tags = CMS.list_tags()
    render(conn, "index.json", tags: tags)
  end

  def create(conn, %{"tag" => tag_params}) do
    with {:ok, %Tag{} = tag} <- CMS.create_tag(tag_params) do
      conn
      |> put_status(:created)
      |> put_resp_header("location", Routes.cms_tag_path(conn, :show, tag))
      |> render("show.json", tag: tag)
    end
  end

  def show(conn, %{"id" => id}) do
    tag = CMS.get_tag!(id)
    render(conn, "show.json", tag: tag)
  end

  def update(conn, %{"id" => id, "tag" => tag_params}) do
    tag = CMS.get_tag!(id)

    with {:ok, %Tag{} = tag} <- CMS.update_tag(tag, tag_params) do
      render(conn, "show.json", tag: tag)
    end
  end

  def delete(conn, %{"id" => id}) do
    tag = CMS.get_tag!(id)

    with {:ok, %Tag{}} <- CMS.delete_tag(tag) do
      send_resp(conn, :no_content, "")
    end
  end
end
