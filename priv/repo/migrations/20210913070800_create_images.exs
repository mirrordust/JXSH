defmodule W.Repo.Migrations.CreateImages do
  use Ecto.Migration

  def change do
    create table(:images) do
      add :name, :string
      add :location, :string
      add :metadata, :string

      timestamps()
    end

    create unique_index(:images, [:name])
    create unique_index(:images, [:location])
  end
end
