defmodule W.Repo.Migrations.CreatePosts do
  use Ecto.Migration

  def change do
    create table(:posts) do
      add :title, :string
      add :body, :text
      add :published, :boolean, default: false, null: false
      add :published_at, :naive_datetime
      add :view_name, :string
      add :views, :integer, default: 0

      timestamps()
    end

    create unique_index(:posts, [:title])
    create unique_index(:posts, [:view_name])
  end
end
