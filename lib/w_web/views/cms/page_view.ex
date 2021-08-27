defmodule WWeb.CMS.PageView do
  use WWeb, :view

  alias W.CMS

  def author_name(%CMS.Page{author: author}) do
    author.user.name
  end
end
