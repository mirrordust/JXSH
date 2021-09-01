defmodule WWeb.CMS.AdminController do
  use WWeb, :controller

  def index(conn, _) do
    # do not use `app.html.eex`
    case Plug.Conn.get_session(conn, :user_id) do
      nil ->
        conn
        |> put_layout(false)
        |> render("index.html")

      user_id ->
        conn
        |> assign(:current_user, W.Accounts.get_user!(user_id))
        |> put_layout(false)
        |> render("index.html")
    end
  end
end
