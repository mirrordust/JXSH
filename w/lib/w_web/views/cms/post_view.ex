defmodule WWeb.CMS.PostView do
  use WWeb, :view
  alias WWeb.CMS.PostView

  def render("index.json", %{posts: posts}) do
    %{data: render_many(posts, PostView, "post.json")}
  end

  def render("show.json", %{post: post}) do
    %{data: render_one(post, PostView, "post.json")}
  end

  def render("post.json", %{post: post}) do
    %{
      id: post.id,
      title: post.title,
      body: post.body,
      published: post.published,
      published_at: post.published_at,
      view_name: post.view_name,
      views: post.views,
      tags: render_many(post.tags, WWeb.CMS.TagView, "tag.json")
    }
  end
end
