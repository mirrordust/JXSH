defmodule WWeb.LayoutView do
  use WWeb, :view

  defp title() do
    WWeb.Gettext.gettext("Unknown Website")
  end
end
