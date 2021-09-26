defmodule WWeb.CMS.TagView do
  use WWeb, :view
  alias WWeb.CMS.TagView
  alias W.CMS.Tag

  def render("index.json", %{tags: tags}) do
    %{data: render_many(tags, TagView, "tag.json")}
  end

  def render("show.json", %{tag: tag}) do
    %{data: render_one(tag, TagView, "tag.json")}
  end

  def render("tag.json", %{tag: %Tag{} = tag}) do
    %{
      id: tag.id,
      name: tag.name,
      inserted_at: NaiveDateTime.to_string(tag.inserted_at),
      updated_at: NaiveDateTime.to_string(tag.updated_at)
    }
  end
end
