defmodule WWeb.CMS.AdminView do
  use WWeb, :view

  defp logged_in?(conn) do
    case Plug.Conn.get_session(conn, :user_id) do
      nil -> false
      _user_id -> true
    end
  end
end
