defmodule WWeb.Auth.SessionController do
  use WWeb, :controller
  require Logger

  alias W.Auth

  def create(conn, %{"user" => %{"email" => email, "password" => password}}) do
    case Auth.authenticate_by_email_password(email, password) do
      {:ok, user} ->
        token = UUID.uuid1()
        :ok = WWeb.Storage.put_token_user_info(token, %{user_id: user.id})
        Logger.debug("[#{__MODULE__}] new token {#{token}} for user_id {#{inspect(user.id)}}")

        conn
        |> put_status(:ok)
        |> json(%{
          data: %{
            access_token: token
          }
        })

      {:error, error} ->
        conn
        |> put_status(:not_found)
        |> json(%{errors: error})
    end
  end

  def delete(conn, %{"token" => token}) do
    case WWeb.Storage.get_user_info_by_token(token) do
      {:ok, %{user_id: _user_id}} ->
        :ok = WWeb.Storage.remove_token_user_info(token)
        send_resp(conn, :no_content, "")

      :error ->
        conn
        |> put_status(:not_found)
        |> json(%{errors: "the token has not been logged in"})
    end

    # status codes: https://hexdocs.pm/plug/Plug.Conn.Status.html#code/1-known-status-codes
  end
end
