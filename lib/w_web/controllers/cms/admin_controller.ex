defmodule WWeb.CMS.AdminController do
  use WWeb, :controller


  def index(conn, _params) do
    render(conn, "index.html")
  end
end
