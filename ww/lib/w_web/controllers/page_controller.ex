defmodule WWeb.PageController do
  use WWeb, :controller

  alias W.CMS.Post

  def index(conn, params) do
    page = String.to_integer(Map.get(params, "page", "1"))
    page_size = String.to_integer(Map.get(params, "page_size", "5"))

    posts = W.CMS.list_published_posts(page, page_size)
    num_of_all_posts = W.CMS.count_posts()
    num_of_pages = WWeb.Utils.count_page(num_of_all_posts, page_size)

    cond do
      1 <= page and page <= num_of_pages ->
        render(conn, "index.html",
          page_title: "Home",
          posts: posts,
          pagination: true,
          page: page,
          num_of_pages: num_of_pages,
          pagination_labels: WWeb.Utils.generate_pagination(num_of_pages, page, 3)
        )

      true ->
        render(conn, "not_found.html", page_title: "Not Found")
    end
  end

  def tags(conn, params) do
    tags = Map.get(params, "tags")

    case WWeb.Utils.parse_tags(tags) do
      {:ok, tag_names} ->
        posts = W.CMS.list_tagged_posts(tag_names)

        render(conn, "index.html",
          page_title: "Tagged Posts",
          posts: posts,
          pagination: false
        )

      {:error, _error} ->
        tags_with_count = W.CMS.count_tagged_posts()

        ordered_tags_with_count =
          Enum.sort(tags_with_count, fn t1, t2 -> t1.tag_name <= t2.tag_name end)

        render(conn, "tags.html",
          page_title: "Tags",
          tags_with_count: ordered_tags_with_count
        )
    end
  end

  def archives(conn, params) do
    yearmonth = Map.get(params, "yearmonth")

    case WWeb.Utils.parse_yearmonth(yearmonth) do
      {:ok, start_date, end_date} ->
        posts = W.CMS.list_archived_posts(start_date, end_date)

        render(conn, "index.html",
          page_title: "Archives of #{yearmonth}",
          posts: posts,
          pagination: false
        )

      {:error, _error} ->
        yearmonth_list = W.CMS.list_posts_yearmonth()

        post_num_in_yearmonth =
          Enum.reduce(yearmonth_list, %{}, fn x, acc ->
            Map.update(acc, x, 1, &(&1 + 1))
          end)

        ym_list =
          yearmonth_list
          |> Enum.uniq()
          |> Enum.sort(&(&1 >= &2))

        render(conn, "archives.html",
          page_title: "Archives",
          yearmonths: ym_list,
          yearmonths_count: post_num_in_yearmonth
        )
    end
  end

  def show(conn, %{"post_view_name" => view_name}) do
    case W.CMS.get_post_by_view_name(view_name) do
      %Post{} = post ->
        post = W.CMS.inc_post_views(post)

        render(conn, "show.html",
          page_title: post.view_name,
          post: post
        )

      _ ->
        render(conn, "not_found.html", page_title: "Not Found")
    end
  end

  def about(conn, _) do
    case W.CMS.get_post_by_view_name("about") do
      %Post{} = post ->
        cond do
          post.published == true ->
            render(conn, "show.html",
              page_title: "About",
              post: post
            )

          post.published == false ->
            render(conn, "about.html", page_title: "About")
        end

      _ ->
        render(conn, "about.html", page_title: "About")
    end
  end
end
