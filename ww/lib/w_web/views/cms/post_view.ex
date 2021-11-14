defmodule WWeb.CMS.PostView do
  use WWeb, :view
  alias WWeb.CMS.PostView
  alias W.CMS.Post

  def render("index.json", %{posts: posts}) do
    %{data: render_many(posts, PostView, "post.json")}
  end

  def render("show.json", %{post: post}) do
    %{data: render_one(post, PostView, "post.json")}
  end

  def render("post.json", %{post: %Post{} = post}) do
    %{
      id: post.id,
      title: post.title,
      body: post.body,
      published: post.published,
      published_at: post.published_at,
      view_name: post.view_name,
      views: post.views,
      # same as result of NaiveDateTime.to_iso8601/2
      inserted_at: post.inserted_at,
      updated_at: NaiveDateTime.to_iso8601(post.updated_at),
      tags: render_many(post.tags, WWeb.CMS.TagView, "tag.json")
    }
  end
end
