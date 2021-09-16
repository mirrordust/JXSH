defmodule W.CMSTest do
  use W.DataCase

  alias W.CMS

  describe "tags" do
    alias W.CMS.Tag

    @valid_attrs %{name: "some name"}
    @update_attrs %{name: "some updated name"}
    @invalid_attrs %{name: nil}

    def tag_fixture(attrs \\ %{}) do
      {:ok, tag} =
        attrs
        |> Enum.into(@valid_attrs)
        |> CMS.create_tag()

      tag
    end

    test "list_tags/0 returns all tags" do
      tag = tag_fixture()
      assert CMS.list_tags() == [tag]
    end

    test "get_tag!/1 returns the tag with given id" do
      tag = tag_fixture()
      assert CMS.get_tag!(tag.id) == tag
    end

    test "create_tag/1 with valid data creates a tag" do
      assert {:ok, %Tag{} = tag} = CMS.create_tag(@valid_attrs)
      assert tag.name == "some name"
    end

    test "create_tag/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = CMS.create_tag(@invalid_attrs)
    end

    test "update_tag/2 with valid data updates the tag" do
      tag = tag_fixture()
      assert {:ok, %Tag{} = tag} = CMS.update_tag(tag, @update_attrs)
      assert tag.name == "some updated name"
    end

    test "update_tag/2 with invalid data returns error changeset" do
      tag = tag_fixture()
      assert {:error, %Ecto.Changeset{}} = CMS.update_tag(tag, @invalid_attrs)
      assert tag == CMS.get_tag!(tag.id)
    end

    test "delete_tag/1 deletes the tag" do
      tag = tag_fixture()
      assert {:ok, %Tag{}} = CMS.delete_tag(tag)
      assert_raise Ecto.NoResultsError, fn -> CMS.get_tag!(tag.id) end
    end

    test "change_tag/1 returns a tag changeset" do
      tag = tag_fixture()
      assert %Ecto.Changeset{} = CMS.change_tag(tag)
    end
  end

  describe "images" do
    alias W.CMS.Image

    @valid_attrs %{location: "some location", metadata: "some metadata", name: "some name"}
    @update_attrs %{location: "some updated location", metadata: "some updated metadata", name: "some updated name"}
    @invalid_attrs %{location: nil, metadata: nil, name: nil}

    def image_fixture(attrs \\ %{}) do
      {:ok, image} =
        attrs
        |> Enum.into(@valid_attrs)
        |> CMS.create_image()

      image
    end

    test "list_images/0 returns all images" do
      image = image_fixture()
      assert CMS.list_images() == [image]
    end

    test "get_image!/1 returns the image with given id" do
      image = image_fixture()
      assert CMS.get_image!(image.id) == image
    end

    test "create_image/1 with valid data creates a image" do
      assert {:ok, %Image{} = image} = CMS.create_image(@valid_attrs)
      assert image.location == "some location"
      assert image.metadata == "some metadata"
      assert image.name == "some name"
    end

    test "create_image/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = CMS.create_image(@invalid_attrs)
    end

    test "update_image/2 with valid data updates the image" do
      image = image_fixture()
      assert {:ok, %Image{} = image} = CMS.update_image(image, @update_attrs)
      assert image.location == "some updated location"
      assert image.metadata == "some updated metadata"
      assert image.name == "some updated name"
    end

    test "update_image/2 with invalid data returns error changeset" do
      image = image_fixture()
      assert {:error, %Ecto.Changeset{}} = CMS.update_image(image, @invalid_attrs)
      assert image == CMS.get_image!(image.id)
    end

    test "delete_image/1 deletes the image" do
      image = image_fixture()
      assert {:ok, %Image{}} = CMS.delete_image(image)
      assert_raise Ecto.NoResultsError, fn -> CMS.get_image!(image.id) end
    end

    test "change_image/1 returns a image changeset" do
      image = image_fixture()
      assert %Ecto.Changeset{} = CMS.change_image(image)
    end
  end

  describe "posts" do
    alias W.CMS.Post

    @valid_attrs %{body: "some body", published: true, published_at: ~N[2010-04-17 14:00:00], title: "some title", view_name: "some view_name", views: 42}
    @update_attrs %{body: "some updated body", published: false, published_at: ~N[2011-05-18 15:01:01], title: "some updated title", view_name: "some updated view_name", views: 43}
    @invalid_attrs %{body: nil, published: nil, published_at: nil, title: nil, view_name: nil, views: nil}

    def post_fixture(attrs \\ %{}) do
      {:ok, post} =
        attrs
        |> Enum.into(@valid_attrs)
        |> CMS.create_post()

      post
    end

    test "list_posts/0 returns all posts" do
      post = post_fixture()
      assert CMS.list_posts() == [post]
    end

    test "get_post!/1 returns the post with given id" do
      post = post_fixture()
      assert CMS.get_post!(post.id) == post
    end

    test "create_post/1 with valid data creates a post" do
      assert {:ok, %Post{} = post} = CMS.create_post(@valid_attrs)
      assert post.body == "some body"
      assert post.published == true
      assert post.published_at == ~N[2010-04-17 14:00:00]
      assert post.title == "some title"
      assert post.view_name == "some view_name"
      assert post.views == 42
    end

    test "create_post/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = CMS.create_post(@invalid_attrs)
    end

    test "update_post/2 with valid data updates the post" do
      post = post_fixture()
      assert {:ok, %Post{} = post} = CMS.update_post(post, @update_attrs)
      assert post.body == "some updated body"
      assert post.published == false
      assert post.published_at == ~N[2011-05-18 15:01:01]
      assert post.title == "some updated title"
      assert post.view_name == "some updated view_name"
      assert post.views == 43
    end

    test "update_post/2 with invalid data returns error changeset" do
      post = post_fixture()
      assert {:error, %Ecto.Changeset{}} = CMS.update_post(post, @invalid_attrs)
      assert post == CMS.get_post!(post.id)
    end

    test "delete_post/1 deletes the post" do
      post = post_fixture()
      assert {:ok, %Post{}} = CMS.delete_post(post)
      assert_raise Ecto.NoResultsError, fn -> CMS.get_post!(post.id) end
    end

    test "change_post/1 returns a post changeset" do
      post = post_fixture()
      assert %Ecto.Changeset{} = CMS.change_post(post)
    end
  end
end
