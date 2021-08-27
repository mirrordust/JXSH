defmodule WWeb.Router do
  use WWeb, :router

  pipeline :browser do
    plug :accepts, ["html", "text"]
    plug :fetch_session
    plug :fetch_flash
    plug :protect_from_forgery
    plug :put_secure_browser_headers
    plug WWeb.Plugs.Locale, "zh"
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/", WWeb do
    pipe_through :browser

    get "/", PageController, :index
    # get "/show", PageController, :show

    get "/redirect_test", PageController, :redirect_test
    get "/err_page", PageController, :err
    get "/about", AboutController, :index
    get "/about/:a", AboutController, :a


    resources "/users", UserController
    resources "/sessions", SessionController, only: [:new, :create, :delete], singleton: true
  end

  scope "/cms", WWeb.CMS, as: :cms do
    pipe_through [:browser, :authenticate_user]

    resources "/pages", PageController
  end

  scope "/admin", WWeb.Admin, as: :admin do
    # pipe_through [:browser, :authenticate_user, :ensure_admin]
    pipe_through [:browser]

    # resources "/posts", PostController
    # resources "/tags", TagController

    # forward "/jobs", BackgroundJob.Plug, name: "w"
  end

  # Other scopes may use custom stacks.
  # scope "/api", WWeb do
  #   pipe_through :api
  # end

  # Enables LiveDashboard only for development
  #
  # If you want to use the LiveDashboard in production, you should put
  # it behind authentication and allow only admins to access it.
  # If your application does not have an admins-only section yet,
  # you can use Plug.BasicAuth to set up some basic authentication
  # as long as you are also using SSL (which you should anyway).
  if Mix.env() in [:dev, :test] do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through :browser
      live_dashboard "/dashboard", metrics: WWeb.Telemetry, ecto_repos: [W.Repo]
    end
  end

  defp authenticate_user(conn, _) do
    case get_session(conn, :user_id) do
      nil ->
        conn
        |> Phoenix.Controller.put_flash(:error, "Login required")
        |> Phoenix.Controller.redirect(to: "/")
        |> halt()

      user_id ->
        assign(conn, :current_user, W.Accounts.get_user!(user_id))
    end
  end
end
