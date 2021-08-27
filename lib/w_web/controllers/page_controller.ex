defmodule WWeb.PageController do
  use WWeb, :controller

  action_fallback WWeb.FallbackController

  def index(conn, _params) do
    render(conn, "index.html")
    # redirect(conn, external: "https://elixir-lang.org/")
    # conn
    # |> put_flash(:info, "Welcome to Phoenix, from flash info!")
    # |> put_flash(:error, "Let's pretend we have an error.")
    # |> redirect(to: Routes.page_path(conn, :redirect_test))
  end


  # def show(conn, _params) do
  #   page = %{title: "foo"}

  #   render(conn, "show.json", page: page)
  # end

  # def index(conn, _params) do
  #   pages = [%{title: "foo"}, %{title: "bar"}]

  #   render(conn, "index.json", pages: pages)
  # end

  def redirect_test(conn, _params) do
    conn
    # |> put_layout("admin.html")
    |> render("index.html")

    # 会丢失_format参数
  end

  def err(conn, _) do
    with {:ok, _a} <- f({:error, :not_found}) do
      render(conn, :index)
    end
  end

  defp f(arg), do: arg
end
