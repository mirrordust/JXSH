defmodule W.CMS.Image do
  use Ecto.Schema
  import Ecto.Changeset

  schema "images" do
    field :location, :string
    field :metadata, :string
    field :name, :string

    timestamps()
  end

  @doc false
  def changeset(image, attrs) do
    image
    |> cast(attrs, [:name, :location, :metadata])
    |> validate_required([:name, :location, :metadata])
    |> unique_constraint(:name)
    |> unique_constraint(:location)
  end
end
