defmodule W.CMS.Post do
  use Ecto.Schema
  import Ecto.Changeset

  schema "posts" do
    field :body, :string
    field :published, :boolean, default: false
    field :published_at, :naive_datetime
    field :title, :string
    field :view_name, :string
    field :views, :integer
    many_to_many :tags, W.CMS.Tag, join_through: "post_tags"#, on_replace: :delete

    timestamps()
  end

  @doc false
  def changeset(post, attrs) do
    post
    |> cast(attrs, [:title, :body, :published, :view_name])
    |> validate_required([:title, :body, :published, :view_name])
    |> unique_constraint(:title)
    |> unique_constraint(:view_name)
    |> cast_assoc(:tags, required: true)
  end
end
