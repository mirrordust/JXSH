defmodule W.Repo.Migrations.AddTablePostTags do
  use Ecto.Migration

  def change do
    # Primary key and timestamps are not required if
    # using `many_to_many` without schemas
    create table(:post_tags, primary_key: false) do
      add :post_id, references(:posts, on_delete: :delete_all), primary_key: true
      add :tag_id, references(:tags, on_delete: :delete_all), primary_key: true
      # timestamps()
    end

    create index(:post_tags, [:post_id])
    create index(:post_tags, [:tag_id])
  end
end
