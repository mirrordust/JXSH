defmodule WWeb.Storage do
  use GenServer

  @name __MODULE__

  ## Client API

  def start_link(_) do
    GenServer.start_link(__MODULE__, [], name: @name)
  end

  def get_user_info_by_token(identity_token) do
    # directly query from ets table, do not call server
    case :ets.lookup(:token_storage, identity_token) do
      [{^identity_token, user_info}] -> {:ok, user_info}
      [] -> :error
    end
  end

  def put_token_user_info(identity_token, %{user_id: _user_id} = user_info) do
    GenServer.call(@name, {:create, {identity_token, user_info}})
  end

  def remove_token_user_info(identity_token) do
    GenServer.call(@name, {:delete, identity_token})
  end

  ## Server callbacks

  @impl true
  def init(_) do
    store = :ets.new(:token_storage, [:named_table])
    {:ok, store}
  end

  @impl true
  def handle_call({:create, {identity_token, user_info}}, _from, state) do
    :ets.insert(:token_storage, {identity_token, user_info})
    {:reply, :ok, state}
  end

  @impl true
  def handle_call({:delete, identity_token}, _from, state) do
    :ets.delete(:token_storage, identity_token)
    {:reply, :ok, state}
  end

  @impl true
  def handle_info(_msg, state) do
    {:noreply, state}
  end
end
