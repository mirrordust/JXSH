defmodule WWeb.CMS.AdminController do
  use WWeb, :controller

  def index(conn, _) do
    # do not use `app.html.eex`
    case Plug.Conn.get_session(conn, :user_id) do
      nil ->
        conn
        |> put_layout(false)
        |> render("login_form.html")

      _user_id ->
        conn
        |> redirect(to: Routes.cms_tag_index_path(conn, :index))
    end
  end
end
