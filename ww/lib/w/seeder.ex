defmodule W.Seeder do
  @moduledoc false

  def seed do
    alias W.Auth.User
    alias W.Repo

    random_string = fn length ->
      :crypto.strong_rand_bytes(length)
      |> Base.url_encode64()
      |> binary_part(0, length)
    end

    email =
      System.get_env("SEEDS_USER_EMAIL") ||
        raise """
        environment variable SEEDS_USER_EMAIL is missing.
        """

    password =
      System.get_env("SEEDS_USER_PASSWORD") ||
        raise """
        environment variable SEEDS_USER_PASSWORD is missing.
        """

    name =
      System.get_env("SEEDS_USER_NAME") ||
        random_string.(10)

    username =
      System.get_env("SEEDS_USER_USERNAME") ||
        "admin_" <> random_string.(5)

    %User{}
    |> User.changeset(%{
      name: name,
      username: username,
      email: email,
      password: password
    })
    |> Repo.insert!()
  end
end
