defmodule WWeb.Plugs.AuthenticateUser do
  import Plug.Conn
  require Logger

  def init(default) do
    default
  end

  def call(%Plug.Conn{} = conn, _default) do
    [auth_token | _] = get_req_header(conn, "authorization")
    Logger.debug("[#{__MODULE__}] headers auth: #{inspect(auth_token)}")

    if String.starts_with?(auth_token, "BASIC ") do
      access_token = String.slice(auth_token, 6..-1)

      case WWeb.Storage.get_user_info_by_token(access_token) do
        {:ok, %{user_id: _user_id}} ->
          conn

        :error ->
          Logger.debug("[#{__MODULE__}] no such user for token {#{inspect(access_token)}}")
          auth_fail(conn)
      end
    else
      auth_fail(conn)
    end
  end

  defp auth_fail(%Plug.Conn{} = conn) do
    conn
    |> put_status(:unauthorized)
    |> Phoenix.Controller.json(%{errors: "authenticate fail"})
    |> halt()
  end
end
