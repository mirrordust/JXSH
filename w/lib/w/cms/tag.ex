defmodule W.CMS.Tag do
  use Ecto.Schema
  import Ecto.Changeset
  alias W.CMS.Tag

  schema "tags" do
    field :name, :string

    timestamps()
  end

  @doc false
  def changeset(%Tag{} = tag, attrs) do
    tag
    |> cast(attrs, [:name])
    |> validate_required([:name])
    |> unique_constraint(:name)
  end
end
