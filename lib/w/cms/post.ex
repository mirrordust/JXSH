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

    timestamps()
  end

  @doc false
  def changeset(post, attrs) do
    post
    |> cast(attrs, [:title, :body, :published, :view_name])
    |> validate_required([:title, :body, :published, :view_name])
    |> unique_constraint(:title)
  end
end
