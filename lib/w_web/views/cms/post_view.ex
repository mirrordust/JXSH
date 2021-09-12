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
  end
end
