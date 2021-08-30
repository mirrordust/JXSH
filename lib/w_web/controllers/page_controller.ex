defmodule WWeb.PageController do
  use WWeb, :controller

  action_fallback WWeb.FallbackController

  def index(conn, _) do
    render(conn, "index.html")
  end

  def show(conn, %{"post_name" => view_name}) do
    render(conn, "show.html", view_name: view_name)
  end
end
