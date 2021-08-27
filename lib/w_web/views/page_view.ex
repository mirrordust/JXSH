defmodule WWeb.PageView do
  use WWeb, :view


  # def render("index.json", %{pages: pages}) do
  #   %{data: render_many(pages, WWeb.PageView, "page.json")}
  # end

  # def render("show.json", %{page: page}) do
  #   %{data: render_one(page, WWeb.PageView, "page.json")}
  # end

  # def render("page.json", %{page: page}) do
  #   %{title: page.title}
  # end

  # def render("index.html", assigns) do
  #   # "rendering with assigns #{inspect(Map.keys(assigns))}"
  #   "rendering with assigns #{inspect(Map.get(assigns, :view_template))}"
  # end
end
