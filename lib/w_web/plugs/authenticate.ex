defmodule WWeb.Plugs.AuthenticateUser do
  import Plug.Conn

  def init(credential_type = _default) do
    credential_type
  end

  def call(%Plug.Conn{} = conn, _default) do
    case get_session(conn, :user_id) do
      nil ->
        conn
        |> Phoenix.Controller.put_flash(:error, "Login required")
        |> Phoenix.Controller.redirect(to: WWeb.Router.Helpers.cms_admin_path(conn, :index))
        |> halt()

      user_id ->
        assign(conn, :current_user, W.Accounts.get_user!(user_id))
    end
  end
end
