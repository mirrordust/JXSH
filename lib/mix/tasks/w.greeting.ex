defmodule Mix.Tasks.W.Greeting do
  use Mix.Task

  @shortdoc "Sends a greeting to us from Hello Phoenix"

  @moduledoc """
  This is where we would put any long form documentation or doctests.
  """

  def run(args) do
    Mix.shell().info("Args: #{args}")
    Mix.Task.run("app.start")
    Mix.shell().info("Greetings from the Hello Phoenix Application!")
  end

  # We can define other functions as needed here.
end
