# Script for populating the database. You can run it as:
#
#     mix run priv/repo/seeds.exs
#
# Inside the script, you can read and write to any of your
# repositories directly:
#
#     W.Repo.insert!(%W.SomeSchema{})
#
# We recommend using the bang functions (`insert!`, `update!`
# and so on) as they will fail if something goes wrong.

alias W.Auth.User
alias W.Repo

if Mix.env() in [:dev, :test] do
  %User{}
  |> User.changeset(%{
    name: "shl",
    username: "sh",
    email: "e@mail.com",
    password: "123"
  })
  |> Repo.insert!()
end

if Mix.env() == :prod do
  random_string = fn length ->
    :crypto.strong_rand_bytes(length) |> Base.url_encode64() |> binary_part(0, length)
  end

  name = System.get_env("SEEDS_USER_NAME") || random_string.(10)

  username = System.get_env("SEEDS_USER_USERNAME") || "admin_" <> random_string.(5)

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

  %User{}
  |> User.changeset(%{
    name: name,
    username: username,
    email: email,
    password: password
  })
  |> Repo.insert!()
end
