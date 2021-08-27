defmodule WWeb.AboutController do
  use WWeb, :controller

  plug WWeb.Plugs.Locale, "zh" when action in [:index]

  def index(conn, _params) do
    render(conn, "index.html")
  end

  def a(conn, %{"a" => a} = _params) do
    render(conn, "show.html", a: a)
  end
end
