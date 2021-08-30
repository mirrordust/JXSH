defmodule WWeb.CMS.ImageLiveTest do
  use WWeb.ConnCase

  import Phoenix.LiveViewTest

  alias W.CMS

  @create_attrs %{location: "some location", name: "some name", size: 120.5}
  @update_attrs %{location: "some updated location", name: "some updated name", size: 456.7}
  @invalid_attrs %{location: nil, name: nil, size: nil}

  defp fixture(:image) do
    {:ok, image} = CMS.create_image(@create_attrs)
    image
  end

  defp create_image(_) do
    image = fixture(:image)
    %{image: image}
  end

  describe "Index" do
    setup [:create_image]

    test "lists all images", %{conn: conn, image: image} do
      {:ok, _index_live, html} = live(conn, Routes.cms_image_index_path(conn, :index))

      assert html =~ "Listing Images"
      assert html =~ image.location
    end

    test "saves new image", %{conn: conn} do
      {:ok, index_live, _html} = live(conn, Routes.cms_image_index_path(conn, :index))

      assert index_live |> element("a", "New Image") |> render_click() =~
               "New Image"

      assert_patch(index_live, Routes.cms_image_index_path(conn, :new))

      assert index_live
             |> form("#image-form", image: @invalid_attrs)
             |> render_change() =~ "can&#39;t be blank"

      {:ok, _, html} =
        index_live
        |> form("#image-form", image: @create_attrs)
        |> render_submit()
        |> follow_redirect(conn, Routes.cms_image_index_path(conn, :index))

      assert html =~ "Image created successfully"
      assert html =~ "some location"
    end

    test "updates image in listing", %{conn: conn, image: image} do
      {:ok, index_live, _html} = live(conn, Routes.cms_image_index_path(conn, :index))

      assert index_live |> element("#image-#{image.id} a", "Edit") |> render_click() =~
               "Edit Image"

      assert_patch(index_live, Routes.cms_image_index_path(conn, :edit, image))

      assert index_live
             |> form("#image-form", image: @invalid_attrs)
             |> render_change() =~ "can&#39;t be blank"

      {:ok, _, html} =
        index_live
        |> form("#image-form", image: @update_attrs)
        |> render_submit()
        |> follow_redirect(conn, Routes.cms_image_index_path(conn, :index))

      assert html =~ "Image updated successfully"
      assert html =~ "some updated location"
    end

    test "deletes image in listing", %{conn: conn, image: image} do
      {:ok, index_live, _html} = live(conn, Routes.cms_image_index_path(conn, :index))

      assert index_live |> element("#image-#{image.id} a", "Delete") |> render_click()
      refute has_element?(index_live, "#image-#{image.id}")
    end
  end

  describe "Show" do
    setup [:create_image]

    test "displays image", %{conn: conn, image: image} do
      {:ok, _show_live, html} = live(conn, Routes.cms_image_show_path(conn, :show, image))

      assert html =~ "Show Image"
      assert html =~ image.location
    end

    test "updates image within modal", %{conn: conn, image: image} do
      {:ok, show_live, _html} = live(conn, Routes.cms_image_show_path(conn, :show, image))

      assert show_live |> element("a", "Edit") |> render_click() =~
               "Edit Image"

      assert_patch(show_live, Routes.cms_image_show_path(conn, :edit, image))

      assert show_live
             |> form("#image-form", image: @invalid_attrs)
             |> render_change() =~ "can&#39;t be blank"

      {:ok, _, html} =
        show_live
        |> form("#image-form", image: @update_attrs)
        |> render_submit()
        |> follow_redirect(conn, Routes.cms_image_show_path(conn, :show, image))

      assert html =~ "Image updated successfully"
      assert html =~ "some updated location"
    end
  end
end
