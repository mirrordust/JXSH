defmodule W.CMS.Post do
  use Ecto.Schema
  import Ecto.Changeset
  alias W.CMS.Post

  schema "posts" do
    field :body, :string
    field :published, :boolean, default: false
    field :published_at, :naive_datetime
    field :title, :string
    field :view_name, :string
    field :views, :integer, default: 0, null: false
    # default value of option on_replace is :raise, which wil raise an error
    # when update Post with a new :tags value,
    # change it to :delete to allow the update of :tags
    # see: https://hexdocs.pm/ecto/2.2.11/associations.html#persistence for detail explanation
    many_to_many :tags, W.CMS.Tag, join_through: "posts_tags", on_replace: :delete

    timestamps()
  end

  @doc false
  def changeset(%Post{} = post, attrs) do
    post
    |> cast(attrs, [:title, :body, :published, :view_name])
    |> validate_required([:title, :body, :published, :view_name])
    |> unique_constraint(:title)
    |> unique_constraint(:view_name)
    |> set_published_at()
  end

  def set_published_at(%Ecto.Changeset{} = changeset) do
    # changeset.changes has key :published and its value is true,
    # means unpublished -> published, then set the published date
    if Map.has_key?(changeset.changes, :published) and
         Map.get(changeset.changes, :published) do
      now = NaiveDateTime.utc_now() |> NaiveDateTime.truncate(:second)

      change(changeset, %{published_at: now})
    else
      changeset
    end
  end
end
