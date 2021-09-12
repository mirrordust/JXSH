defmodule WWeb.CMS.PostController do
  use WWeb, :controller

  alias W.CMS

  def index(conn, _params) do
    posts = CMS.list_posts()
    render(conn, "index.json", posts: posts)
  end

  def show(conn, %{"id" => id}) do
    post = CMS.get_post!(id)
    render(conn, "show.json", post: post)
  end

  def create(conn, %{"post" => post_params}) do
    case CMS.create_post(post_params) do
      {:ok, post} ->
        render(conn, "show.json", post: post)

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, "error.json", changeset: changeset)
    end
  end

  def update(conn, %{"id"=>id, "post"=>post_params}) do
    post = CMS.get_post!(id)

    case CMS.update_post(post, post_params) do
      {:ok, post} -> render(conn, "show.json", post: post)
      {:error, %Ecto.Changeset{}=changeset} ->
        render(conn, "error.json",changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
  end
end
