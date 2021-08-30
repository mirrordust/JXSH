defmodule WWeb.SessionController do
  use WWeb, :controller

  alias W.Accounts

  def new(conn, %{"from" => from}) do
    IO.inspect "🐸"
    IO.inspect from
    render(conn, "new.html", from: from)
  end

  def create(conn, %{"user" => %{"email" => email, "password" => password, "from" => from}}) do
    case Accounts.authenticate_by_email_password(email, password) do
      {:ok, user} ->
        conn
        |> put_flash(:info, "Welcome back!")
        |> put_session(:user_id, user.id)
        |> configure_session(renew: true)
        |> redirect(to: from)

      {:error, :unauthorized} ->
        conn
        |> put_flash(:error, "Bad email/password combination")
        |> redirect(to: Routes.session_path(conn, :new))
    end
  end

  def delete(conn, _) do
    conn
    |> configure_session(drop: true)
    |> redirect(to: "/")
  end
end
