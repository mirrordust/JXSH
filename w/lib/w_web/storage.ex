defmodule WWeb.Storage do
  use GenServer

  @name __MODULE__

  ## Client API

  def start_link(_) do
    GenServer.start_link(__MODULE__, [], name: @name)
  end

  def get_token_by_user_id(user_id) do
    # directly query from ets table, do not call server
    case :ets.lookup(:token_storage, user_id) do
      [{^user_id, token}] -> {:ok, token}
      [] -> :error
    end
  end

  def put_user_id_token(user_id, %{access: _access_token} = token) do
    GenServer.call(@name, {:create, {user_id, token}})
  end

  def remove_user_id_token(user_id) do
    GenServer.call(@name, {:delete, user_id})
  end

  ## Server callbacks

  @impl true
  def init(_) do
    store = :ets.new(:token_storage, [:named_table])
    {:ok, store}
  end

  @impl true
  def handle_call({:create, {user_id, token}}, _from, state) do
    :ets.insert(:token_storage, {user_id, token})
    {:reply, :ok, state}
  end

  @impl true
  def handle_call({:delete, user_id}, _from, state) do
    :ets.delete(:token_storage, user_id)
    {:reply, :ok, state}
  end

  @impl true
  def handle_info(_msg, state) do
    {:noreply, state}
  end
end
