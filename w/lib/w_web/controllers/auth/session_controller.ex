defmodule WWeb.Auth.SessionController do
  use WWeb, :controller

  alias W.Auth

  def create(conn, %{"user" => %{"email" => email, "password" => password}}) do
    case Auth.authenticate_by_email_password(email, password) do
      {:ok, user} ->
        token = UUID.uuid1()
        :ok = WWeb.Storage.put_user_id_token(user.id, %{access: token})

        conn
        |> put_session(:user_id, user.id)
        |> configure_session(renew: true)
        |> json(%{
          token: %{
            access_token: token
          }
        })

      {:error, error} ->
        conn
        |> json(%{errors: error})
    end
  end

  def delete(conn, _) do
    user_id = get_session(conn, :user_id)
    :ok = WWeb.Storage.remove_user_id_token(user_id)

    conn
    |> configure_session(drop: true)
    |> send_resp(:no_content, "")

    # status codes: https://hexdocs.pm/plug/Plug.Conn.Status.html#code/1-known-status-codes
  end
end
