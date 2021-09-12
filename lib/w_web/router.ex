defmodule WWeb.Router do
  use WWeb, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_live_flash
    plug :put_root_layout, {WWeb.LayoutView, :root}
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  pipeline :restricted_api do
    plug :api
    plug WWeb.Plugs.AuthenticateUser
  end

  scope "/", WWeb do
    pipe_through :browser

    get "/", PageController, :index
    get "/about", AboutController, :index

    live "/pagelive", PageLive, :index, as: :livepage
    resources "/users", UserController

    get "/:post_name", PageController, :show
  end

  scope "/auth", WWeb.Auth do
    pipe_through :api

    resources "/sessions", SessionController, only: [:create, :delete], singleton: true
  end

  scope "/cms", WWeb.CMS, as: :cms do
    pipe_through :restricted_api

    resources "/posts", PostController, only: [:index, :show, :create, :update, :delete]
    resources "/tags", TagController, only: [:index, :show, :create, :update, :delete]
    resources "/images", ImageController, only: [:index, :show, :create, :update, :delete]
  end

  scope "/mon" do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through :browser

      live_dashboard "/dashboard", metrics: WWeb.Telemetry, ecto_repos: [W.Repo]
    end
  end
end
