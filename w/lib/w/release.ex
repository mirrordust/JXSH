defmodule W.Release do
  @moduledoc false

  @start_apps [:postgrex, :ecto, :ecto_sql]

  @app :w

  def createdb_and_migrate do
    createdb()
    migrate()
  end

  def createdb do
    load_app()

    Enum.each(@start_apps, &Application.ensure_all_started/1)

    for repo <- repos() do
      :ok = ensure_repo_created(repo)
    end
  end

  defp ensure_repo_created(repo) do
    IO.puts("Create #{inspect(repo)} database if it doesn't exist...")

    case repo.__adapter__.storage_up(repo.config) do
      :ok -> :ok
      {:error, :already_up} -> :ok
      {:error, term} -> {:error, term}
    end
  end

  def migrate do
    load_app()

    Enum.each(@start_apps, &Application.ensure_all_started/1)

    for repo <- repos() do
      {:ok, _, _} = Ecto.Migrator.with_repo(repo, &Ecto.Migrator.run(&1, :up, all: true))
    end
  end

  def migrate1 do
    load_app()

    for repo <- repos() do
      {:ok, _, _} = Ecto.Migrator.with_repo(repo, &Ecto.Migrator.run(&1, :up, all: true))
    end
  end

  def rollback(repo, version) do
    load_app()
    {:ok, _, _} = Ecto.Migrator.with_repo(repo, &Ecto.Migrator.run(&1, :down, to: version))
  end

  defp repos do
    Application.fetch_env!(@app, :ecto_repos)
  end

  defp load_app do
    Application.load(@app)
  end

  # And then in your application you check Application.get_env(@app, :minimal)
  # and start only part of the children when it is set.
  def start_app do
    load_app()
    Application.put_env(@app, :minimal, true)
    Application.ensure_all_started(@app)
  end
end
