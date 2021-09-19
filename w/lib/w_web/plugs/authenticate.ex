defmodule WWeb.Plugs.AuthenticateUser do
  import Plug.Conn

  def init(default) do
    default
  end

  def call(%Plug.Conn{} = conn, _default) do
    [auth_token | _] = get_req_header(conn, "authorization")

    if String.starts_with?(auth_token, "BASIC ") do
      access_token = String.slice(auth_token, 6..-1)
      user_id = get_session(conn, :user_id)

      case WWeb.Storage.get_token_by_user_id(user_id) do
        {:ok, %{access: ^access_token}} -> conn
        # no such user
        :error -> auth_fail(conn)
        # access_token does not match
        _ -> auth_fail(conn)
      end
    else
      auth_fail(conn)
    end
  end

  defp auth_fail(%Plug.Conn{} = conn) do
    conn
    |> Phoenix.Controller.json(%{errors: "authenticate fail"})
    |> halt()
  end
end
