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

  scope "/", WWeb do
    pipe_through :browser

    get "/", PageController, :index
    get "/about", AboutController, :index

    live "/pagelive", PageLive, :index, as: :livepage
    resources "/users", UserController

    get "/:post_name", PageController, :show
  end

  scope "/cms", WWeb.CMS, as: :cms do
    pipe_through :browser

    get "/admin", AdminController, :index
    resources "/sessions", SessionController, only: [:create, :delete], singleton: true
  end

  scope "/cms", WWeb.CMS, as: :cms do
    pipe_through [:browser, WWeb.Plugs.AuthenticateUser]

    # post
    live "/posts", PostLive.Index, :index
    live "/posts/new", PostLive.Index, :new
    live "/posts/:id/edit", PostLive.Index, :edit

    live "/posts/:id", PostLive.Show, :show
    live "/posts/:id/show/edit", PostLive.Show, :edit

    # tag
    live "/tags", TagLive.Index, :index
    live "/tags/new", TagLive.Index, :new
    live "/tags/:id/edit", TagLive.Index, :edit

    live "/tags/:id", TagLive.Show, :show
    live "/tags/:id/show/edit", TagLive.Show, :edit

    # image
    live "/images", ImageLive.Index, :index
    live "/images/new", ImageLive.Index, :new
    live "/images/:id/edit", ImageLive.Index, :edit

    live "/images/:id", ImageLive.Show, :show
    live "/images/:id/show/edit", ImageLive.Show, :edit
  end

  scope "/mon" do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through :browser
      live_dashboard "/dashboard", metrics: WWeb.Telemetry, ecto_repos: [W.Repo]
    end
  end
end
