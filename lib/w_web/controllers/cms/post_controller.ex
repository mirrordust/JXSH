defmodule WWeb.CMS.PostController do
  use WWeb, :controller

  alias W.CMS
  alias W.CMS.Post

  action_fallback WWeb.FallbackController

  def index(conn, _params) do
    posts = CMS.list_posts()
    render(conn, "index.json", posts: posts)
  end

  def create(conn, %{"post" => post_params}) do
    with {:ok, %Post{} = post} <- CMS.create_post(post_params) do
      conn
      |> put_status(:created)
      |> put_resp_header("location", Routes.cms_post_path(conn, :show, post))
      |> render("show.json", post: post)
    end
  end

  def show(conn, %{"id" => id}) do
    post = CMS.get_post!(id)
    render(conn, "show.json", post: post)
  end

  def update(conn, %{"id" => id, "post" => post_params}) do
    post = CMS.get_post!(id)

    with {:ok, %Post{} = post} <- CMS.update_post(post, post_params) do
      render(conn, "show.json", post: post)
    end
  end

  def delete(conn, %{"id" => id}) do
    post = CMS.get_post!(id)

    with {:ok, %Post{}} <- CMS.delete_post(post) do
      send_resp(conn, :no_content, "")
    end
  end
end
