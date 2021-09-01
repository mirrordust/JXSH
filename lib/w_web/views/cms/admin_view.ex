defmodule WWeb.CMS.AdminView do
  use WWeb, :view

  defp logged_in?(conn) do
    case Plug.Conn.get_session(conn, :user_id) do
      nil -> false
      _user_id -> true
    end
  end

  defp logged_in_username(conn) do
    W.Accounts.get_user!(Plug.Conn.get_session(conn, :user_id)).username
  end
end
