defmodule W.Repo.Migrations.AddTablePostsTags do
  use Ecto.Migration

  def change do
    # Primary key and timestamps are not required if
    # using `many_to_many` without schemas
    create table(:posts_tags, primary_key: false) do
      add :post_id, references(:posts, on_delete: :delete_all), primary_key: true
      add :tag_id, references(:tags, on_delete: :delete_all), primary_key: true
      # timestamps()
    end

    create index(:posts_tags, [:post_id])
    create index(:posts_tags, [:tag_id])
    create unique_index(:posts_tags, [:post_id, :tag_id])
  end
end
