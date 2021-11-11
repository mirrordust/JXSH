defmodule WWeb.PageController do
  use WWeb, :controller

  alias W.CMS.Post

  def index(conn, _params) do
    render(conn, "index.html")
  end

  def show(conn, %{"post_view_name" => view_name}) do
    case W.CMS.get_post_by_view_name(view_name) do
      %Post{} = post ->
        render(conn, "show.html", post: post)

      _ ->
        render(conn, "not_found.html")
    end
  end

  def about(conn, _) do
    case W.CMS.get_post_by_view_name("about") do
      %Post{} = post ->
        if post.published == true do
          render(conn, "show.html", post: post)
        else
          render(conn, "about.html")
        end

      _ ->
        render(conn, "about.html")
    end
  end
end
