defmodule W.CMS.Image do
  use Ecto.Schema
  import Ecto.Changeset

  schema "images" do
    field :location, :string
    field :name, :string
    field :size, :float

    timestamps()
  end

  @doc false
  def changeset(image, attrs) do
    image
    |> cast(attrs, [:name, :location, :size])
    |> validate_required([:name, :location, :size])
    |> unique_constraint(:name)
    |> unique_constraint(:location)
  end
end
